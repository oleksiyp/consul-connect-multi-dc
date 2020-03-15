package app;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import lombok.extern.slf4j.Slf4j;
import org.springframework.http.HttpStatus;
import org.springframework.util.MultiValueMap;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestHeader;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.reactive.function.client.WebClient;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.net.InetAddress;
import java.net.UnknownHostException;
import java.util.*;
import java.util.concurrent.ThreadLocalRandom;

@RestController
@Slf4j
public class RecursiveController {
    private final String dc;
    private final String serviceName;
    private final List<String> endpoints;
    private WebClient client;
    private final String hostname;

    public RecursiveController(WebClient.Builder builder) {
        this.client = builder.build();
        dc = System.getenv("DC");
        serviceName = System.getenv("SERVICE_NAME");
        endpoints = Arrays.asList(System.getenv("ENDPOINTS").split(";"));

        try {
            hostname = InetAddress.getLocalHost().getHostName();
        } catch (UnknownHostException e) {
            throw new RuntimeException(e);
        }
    }

    @GetMapping("/")
    public Mono<String> doHealthCheck() {
        return Mono.just("OK");
    }

    @GetMapping(value = "/traffic/${SERVICE_NAME}", produces = "application/json")
    public Mono<String> doRequest(
            @RequestHeader MultiValueMap<String, String> headers
    ) {

        return Flux.fromIterable(endpoints)
                .map(it -> Arrays.asList(it.split(":")))
                .filter(it -> ThreadLocalRandom.current().nextDouble() < Double.parseDouble(it.get(0)))
                .flatMap(it -> client.get()
                        .uri("http://localhost:" + it.get(1) + "/traffic/" + it.get(2))
                        .retrieve()
                        .onStatus(HttpStatus::isError, (resp) -> Mono.error(
                                new RuntimeException("failed: " + resp.statusCode().getReasonPhrase())))
                        .bodyToMono(Response.class)
                        .onErrorResume(thr -> {
                            Response errorResponse = new Response();
                            errorResponse.setErrors(Collections.singletonMap(dc +"/" + serviceName, 1));
                            thr.printStackTrace();
                            return Mono.just(errorResponse);
                        })
                )
                .collectList()
                .map(subResponses -> {
                    Response resp = new Response();
                    Map<String, Integer> oks = new HashMap<>();
                    oks.put(dc + "/" + serviceName, 1);
                    Map<String, Integer> errors = new HashMap<>();
                    for (Response response : subResponses) {
                        addMaps(errors, response.getErrors());
                    }
                    resp.setOks(new TreeMap<>(oks));
                    resp.setErrors(new TreeMap<>(oks));
                    return resp;
                })
                .map(it -> {
                    try {
                        return new ObjectMapper().writerWithDefaultPrettyPrinter().writeValueAsString(it);
                    } catch (JsonProcessingException e) {
                        throw new RuntimeException(e);
                    }
                });
    }

    private void addMaps(Map<String, Integer> accumulator, Map<String, Integer> map) {
        for (String host : map.keySet()) {
            accumulator.compute(host, (k, v) -> (v == null ? 0 : v) + map.get(host));
        }
    }
}
