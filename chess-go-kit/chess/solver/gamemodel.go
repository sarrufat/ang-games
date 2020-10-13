package solver

type BoardDim = int

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
}

type IBoard interface {
	isInside(x, y int) bool
	threateningForPType(byte) map[Pos]ThreateningVector
}
type pieceBase struct {
	PieceBase
	cType    byte
	Position Pos
	Board    *Board
}

type King struct {
	pieceBase
}
type Knight struct {
	pieceBase
}
type Bishop struct {
	pieceBase
}
type Rook struct {
	pieceBase
}
type Queen struct {
	pieceBase
}

type PieceBase interface {
	Threatening() ThreateningVector
	SetPos(pos Pos)
}

type PieceBuilder interface {
	Build(pt byte, pos Pos, board *Board) PieceBase
}

type pieceBuilder struct {
}

func (builder *pieceBuilder) Build(pt byte, pos Pos, board *Board) PieceBase {
	var newPiece PieceBase
	switch pt {
	case CKing:
		newPiece = &King{pieceBase: pieceBase{
			cType:    CKing,
			Position: pos,
			Board:    board,
		}}
	case CBishop:
		newPiece = &Bishop{pieceBase: pieceBase{
			cType:    CBishop,
			Position: pos,
			Board:    board,
		}}
	case CKnight:
		newPiece = &Knight{pieceBase: pieceBase{
			cType:    CKnight,
			Position: pos,
			Board:    board,
		}}
	case CQueen:
		newPiece = &Queen{pieceBase: pieceBase{
			cType:    CQueen,
			Position: pos,
			Board:    board,
		}}
	case CRook:
		newPiece = &Rook{pieceBase: pieceBase{
			cType:    CRook,
			Position: pos,
			Board:    board,
		}}
	}
	return newPiece
}

func NewPieceBuilder() PieceBuilder {
	return &pieceBuilder{}
}

func (board *Board) isInside(x, y int) bool {
	return x >= 0 && y >= 0 && x < board.Dimension && y < board.Dimension
}

func (board *Board) threateningForPType(pt byte) map[Pos]ThreateningVector {
	builder := NewPieceBuilder()
	piece := builder.Build(pt, Pos{}, board)
	rMap := make(map[Pos]ThreateningVector)
	for x := 0; x < board.Dimension; x++ {
		for y := 0; y < board.Dimension; y++ {
			pos := Pos{X: x, Y: y}
			piece.SetPos(pos)
			rMap[pos] = piece.Threatening()
		}
	}
	return rMap
}

func (p *pieceBase) SetPos(pos Pos) {
	p.Position.X = pos.X
	p.Position.Y = pos.Y
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
	asBishop := Bishop{pieceBase: pieceBase{Position: q.Position, Board: q.Board}}
	asRook := Rook{pieceBase: pieceBase{Position: q.Position, Board: q.Board}}
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

func vectorIncr(pos, iPos Pos, check func(Pos) bool) Positions {
	var rPos Positions
	newPos := Pos{X: pos.X + iPos.X, Y: pos.Y + iPos.Y}
	if check(newPos) {
		rPos = append(rPos, newPos)
		rPos = append(rPos, vectorIncr(newPos, iPos, check)...)
	}
	return rPos
}

const (
	CKing   byte = 'K'
	CBishop byte = 'B'
	CRook   byte = 'R'
	CQueen  byte = 'Q'
	CKnight byte = 'N'
)

var (
	// (1, -2), (2, -1), (2, 1), (1, 2), (-1, 2), (-2, 1), (-2, -1), (-1, -2)
	knightMVectors = []Pos{{X: 1, Y: -2}, {X: 2, Y: -1}, {X: 2, Y: 1}, {X: 1, Y: 2}, {X: -1, Y: 2}, {X: -2, Y: 1},
		{X: -2, Y: -1}, {X: -1, Y: -2}}
)
