package main

import (
	"./chess"
	"flag"
	"github.com/go-kit/kit/log"
	"net/http"
	"os"
)

const (
	defaultPort = "9000"
	//defaultRoutingServiceURL = "http://localhost:7878"
)

func main() {
	var (
		addr = envString("PORT", defaultPort)
		// rsurl = envString("ROUTINGSERVICE_URL", defaultRoutingServiceURL)

		httpAddr = flag.String("http.addr", ":"+addr, "HTTP listen address")
		//	routingServiceURL = flag.String("service.routing", rsurl, "routing service URL")

		// ctx = context.Background()
	)

	flag.Parse()
	// Logger
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	var bs chess.Service
	bs = chess.ServiceImpl{}
	httpLogger := log.With(logger, "component", "http")
	mux := http.NewServeMux()
	mux.Handle("/v1/games/", chess.MakeHandler(bs, httpLogger))
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
