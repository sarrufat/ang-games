package chess

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	. "github.com/sarrufat/ang-games/chess-go-kit/chess/common"
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
