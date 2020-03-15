package app.consul.api;

import retrofit2.Call;
import retrofit2.http.*;

public interface ConsulApi {
    @PUT("config")
    Call<Void> putConfig(
            @Query("dc") String dc,
            @Body ServiceRouterConfig config
    );

    @DELETE("config/{kind}/{service}")
    Call<Void> deleteConfig(
            @Path("kind") String kind,
            @Path("service") String service,
            @Query("dc") String dc
    );

    @GET("config/service-router/{service}")
    Call<ServiceRouterConfig> getServiceRoute(
            @Path("service") String service,
            @Query("dc") String dc
    );
}