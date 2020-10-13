package chess

import (
	. "./common"
	"context"
	"github.com/go-kit/kit/endpoint"
)

func makeChessEndpont(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(Problem)
		tid, err := s.Solve(req)
		return tid, err
	}
}

func makeCheckEndpont(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(TaskId)
		tid, err := s.CheckResult(req)
		return tid, err
	}
}
