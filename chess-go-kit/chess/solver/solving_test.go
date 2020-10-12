package solver

import (
	"../../chess"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

var _ = Describe("Solving", func() {
	var s Solver
	pieces := []chess.Piece{{Letter: "K", Npieces: 1}, {Letter: "B", Npieces: 1}, {Letter: "R", Npieces: 1},
		{Letter: "Q", Npieces: 1}, {Letter: "N", Npieces: 1}}
	var testData map[byte]map[Pos]ThreateningVector
	BeforeSuite(func() {
		s = NewSolver()
		s.setBoard(&Board{Dimension: 8})
		testData = s.threateningForPType(pieces)
	})
	Describe("Threatenings", func() {
		Context("For all pieces", func() {
			It("Return a correct size map", func() {
				for _, p := range pieces {
					Expect(len(testData[p.Letter[0]])).Should(Equal(64))
				}
			})
		})
	})
	Describe("Solving", func() {
		Context("Simple 8 Queens", func() {
			It("Should return correct number of results", func() {
				problem := chess.Problem{
					Dim: "8x8",
					Pieces: []chess.Piece{{
						Letter:  "Q",
						Npieces: 8,
					}},
				}
				solution, _ := s.Solve(problem)
				Expect(len(solution)).Should(Equal(92))
			})
		})
	})

})

func TestSolving(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SolverTest Suite")
}
