package app

import com.fasterxml.jackson.databind.SerializationFeature
import io.ktor.application.call
import io.ktor.application.install
import io.ktor.client.HttpClient
import io.ktor.client.features.json.JacksonSerializer
import io.ktor.client.features.json.JsonFeature
import io.ktor.client.request.accept
import io.ktor.client.request.get
import io.ktor.features.ContentNegotiation
import io.ktor.http.ContentType
import io.ktor.http.HttpStatusCode
import io.ktor.jackson.jackson
import io.ktor.response.respond
import io.ktor.response.respondText
import io.ktor.routing.get
import io.ktor.routing.routing
import io.ktor.server.engine.embeddedServer
import io.ktor.server.netty.Netty
import org.slf4j.LoggerFactory
import java.lang.RuntimeException
import java.util.concurrent.ThreadLocalRandom

val log = LoggerFactory.getLogger("app.AppKt")

class Response(
    var name: String,
    var oks: Map<String, Int> = mapOf()
)

data class Upstream(
    val name: String,
    val port: Int
)

private operator fun <K> Map<K, Int>.times(other: Map<K, Int>): Map<K, Int> {
    val map = this.toMutableMap()
    for ((key, value) in other) {
        map[key] = (map[key] ?: 0) + value
    }
    return map.toMap()
}

val dc = envVar("dc")
val tag = envVar("tag")
val countIterations = envVar("countIterations").toInt()
val serviceName = envVar("serviceName")
val upstreams = envVar("upstreams")
    .split(",")
    .map {
        val arr = it.split(":")
        Upstream(
            arr[0],
            arr[1].toInt()
        )
    }

val client = HttpClient() {
    install(JsonFeature) {
        serializer = JacksonSerializer {
        }
    }
}

suspend fun queryServices(n: Int): Response {
    if (n == 0) {
        return Response("$dc/$serviceName/$tag")
    }

    val upstream = upstreams[ThreadLocalRandom.current().nextInt(upstreams.size)]
    val url = "http://127.0.0.1:${upstream.port}/traffic/${upstream.name}?n=$n"
    return try {
        val response = client.get<Response>(url) {
            accept(ContentType.Application.Json)
        }
        Response("$dc/$serviceName/$tag", response.oks * mapOf(response.name to 1))
    } catch (err: Exception) {
        log.warn("Exception happened for '{}'", url, err)
        Response("$dc/$serviceName/$tag")
    }
}

fun main() {
    var respondOk = true
    val server = embeddedServer(Netty, 8080) {
        install(ContentNegotiation) {
            jackson {
                enable(SerializationFeature.INDENT_OUTPUT)
            }
        }
        routing {
            get("/healthz") {
                if (respondOk) {
                    call.respondText("OK", ContentType.Text.Plain)
                } else {
                    call.respondText("Failure", status = HttpStatusCode.ServiceUnavailable)
                }
            }
            get("/health_k8s") {
                call.respondText("OK", ContentType.Text.Plain)
            }
            get("/traffic/$serviceName") {
                if (respondOk) {
                    val n = call.parameters["n"]?.toIntOrNull() ?: countIterations
                    call.respond(queryServices(n - 1))
                } else {
                    call.respondText("Failure", status = HttpStatusCode.ServiceUnavailable)
                }
            }
            get("/on") {
                respondOk = true
                call.respondText("OK", ContentType.Text.Plain)
            }
            get("/off") {
                respondOk = false
                call.respondText("OK", ContentType.Text.Plain)
            }
        }
    }
    server.start(wait = true)

}


private fun envVar(name: String) = System.getenv(name)
    ?: throw RuntimeException("missing '$name' environment variable")