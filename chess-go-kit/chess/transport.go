package chess

import (
	"context"
	"encoding/json"
	"errors"
	kitlog "github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	. "github.com/sarrufat/ang-games/chess-go-kit/chess/common"
	"net/http"
)

var (
	ErrBadRouting = errors.New("bad routing")
)

func MakeHandler(bs Service, logger kitlog.Logger, registry *prometheus.Registry) http.Handler {
	solveTrans := httptransport.NewServer(
		makeChessEndpont(bs),
		decodeProblemRequest,
		encodeResponse)
	r := mux.NewRouter()
	r.Handle("/v1/games/chess", solveTrans).Methods("POST")
	// Prometheus
	requests := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "chess",
		Name:      "check_counter",
		Help:      "Total number of requests",
	}, []string{"code"})
	registry.MustRegister(requests)
	checkTrans := httptransport.NewServer(
		makeCheckEndpont(bs),
		decodeCheckRequest,
		encodeResponse)

	r.Handle("/v1/games/chess/{id}", checkTrans).Methods("GET")
	return promhttp.InstrumentHandlerCounter(requests, r)
}
func decodeProblemRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request Problem
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeCheckRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id, ok := mux.Vars(r)["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return TaskId{TaskId: id}, nil
}
