package main

import (
	"flag"
	"github.com/go-kit/kit/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sarrufat/ang-games/chess-go-kit/chess"
	"net/http"
	"os"
)

const (
	defaultPort = "9000"
	//defaultRoutingServiceURL = "http://localhost:7878"
)

func getPid() (int, error) {
	return os.Getegid(), nil
}
func main() {
	var (
		addr = envString("PORT", defaultPort)
		// rsurl = envString("ROUTINGSERVICE_URL", defaultRoutingServiceURL)

		httpAddr = flag.String("http.addr", ":"+addr, "HTTP listen address")
		regsitry = prometheus.NewRegistry()
		//	routingServiceURL = flag.String("service.routing", rsurl, "routing service URL")

		// ctx = context.Background()
	)

	flag.Parse()
	// Logger
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	// Prometheus
	regsitry.MustRegister(prometheus.NewGoCollector(),
		prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{
			PidFn: getPid,
		}))
	var bs chess.Service
	bs = chess.NewService(regsitry)
	httpLogger := log.With(logger, "component", "http")
	mux := http.NewServeMux()
	mux.Handle("/v1/games/", chess.MakeHandler(bs, httpLogger, regsitry))
	mux.Handle("/metrics", promhttp.HandlerFor(regsitry, promhttp.HandlerOpts{}))
	http.Handle("/", mux)
	errs := make(chan error, 2)
	go func() {
		rCon := chess.NewResultConsumer()
		rCon()
	}()
	go func() {
		logger.Log("transport", "http", "address", *httpAddr, "msg", "listening")
		errs <- http.ListenAndServe(*httpAddr, nil)
	}()
	logger.Log("terminated", <-errs)
}
func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
