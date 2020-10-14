package solver

import (
	"errors"
	//	"github.com/go-kit/kit/metrics"
	"github.com/prometheus/client_golang/prometheus"
	chess "github.com/sarrufat/ang-games/chess-go-kit/chess/common"
	"sort"
	"strconv"
	"time"
)

var (
	requestDuration      prometheus.Histogram
	solutionCountByDim   *prometheus.CounterVec
	prometheusRegistered = false
)

type Solver interface {
	threateningForPType([]chess.Piece) map[byte]map[Pos]ThreateningVector
	Solve(p chess.Problem, observe func(ms int64, iter int32, nsol int, res [][]chess.ResultPosition, err error))
	setBoard(b *Board)
}

type solver struct {
	board *Board
}

func NewSolver(r *prometheus.Registry) Solver {
	if !prometheusRegistered {
		prometheusRegistered = true
		r.MustRegister(requestDuration, solutionCountByDim)
	}
	return &solver{}
}

func init() {
	requestDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: "chess",
		Name:      "solve_time",
		Help:      "Solver time in milliseconds",
	})
	solutionCountByDim = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "chess",
		Name:      "solutions",
		Help:      "Solutions found",
	}, []string{"dim"})
}

type anonRec func(chess.Piece, []chess.Piece, ThreateningVector, []chess.ResultPosition)

func (s *solver) Solve(p chess.Problem, observe func(ms int64, iter int32, nsolutions int, res [][]chess.ResultPosition, err error)) {
	dim, err := strconv.Atoi(p.Dim[0:1])
	if err != nil {
		observe(0, 0, 0, [][]chess.ResultPosition{}, err)
		return
	}
	s.board = &Board{
		Dimension: dim,
	}
	tMap := s.threateningForPType(p.Pieces)
	var iterations int32
	var foundSolutions int
	var recSolve anonRec
	var results [][]chess.ResultPosition
	t0 := time.Now()
	defer func() {
		elapsed := time.Since(t0)
		requestDuration.Observe(float64(elapsed.Milliseconds()))
		solutionCountByDim.WithLabelValues(strconv.Itoa(dim)).Add(float64(foundSolutions))
	}()
	// Prometheus
	recSolve = func(actual chess.Piece, reaming []chess.Piece, tVector ThreateningVector, resPos []chess.ResultPosition) {
		if err != nil {
			return
		}
		iterations += 1
		// Timeout
		if iterations%10000 == 0 {
			if time.Since(t0).Seconds() >= 60 {
				err = errors.New("Computation time exceeded")
				return
			}
		}
		//	fmt.Printf("iterations = %d %d %v\n", iterations, len(reaming), results)
		for x := 0; x < dim; x++ {
			for y := 0; y < dim; y++ {
				lastIdx := len(resPos) - 1
				// Skip permutation if already calculated tree
				if len(resPos) > 0 && resPos[lastIdx].Piece[0] == actual.Letter[0] {
					nCurIdx := x*dim + y
					nLastIdx := resPos[lastIdx].X*dim + resPos[lastIdx].Y
					if nCurIdx <= nLastIdx {
						continue
					}
				}
				actualRPos := append(resPos, chess.ResultPosition{Piece: actual.Letter, X: x, Y: y})
				if !searchResult(resPos, x, y) && !threatPos(tVector, x, y) {
					vector := tMap[actual.Letter[0]][Pos{X: x, Y: y}]
					if !threatPosInResult(vector, resPos) {
						if len(reaming) == 0 { // solution foun
							foundSolutions += 1
							if foundSolutions < 10000 {
								var solution = make([]chess.ResultPosition, len(actualRPos))
								copy(solution, actualRPos)
								results = append(results, solution)
							}
							// Limit max results
							// if len(results) > 10000 {
							//	err = errors.New("Results limit exceeded")
							// }
						} else {
							expVector := append(tVector, vector...)
							expResult := actualRPos
							recSolve(reaming[0], reaming[1:], expVector, expResult)
						}

					}
				}
			}
		}
	}
	fPieces := flatten(p.Pieces, tMap)
	//	for x := 0; x < dim; x++ {
	//for y := 0; y < dim; y++ {
	// resPos := []chess.ResultPosition{{Piece: fPieces[0].Letter, X: x, Y: x}}
	recSolve(fPieces[0], fPieces[1:], ThreateningVector{}, []chess.ResultPosition{})
	//	}
	//	}
	observe(time.Since(t0).Milliseconds(), iterations, foundSolutions, results, err)
}
func flatten(pieces []chess.Piece, pmap map[byte]map[Pos]ThreateningVector) []chess.Piece {
	var out []chess.Piece
	for _, p := range priority(pieces, pmap) {
		for i := 0; i < p.Npieces; i++ {
			out = append(out, p)
		}
	}
	return out
}

type arrayPieces struct {
	pieces []chess.Piece
	pmap   map[string]int
}

func (ap arrayPieces) Len() int {
	return len(ap.pieces)
}
func (ap arrayPieces) Swap(i, j int) {
	ap.pieces[j], ap.pieces[i] = ap.pieces[i], ap.pieces[j]
}
func (ap arrayPieces) Less(i, j int) bool {
	v1, ok1 := ap.pmap[ap.pieces[i].Letter]
	if ok1 {
		v2, ok2 := ap.pmap[ap.pieces[j].Letter]
		if ok2 {
			return v2 < v1
		}
	}
	return false
}
func priority(pieces []chess.Piece, pmap map[byte]map[Pos]ThreateningVector) []chess.Piece {
	priorityMap := make(map[string]int)
	for k, e := range pmap {
		priorityMap[strconv.Itoa(int(k))] = len(e)
	}
	aps := arrayPieces{
		pieces: pieces,
		pmap:   priorityMap,
	}
	sort.Sort(aps)
	return aps.pieces
}
func threatPosInResult(tVector ThreateningVector, results []chess.ResultPosition) bool {
	for _, v := range tVector {
		for _, r := range results {
			if v.X == r.X && v.Y == r.Y {
				return true
			}
		}
	}
	return false
}

func threatPos(tVector ThreateningVector, x int, y int) bool {
	for _, rPos := range tVector {
		if rPos.X == x && rPos.Y == y {
			return true
		}
	}
	return false
}
func searchResult(results []chess.ResultPosition, x int, y int) bool {
	for _, rPos := range results {
		if rPos.X == x && rPos.Y == y {
			return true
		}
	}
	return false
}

func (s *solver) setBoard(b *Board) {
	s.board = b
}

func (s *solver) threateningForPType(pieces []chess.Piece) map[byte]map[Pos]ThreateningVector {
	retMap := make(map[byte]map[Pos]ThreateningVector)
	for _, p := range pieces {
		pt := p.Letter[0]
		mp := s.board.threateningForPType(pt)
		retMap[pt] = mp
	}
	return retMap
}
