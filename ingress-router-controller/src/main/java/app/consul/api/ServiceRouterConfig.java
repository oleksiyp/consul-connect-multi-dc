package app.consul.api;

import com.fasterxml.jackson.databind.PropertyNamingStrategy.UpperCamelCaseStrategy;
import com.fasterxml.jackson.databind.annotation.JsonNaming;
import lombok.Builder;
import lombok.Data;
import lombok.Value;

import java.util.List;

@Data
@Value
@Builder
@JsonNaming(UpperCamelCaseStrategy.class)
public class ServiceRouterConfig {
    public static final String KIND = "service-router";

    String kind = KIND;
    String name;
    List<Route> routes;

    @Data
    @Value
    @Builder
    @JsonNaming(UpperCamelCaseStrategy.class)
    public static class Route {
        Match match;
        Destination destination;
    }

    @Data
    @Value
    @Builder
    @JsonNaming(UpperCamelCaseStrategy.class)
    public static class Match {
        HTTP http;
    }

    @Data
    @Value
    @Builder
    @JsonNaming(UpperCamelCaseStrategy.class)
    public static class HTTP {
        String pathExact;
        String pathPrefix;
        String pathRegex;
        Header header;
        QueryParam queryParam;
        List<String> methods;
    }

    @Data
    @Value
    @Builder
    @JsonNaming(UpperCamelCaseStrategy.class)
    public static class Header {
        String name;
        boolean present;
        String exact;
        String prefix;
        String suffix;
        String regex;
        boolean invert;
    }

    @Data
    @Value
    @Builder
    @JsonNaming(UpperCamelCaseStrategy.class)
    public static class QueryParam {
        String name;
        boolean present;
        String exact;
        String regex;
    }

    @Data
    @Value
    @Builder
    @JsonNaming(UpperCamelCaseStrategy.class)
    public static class Destination {
        String service;
        String serviceSubset;
        String namespace;
        String prefixRewrite;
        long requestTimeout;
        int numRetries;
        boolean retryOnConnectFailure;
        List<Integer> retryOnStatusCodes;
    }
}
