package app;

import lombok.Getter;
import lombok.Setter;

import java.util.List;
import java.util.Map;

@Getter
@Setter
public class Response {
    private Map<String, Integer> oks;
    private Map<String, Integer> errors;
}
