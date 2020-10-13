package chess

import (
	. "./common"
	"context"
	"encoding/json"
	"errors"
	kitlog "github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	ErrBadRouting = errors.New("bad routing")
)
func MakeHandler(bs Service, logger kitlog.Logger) http.Handler {
	solveTrans := httptransport.NewServer(
		makeChessEndpont(bs),
		decodeProblemRequest,
		encodeResponse)
	r := mux.NewRouter()
	r.Handle("/v1/games/chess", solveTrans).Methods("POST")
	//
	checkTrans := httptransport.NewServer(
		makeCheckEndpont(bs),
		decodeCheckRequest,
		encodeResponse)
	r.Handle("/v1/games/chess/{id}", checkTrans).Methods("GET")
	return r
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

