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

		kings   []PieceBase
		knights []PieceBase
		bishops []PieceBase
		rooks   []PieceBase
		queens  []PieceBase
	)
	BeforeSuite(func() {
		board = Board{Dimension: 8}
		testPositions := []Pos{{}, {X: 7}, {Y: 7}, {X: 7, Y: 7}}
		builder := NewPieceBuilder()
		for _, pos := range testPositions {
			kings = append(kings, builder.Build(CKing, pos, &board))
			knights = append(knights, builder.Build(CKnight, pos, &board))
			bishops = append(bishops, builder.Build(CBishop, pos, &board))
			rooks = append(rooks, builder.Build(CRook, pos, &board))
			queens = append(queens, builder.Build(CQueen, pos, &board))

		}
	})

	Describe("PieceBase", func() {
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

func checkIt(ts []PieceBase, expeted int) {
	for _, k := range ts {
		Expect(len(k.Threatening())).Should(Equal(expeted))
	}
}

func TestSolverTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SolverTest Suite")
}
