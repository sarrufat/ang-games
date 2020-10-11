package solver_test_test

import (
	. "../../solver"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

var _ = Describe("Solver", func() {
	var (
		board Board

		kings   []Threatening
		knights []Threatening
		bishops []Threatening
		rooks   []Threatening
		queens  []Threatening
	)
	BeforeSuite(func() {
		board = Board{Dimension: [2]int{8, 8}}
		testPositions := []Pos{{}, {7, 0}, {0, 7}, {7, 7}}
		for _, pos := range testPositions {
			kings = append(kings, &King{PieceBase: PieceBase{Board: &board, Position: pos}})
			knights = append(knights, &Knight{PieceBase: PieceBase{Board: &board, Position: pos}})
			bishops = append(bishops, &Bishop{PieceBase: PieceBase{Board: &board, Position: pos}})
			rooks = append(rooks, &Rook{PieceBase: PieceBase{Position: pos, Board: &board}})
			queens = append(queens, &Queen{PieceBase: PieceBase{Position: pos, Board: &board}})
		}
	})

	Describe("Threatening", func() {
		Context("King With different board positions", func() {
			It("Should report correct lenght on the corners", func() {
				checkIt(kings, 3)
			})

		})
		Context("Knight With different board positions", func() {
			It("Should report correct lenght on the corners", func() {
				checkIt(knights, 2)
			})
		})
		Context("Bichop With different board positions", func() {
			It("Should report correct lenght on the corners", func() {
				checkIt(bishops, 7)
			})
		})
		Context("Rook With different board positions", func() {
			It("Should report correct lenght on the corners", func() {
				checkIt(rooks, 14)
			})
		})
		Context("Queen With different board positions", func() {
			It("Should report correct lenght on the corners", func() {
				checkIt(queens, 21)
			})
		})
	})
})

func checkIt(ts []Threatening, expeted int) {
	for _, k := range ts {
		Expect(len(k.Threatening())).Should(Equal(expeted))
	}
}

func TestSolverTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SolverTest Suite")
}
