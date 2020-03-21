FROM alpine:3.11

RUN addgroup -S service \
    && adduser -S -g service service \
    && apk --no-cache add ca-certificates

WORKDIR /home/service

COPY /bin/service .

RUN chown -R service:service ./

USER service

ENTRYPOINT ["./service"]

