FROM openjdk:8-jdk-alpine
COPY target/service-*.jar service.jar
WORKDIR /
ENTRYPOINT ["java","-Djava.security.egd=file:/dev/./urandom","-jar","/service.jar"]
