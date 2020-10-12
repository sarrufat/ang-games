package solver

import (
	"../../chess"
	// "fmt"
	"strconv"
)

type Solver interface {
	threateningForPType([]chess.Piece) map[byte]map[Pos]ThreateningVector
	Solve(p chess.Problem) ([][]chess.ResultPosition, error)
	setBoard(b *Board)
}

type solver struct {
	board *Board
}

func NewSolver() Solver {
	return &solver{}
}

type anonRec func(chess.Piece, []chess.Piece, ThreateningVector, []chess.ResultPosition)

func (s *solver) Solve(p chess.Problem) ([][]chess.ResultPosition, error) {
	dim, err := strconv.Atoi(p.Dim[0:1])
	if err != nil {
		return nil, err
	}
	s.board = &Board{
		Dimension: dim,
	}
	tMap := s.threateningForPType(p.Pieces)
	var iterations int32
	var recSolve anonRec
	var results [][]chess.ResultPosition

	recSolve = func(actual chess.Piece, reaming []chess.Piece, tVector ThreateningVector, resPos []chess.ResultPosition) {
		iterations += 1
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
						//			fmt.Printf("OK %d (%d, %d) = respos %v, pieces = %v\n", iterations, x, y, resPos, reaming)
						if len(reaming) == 0 { // solution found
							results = append(results, actualRPos)
							//				fmt.Println(actualRPos)
							// TODO: Limit max results
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
	fPieces := flatten(p.Pieces)
	//	for x := 0; x < dim; x++ {
	//for y := 0; y < dim; y++ {
	// resPos := []chess.ResultPosition{{Piece: fPieces[0].Letter, X: x, Y: x}}
	recSolve(fPieces[0], fPieces[1:], ThreateningVector{}, []chess.ResultPosition{})
	//	}
	//	}

	return results, nil
}
func flatten(pieces []chess.Piece) []chess.Piece {
	var out []chess.Piece
	for _, p := range pieces {
		for i := 0; i < p.Npieces; i++ {
			out = append(out, p)
		}
	}
	return out
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
