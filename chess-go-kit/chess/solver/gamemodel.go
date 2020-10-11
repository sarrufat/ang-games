package solver

type BoardDim = [2]int

type PieceParam struct {
	Piece  byte
	Number int
}
type Pos struct {
	X, Y int
}

type Positions = []Pos
type ThreateningVector = Positions

type ChessParam = []PieceParam

type Board struct {
	Dimension BoardDim
	pMap      map[Pos]PieceBase
}
type IBoard interface {
	isInside(x, y int) bool
}
type PieceBase struct {
	Threatening
	Position Pos
	Board    *Board
}

type King struct {
	PieceBase
}
type Knight struct {
	PieceBase
}
type Bishop struct {
	PieceBase
}
type Rook struct {
	PieceBase
}
type Queen struct {
	PieceBase
}

type Threatening interface {
	Threatening() ThreateningVector
}

type DirectionalMove interface {
}

func (board *Board) isInside(x, y int) bool {
	return x >= 0 && y >= 0 && x < board.Dimension[0] && y < board.Dimension[1]
}

func (k *King) Threatening() ThreateningVector {
	var tres ThreateningVector
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x != 0 || y != 0 {
				nx := k.Position.X + x
				ny := k.Position.Y + y
				if k.Board.isInside(nx, ny) {
					tres = append(tres, Pos{X: nx, Y: ny})
				}
			}

		}
	}
	return tres
}

func (k *Knight) Threatening() ThreateningVector {
	var tres ThreateningVector
	for _, mv := range knightMVectors {
		nx := k.Position.X + mv.X
		ny := k.Position.Y + mv.Y
		if k.Board.isInside(nx, ny) {
			tres = append(tres, Pos{X: nx, Y: ny})
		}
	}
	return tres
}

func (b *Bishop) Threatening() ThreateningVector {
	var tres ThreateningVector
	check := func(p Pos) bool {
		return b.Board.isInside(p.X, p.Y)
	}
	northEst := func() Positions {
		return vectorIncr(b.Position, Pos{X: 1, Y: -1}, check)
	}
	northWest := func() Positions {
		return vectorIncr(b.Position, Pos{X: -1, Y: -1}, check)
	}
	southWest := func() Positions {
		return vectorIncr(b.Position, Pos{X: -1, Y: 1}, check)
	}
	southEst := func() Positions {
		return vectorIncr(b.Position, Pos{X: 1, Y: 1}, check)
	}

	tres = append(tres, northEst()...)
	tres = append(tres, northWest()...)
	tres = append(tres, southWest()...)
	tres = append(tres, southEst()...)
	return tres
}

func (q *Queen) Threatening() ThreateningVector {
	var tres ThreateningVector
	asBishop := Bishop{PieceBase: PieceBase{Position: q.Position, Board: q.Board}}
	asRook := Rook{PieceBase: PieceBase{Position: q.Position, Board: q.Board}}
	tres = append(tres, asBishop.Threatening()...)
	tres = append(tres, asRook.Threatening()...)
	return tres
}

func (b *Rook) Threatening() ThreateningVector {
	var tres ThreateningVector
	check := func(p Pos) bool {
		return b.Board.isInside(p.X, p.Y)
	}
	north := func() Positions {
		return vectorIncr(b.Position, Pos{X: 0, Y: -1}, check)
	}
	south := func() Positions {
		return vectorIncr(b.Position, Pos{X: 0, Y: 1}, check)
	}
	west := func() Positions {
		return vectorIncr(b.Position, Pos{X: -1, Y: 0}, check)
	}

	est := func() Positions {
		return vectorIncr(b.Position, Pos{X: 1, Y: 0}, check)
	}

	tres = append(tres, north()...)
	tres = append(tres, south()...)
	tres = append(tres, west()...)
	tres = append(tres, est()...)

	return tres
}

func vectorIncr(pos, ipos Pos, check func(Pos) bool) Positions {
	var rPos Positions
	newPos := Pos{X: pos.X + ipos.X, Y: pos.Y + ipos.Y}
	if check(newPos) {
		rPos = append(rPos, newPos)
		rPos = append(rPos, vectorIncr(newPos, ipos, check)...)
	}
	return rPos
}

var (
	// (1, -2), (2, -1), (2, 1), (1, 2), (-1, 2), (-2, 1), (-2, -1), (-1, -2)
	knightMVectors = []Pos{Pos{X: 1, Y: -2}, {X: 2, Y: -1}, {X: 2, Y: 1}, {X: 1, Y: 2}, {X: -1, Y: 2}, {X: -2, Y: 1},
		{X: -2, Y: -1}, {X: -1, Y: -2}}
)
