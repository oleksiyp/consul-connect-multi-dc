package app;

import app.consul.api.ConsulApi;
import app.consul.api.ConsulBlockingApi;
import com.fasterxml.jackson.databind.ObjectMapper;
import okhttp3.OkHttpClient;
import org.springframework.boot.context.properties.EnableConfigurationProperties;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import retrofit2.Retrofit;
import retrofit2.converter.jackson.JacksonConverterFactory;

import java.time.Duration;

@Configuration
@EnableConfigurationProperties(AppProps.class)
public class ConsulApiConfig {
    private final Retrofit retrofit;
    private final Retrofit retrofitLongPolling;

    public ConsulApiConfig(AppProps properties) {
        retrofit = new Retrofit.Builder()
                .baseUrl(properties.getConsul().getBaseUrl())
                .addConverterFactory(JacksonConverterFactory.create(
                        new ObjectMapper().findAndRegisterModules()
                ))
                .build();

        retrofitLongPolling = retrofit.newBuilder()
                .client(new OkHttpClient.Builder()
                        .callTimeout(Duration.ZERO)
                        .readTimeout(Duration.ZERO)
                        .build())
                .build();
    }


    @Bean
    public ConsulApi connectApi() {
        return retrofit.create(ConsulApi.class);
    }

    @Bean
    public ConsulBlockingApi connectBlockingApi() {
        return retrofitLongPolling.create(ConsulBlockingApi.class);
    }
}
