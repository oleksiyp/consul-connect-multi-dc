package app.consul.api;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.databind.PropertyNamingStrategy.UpperCamelCaseStrategy;
import com.fasterxml.jackson.databind.annotation.JsonNaming;
import com.google.common.collect.ImmutableList;
import lombok.Builder;
import lombok.Data;
import lombok.Value;

@Data
@Builder
@Value
@JsonNaming(UpperCamelCaseStrategy.class)
public class RegisterServiceRequest {
    String id;
    String name;
    int port;
    Connect connect;
    String address;
    String kind;
    ImmutableList<Check> checks;

    @Data
    @Builder
    @Value
    @JsonNaming(UpperCamelCaseStrategy.class)
    public static class Connect {
        @JsonProperty("Native")
        boolean isNative;
    }

//    @Data
//    @Builder
//    @Value
//    @JsonNaming(UpperCamelCaseStrategy.class)
//    public static class Proxy {
//        String destinationServiceName;
//        String destinationServiceID;
//        String localServiceAddress;
//        int localServicePort;
//        MeshGateway meshGateway;
//    }
//
//    @Data
//    @Builder
//    @Value
//    @JsonNaming(UpperCamelCaseStrategy.class)
//    public static class MeshGateway {
//        String mode;
//    }

    @Data
    @Builder
    @Value
    @JsonNaming(UpperCamelCaseStrategy.class)
    public static class Check {
        int interval;
        @JsonProperty("HTTP")
        String http;
    }
}
