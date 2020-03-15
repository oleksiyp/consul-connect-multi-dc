package app;

import lombok.Getter;
import lombok.Setter;
import lombok.ToString;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.boot.context.properties.NestedConfigurationProperty;

@Getter
@Setter
@ToString
@ConfigurationProperties("app")
public class AppProps {
    String dc = System.getenv("DC");

    String routerServiceName = System.getenv("ROUTER_SERVICE_NAME");

    int resyncMins = 10;

    @NestedConfigurationProperty
    ConsulProperties consul = new ConsulProperties();

    @Getter
    @Setter
    public static class ConsulProperties {
        String baseUrl = "http://172.17.0.12:8500/v1/";
    }
}
