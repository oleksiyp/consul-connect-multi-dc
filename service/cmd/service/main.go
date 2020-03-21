package main

import (
	"context"
	"encoding/json"
	"flag"
	"github.com/oleksiyp/service/logger"
	"github.com/oleksiyp/service/signals"
	"go.uber.org/zap"
	_ "k8s.io/code-generator/cmd/client-gen/generators"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	serviceName string
	dc          string
	logLevel    string
	zapEncoding string
	port        string
	upstreams   string
)

type Upstream struct {
	probability float64
	name        string
	port        string
}

func init() {
	flag.StringVar(&serviceName, "serviceName", "", "Name of service")
	flag.StringVar(&dc, "dc", "", "Datacenter")
	flag.StringVar(&logLevel, "log-level", "debug", "Log level can be: debug, info, warning, error.")
	flag.StringVar(&zapEncoding, "zap-encoding", "json", "Zap logger encoding.")
	flag.StringVar(&port, "port", "8080", "Port to listen on.")
	flag.StringVar(&upstreams, "upstreams", "", "Upstreams")
}

func main() {
	flag.Parse()

	logger, err := logger.NewLoggerWithEncoding(logLevel, zapEncoding)
	if err != nil {
		log.Fatalf("Error creating logger: %v", err)
	}

	zap.ReplaceGlobals(logger.Desugar())

	defer logger.Sync()

	stopCh := signals.SetupSignalHandler()

	logger.Infof("Starting3 service %s", serviceName)

	upstreamArr := parseUpstreams(logger)

	ListenAndServe(logger, stopCh, upstreamArr)
}

func parseUpstreams(logger *zap.SugaredLogger) []Upstream {
	upstreamArr := make([]Upstream, 0)
	for _, upstream := range strings.Split(upstreams, ";") {
		if len(strings.TrimSpace(upstream)) == 0 {
			continue
		}
		arr := strings.Split(upstream, ":")

		value, err := strconv.ParseFloat(arr[0], 64)
		if err != nil {
			logger.Fatalf("Failed to parse %v", err)
		}

		upstreamObj := Upstream{
			probability: value,
			name:        arr[2],
			port:        arr[1],
		}
		upstreamArr = append(upstreamArr, upstreamObj)
		logger.Infof("Adding upsteream: %s", upstreamObj)
	}
	return upstreamArr
}

type Response struct {
	Name   string
	Oks    map[string]int `json:"oks"`
	Errors map[string]int `json:"errors"`
}

func ListenAndServe(logger *zap.SugaredLogger, stopCh <-chan struct{}, upstreams []Upstream) {
	mux := http.DefaultServeMux
	mux.HandleFunc("/traffic/"+serviceName, func(w http.ResponseWriter, r *http.Request) {
		responses := make(chan Response)
		cnt := 0
		for _, upstream := range upstreams {
			if rand.Float64() > upstream.probability {
				continue
			}

			cnt = cnt + 1
			go func() {
				r, err := http.Get("http://localhost:" + upstream.port + "/traffic/" + upstream.name)
				if err != nil {
					logger.Errorf("Failed to fetch %v", err)
					replyError(responses)
					return
				}

				var resp Response

				err = json.NewDecoder(r.Body).Decode(&resp)
				if err != nil {
					logger.Errorf("Failed to parse json %v", err)
					replyError(responses)
					return
				}

				responses <- resp
			}()
		}

		resultingResponse := aggregateResponse(cnt, responses)

		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "    ")

		w.WriteHeader(http.StatusOK)
		_ = encoder.Encode(&resultingResponse)
	})
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 1 * time.Minute,
		IdleTimeout:  15 * time.Second,
	}

	logger.Infof("Starting HTTP server on port %s", port)

	// run server in background
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			logger.Fatalf("HTTP server crashed %v", err)
		}
	}()

	// wait for SIGTERM or SIGINT
	<-stopCh
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Errorf("HTTP server graceful shutdown failed %v", err)
	} else {
		logger.Info("HTTP server stopped")
	}
}

func aggregateResponse(cnt int, responses chan Response) Response {
	oks := make(map[string]int, 0)
	errors := make(map[string]int, 0)
	for i := 0; i < cnt; i++ {
		response := <-responses
		for k, v := range response.Oks {
			oks[k] = oks[k] + v
		}
		oks[response.Name] = oks[response.Name] + 1
		for k, v := range response.Errors {
			errors[k] = errors[k] + v
		}
	}

	resultingResponse := Response{
		Name:   dc + "/" + serviceName,
		Oks:    oks,
		Errors: errors,
	}
	return resultingResponse
}

func replyError(responses chan Response) {
	errors := make(map[string]int, 0)
	errors[dc+"/"+serviceName] = 1

	errorResponse := Response{
		Name:   dc + "/" + serviceName,
		Oks:    make(map[string]int, 0),
		Errors: errors,
	}
	responses <- errorResponse
}
