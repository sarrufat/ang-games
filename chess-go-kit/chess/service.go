package chess

import "errors"

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("invalid argument")

// Service
type Service interface {
	Solve(problem Problem) (TaskId, error)
	CheckResult(id TaskId) (Result, error)
}

// Piece
type Piece struct {
	// Not used
	//	Label   string `json:"label"`
	Letter  string `json:"letter"`
	Npieces int    `json:"npieces"`
}

// Problem
type Problem struct {
	Dim    string  `json:"dim""`
	Pieces []Piece `json:"pieces"`
}

// Return TaskID

type TaskId struct {
	TaskId string `json:"taskId"`
}

// Result
type Result struct {
	Done        bool             `json:"done"`
	Millis      int64            `json:"ms"`
	NIterations int32            `json:"iterations"`
	Positions   []ResultPosition `json:"positions"`
}

// Position
type ResultPosition struct {
	Piece string `json:"piece"`
	X     int    `json:"x"`
	Y     int    `json:"y"`
}

type ServiceImpl struct {
}

func (ServiceImpl) Solve(problem Problem) (TaskId, error) {
	if len(problem.Dim) < 1 {
		return TaskId{}, ErrInvalidArgument
	}
	return TaskId{TaskId: "12345"}, nil
}

func (ServiceImpl) CheckResult(id TaskId) (Result, error) {
	if len(id.TaskId) < 1 {
		return Result{}, ErrInvalidArgument
	}
	return Result{}, nil
}
