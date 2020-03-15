package app;

import lombok.Getter;
import org.springframework.context.ApplicationEvent;

@Getter
public class RouteEvent extends ApplicationEvent {
    private final boolean delete;
    private final String prefix;
    private final String service;

    public RouteEvent(Object source, boolean delete, String prefix, String service) {
        super(source);
        this.delete = delete;
        this.prefix = prefix;
        this.service = service;
    }
}
