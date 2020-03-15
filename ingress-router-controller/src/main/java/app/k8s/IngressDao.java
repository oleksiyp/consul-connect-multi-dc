package app.k8s;

import app.AppProps;
import app.RouteEvent;
import io.kubernetes.client.ApiClient;
import io.kubernetes.client.apis.NetworkingV1Api;
import io.kubernetes.client.apis.NetworkingV1beta1Api;
import io.kubernetes.client.informer.ResourceEventHandler;
import io.kubernetes.client.informer.SharedIndexInformer;
import io.kubernetes.client.informer.SharedInformerFactory;
import io.kubernetes.client.models.NetworkingV1beta1HTTPIngressPath;
import io.kubernetes.client.models.NetworkingV1beta1Ingress;
import io.kubernetes.client.models.NetworkingV1beta1IngressList;
import io.kubernetes.client.models.NetworkingV1beta1IngressRule;
import io.kubernetes.client.util.ClientBuilder;
import lombok.extern.slf4j.Slf4j;
import org.springframework.context.ApplicationEventPublisher;
import org.springframework.stereotype.Component;

import javax.annotation.PostConstruct;
import javax.annotation.PreDestroy;
import java.io.IOException;
import java.util.List;
import java.util.concurrent.TimeUnit;

@Component
@Slf4j
public class IngressDao {
    private final ApiClient apiClient;
    private final NetworkingV1beta1Api networkingV1BetaApi;
    private final AppProps props;
    private SharedInformerFactory factory;
    private final ApplicationEventPublisher publisher;

    public IngressDao(AppProps props, ApplicationEventPublisher publisher) throws IOException {
        this.props = props;
        this.publisher = publisher;
        apiClient = ClientBuilder.defaultClient();
        networkingV1BetaApi = new NetworkingV1beta1Api(apiClient);
    }

    @PostConstruct
    public void init() {
        factory = new SharedInformerFactory(apiClient);

        SharedIndexInformer<NetworkingV1beta1Ingress> informer = factory.sharedIndexInformerFor((params) ->
                        networkingV1BetaApi.listIngressForAllNamespacesCall(
                                null,
                                null,
                                null,
                                null,
                                null,
                                params.resourceVersion,
                                params.timeoutSeconds,
                                params.watch,
                                null,
                                null
                        ),
                NetworkingV1beta1Ingress.class,
                NetworkingV1beta1IngressList.class,
                TimeUnit.MINUTES.toMillis(props.getResyncMins())
        );


        informer.addEventHandler(new IngressEventHandler());

        factory.startAllRegisteredInformers();
    }

    @PreDestroy
    public void destroy() {
        factory.stopAllRegisteredInformers();
    }

    private class IngressEventHandler implements ResourceEventHandler<NetworkingV1beta1Ingress> {
        @Override
        public void onAdd(NetworkingV1beta1Ingress obj) {
            try {
                findOutRules(obj, false);
            } catch (Exception ex) {
                log.error("Failed to add '{}'", obj.getMetadata().getName(), ex);
            }
        }

        @Override
        public void onUpdate(NetworkingV1beta1Ingress oldObj, NetworkingV1beta1Ingress newObj) {
            try {
                findOutRules(newObj, false);
            } catch (Exception ex) {
                log.error("Failed to update '{}'", newObj.getMetadata().getName(), ex);
            }
        }

        @Override
        public void onDelete(NetworkingV1beta1Ingress obj, boolean deletedFinalStateUnknown) {
            try {
                findOutRules(obj, true);
            } catch (Exception ex) {
                log.error("Failed to delete '{}'", obj.getMetadata().getName(), ex);
            }
        }

        private void findOutRules(NetworkingV1beta1Ingress newObj, boolean delete) {
            if (newObj.getSpec() == null) {
                return;
            }

            if (newObj.getSpec().getRules() == null) {
                return;
            }

            for (NetworkingV1beta1IngressRule rule : newObj.getSpec().getRules()) {

                if (rule.getHttp() == null) {
                    continue;
                }
                if (rule.getHttp().getPaths() == null) {
                    continue;
                }

                List<NetworkingV1beta1HTTPIngressPath> pathsList = rule.getHttp().getPaths();

                for (NetworkingV1beta1HTTPIngressPath ingressPath : pathsList) {
                    publisher.publishEvent(new RouteEvent(
                            this,
                            delete,
                            ingressPath.getPath(),
                            ingressPath.getBackend().getServiceName()
                    ));
                }
            }
        }
    }
}
