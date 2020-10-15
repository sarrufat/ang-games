package chess

import (
	"errors"
	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sarrufat/ang-games/chess-go-kit/appconf"
	. "github.com/sarrufat/ang-games/chess-go-kit/chess/common"
	"github.com/sarrufat/ang-games/chess-go-kit/chess/solver"
	"log"
	"os"
	"time"
)

// ErrInvalidArgument is returned when one or more arguments are invalid.

var (
	logger             = log.New(os.Stdout, "service:", log.LstdFlags)
	ErrInvalidArgument = errors.New("invalid argument")
	solMap             = make(map[string]Result)
	mapChan            = make(chan TResult)
	timoutQueue        []toutElement
)

// Tout element
type toutElement struct {
	id   string
	tout time.Time
}

// Service
type Service interface {
	Solve(problem Problem) (TaskId, error)
	CheckResult(id TaskId) (Result, error)
}

// Return TaskID

type TaskId struct {
	TaskId string `json:"taskId"`
}

func newTask() TaskId {
	return TaskId{TaskId: uuid.New().String()}
}

type serviceImpl struct {
	registry *prometheus.Registry
}

func NewService(r *prometheus.Registry) Service {
	return serviceImpl{r}
}

func (simpl serviceImpl) Solve(problem Problem) (TaskId, error) {
	if len(problem.Dim) < 1 {
		return TaskId{}, ErrInvalidArgument
	}
	task := newTask()
	s := solver.NewSolver(simpl.registry)
	go func(out chan<- TResult) {
		out <- TResult{
			Result: Result{},
			TaskId: task.TaskId,
		}
		s.Solve(problem, func(ms int64, iter int32, nsol int, res [][]ResultPosition, err error) {
			if err != nil {
				logger.Printf("Solve error: %v", err)
				out <- TResult{
					Result: Result{Done: true, Millis: ms, NIterations: iter},
					TaskId: task.TaskId,
				}
			} else {
				var combs []Combination
				for _, comb := range res {
					combs = append(combs, Combination{Positions: comb})
				}
				out <- TResult{

					Result: Result{Done: true, Millis: ms, NIterations: iter, Combination: combs, NumCombinations: nsol},
					TaskId: task.TaskId,
				}
			}
		})

	}(mapChan)

	return task, nil
}

func NewResultConsumer() func() {
	return func() {
		for r := range mapChan {
			logger.Printf("solMap[r.TaskId] %s %s %v %s %d", r.TaskId, "done", r.Done, "Results", r.Result.NumCombinations)
			solMap[r.TaskId] = r.Result
			timoutQueue = append(timoutQueue, toutElement{
				id:   r.TaskId,
				tout: time.Now().Add(appconf.CacheTimeOut),
			})
		}
	}
}

func NewCleaningCacheResultTask() func() {
	return func() {
		for true {
			time.Sleep(appconf.CacheCleaningPeriod)
			now := time.Now()
			var newQueue []toutElement
			for _, ele := range timoutQueue {
				if ele.tout.Before(now) {
					logger.Println("Removing ", ele.id)
					delete(solMap, ele.id)
				} else {
					newQueue = append(newQueue, ele)
				}
			}
			timoutQueue = newQueue
		}
	}
}
func (serviceImpl) CheckResult(id TaskId) (Result, error) {
	logger.Printf("CheckResult %s", id.TaskId)

	if len(id.TaskId) < 1 {
		return Result{}, ErrInvalidArgument
	}
	res, ok := solMap[id.TaskId]
	if ok {
		//	logger.Printf("%v", res.Combination)
		return res, nil
	}
	logger.Print("CheckResult not found")

	return Result{}, nil
}
