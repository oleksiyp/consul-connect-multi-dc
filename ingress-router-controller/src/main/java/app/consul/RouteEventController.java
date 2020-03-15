package app.consul;

import app.AppProps;
import app.RouteEvent;
import app.consul.api.ConsulApi;
import app.consul.api.ServiceRouterConfig;
import app.consul.api.ServiceRouterConfig.Match;
import lombok.RequiredArgsConstructor;
import okhttp3.ResponseBody;
import org.springframework.boot.context.properties.EnableConfigurationProperties;
import org.springframework.context.ApplicationListener;
import org.springframework.stereotype.Component;
import retrofit2.Call;
import retrofit2.Response;

import javax.annotation.PostConstruct;
import java.io.IOException;
import java.util.HashMap;
import java.util.Map;
import java.util.stream.Collectors;

import static app.consul.api.ServiceRouterConfig.HTTP;
import static app.consul.api.ServiceRouterConfig.Route;

@Component
@RequiredArgsConstructor
@EnableConfigurationProperties(AppProps.class)
public class RouteEventController implements ApplicationListener<RouteEvent> {
    private final Map<String, String> routingTable = new HashMap<>();
    private final ConsulApi api;
    private final AppProps props;

    @PostConstruct
    public void init() {
        if (props.getRouterServiceName() == null) {
            throw new RuntimeException("bad 'app.routerServiceName' property");
        }
    }

    @Override
    public void onApplicationEvent(RouteEvent event) {
        if (!event.isDelete()) {
            routingTable.put(event.getPrefix(), event.getService());
        } else {
            routingTable.remove(event.getPrefix());
        }

        ServiceRouterConfig routerConfig = buildServiceRouter();
        String dc = props.getDc();
        if (dc == null) {
            dc = "";
        }
        execute(api.putConfig(dc, routerConfig));
    }

    private void execute(Call<?> call) {
        try {
            Response<?> response = call.execute();
            if (!response.isSuccessful()) {
                ResponseBody body = response.errorBody();
                if (body != null) {
                    throw new RuntimeException("failed to change configuration: " + body.string());
                } else {
                    throw new RuntimeException("failed to change configuration");
                }
            }
        } catch (IOException e) {
            throw new RuntimeException("failed to call Consul", e);
        }
    }

    private ServiceRouterConfig buildServiceRouter() {
        return ServiceRouterConfig.builder()
                .name(props.getRouterServiceName())
                .routes(routingTable
                        .entrySet()
                        .stream()
                        .map(entry -> Route.builder()
                                .match(Match.builder()
                                        .http(HTTP.builder()
                                                .pathPrefix(entry.getKey())
                                                .build())
                                        .build())
                                .destination(ServiceRouterConfig.Destination.builder()
                                        .service(entry.getValue())
                                        .build())
                                .build())
                        .collect(Collectors.toList()))
                .build();
    }
}
