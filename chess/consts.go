package chess

import "fmt"

//go:generate stringer -type=Piece,Player,Square,State,MovementType -output=stringer_gen.go

type Player byte
type Square int
type State byte
type MovementType byte
type Piece int8

const (
	BlackKing   Piece = iota - 6
	BlackQueen        // -5
	BlackRook         // -4
	BlackBishop       // -3
	BlackKnight       // -2
	BlackPawn         // -1
	Empty             // 0
	WhitePawn         // 1
	WhiteKnight       // 2
	WhiteBishop       // 3
	WhiteRook         // 4
	WhiteQueen        // 5
	WhiteKing         // 6
	Pawn              // Used for parsing PGN
	Knight
	Bishop
	Rook
	Queen
	King
)

type NoMoveError struct {
	Move string
}

func (nm *NoMoveError) Error() string {
	return fmt.Sprintf("no move: %s", nm.Move)
}

const (
	Noone Player = iota
	White        // 1
	Black        // 2
	Both         // 2
)

const (
	Playing State = iota + 1
	Check
	CheckMate
	Draw
	Promo
)

const (
	Regular MovementType = iota
	PawnMove
	Capture
	Castle
	Promotion
	CapturePromotion
	CaptureEnPassant
	CheckMove
)

const (
	none Square = iota - 1
	a1
	b1
	c1
	d1
	e1
	f1
	g1
	h1
	a2
	b2
	c2
	d2
	e2
	f2
	g2
	h2
	a3
	b3
	c3
	d3
	e3
	f3
	g3
	h3
	a4
	b4
	c4
	d4
	e4
	f4
	g4
	h4
	a5
	b5
	c5
	d5
	e5
	f5
	g5
	h5
	a6
	b6
	c6
	d6
	e6
	f6
	g6
	h6
	a7
	b7
	c7
	d7
	e7
	f7
	g7
	h7
	a8
	b8
	c8
	d8
	e8
	f8
	g8
	h8
)

const (
	None Square = iota - 1
	A1
	B1
	C1
	D1
	E1
	F1
	G1
	H1
	A2
	B2
	C2
	D2
	E2
	F2
	G2
	H2
	A3
	B3
	C3
	D3
	E3
	F3
	G3
	H3
	A4
	B4
	C4
	D4
	E4
	F4
	G4
	H4
	A5
	B5
	C5
	D5
	E5
	F5
	G5
	H5
	A6
	B6
	C6
	D6
	E6
	F6
	G6
	H6
	A7
	B7
	C7
	D7
	E7
	F7
	G7
	H7
	A8
	B8
	C8
	D8
	E8
	F8
	G8
	H8
)

var stringToSquare = map[string]Square{
	"a1": a1,
	"b1": b1,
	"c1": c1,
	"d1": d1,
	"e1": e1,
	"f1": f1,
	"g1": g1,
	"h1": h1,
	"a2": a2,
	"b2": b2,
	"c2": c2,
	"d2": d2,
	"e2": e2,
	"f2": f2,
	"g2": g2,
	"h2": h2,
	"a3": a3,
	"b3": b3,
	"c3": c3,
	"d3": d3,
	"e3": e3,
	"f3": f3,
	"g3": g3,
	"h3": h3,
	"a4": a4,
	"b4": b4,
	"c4": c4,
	"d4": d4,
	"e4": e4,
	"f4": f4,
	"g4": g4,
	"h4": h4,
	"a5": a5,
	"b5": b5,
	"c5": c5,
	"d5": d5,
	"e5": e5,
	"f5": f5,
	"g5": g5,
	"h5": h5,
	"a6": a6,
	"b6": b6,
	"c6": c6,
	"d6": d6,
	"e6": e6,
	"f6": f6,
	"g6": g6,
	"h6": h6,
	"a7": a7,
	"b7": b7,
	"c7": c7,
	"d7": d7,
	"e7": e7,
	"f7": f7,
	"g7": g7,
	"h7": h7,
	"a8": a8,
	"b8": b8,
	"c8": c8,
	"d8": d8,
	"e8": e8,
	"f8": f8,
	"g8": g8,
	"h8": h8,
}

var squareToString = map[Square]string{
	a1: "a1",
	b1: "b1",
	c1: "c1",
	d1: "d1",
	e1: "e1",
	f1: "f1",
	g1: "g1",
	h1: "h1",
	a2: "a2",
	b2: "b2",
	c2: "c2",
	d2: "d2",
	e2: "e2",
	f2: "f2",
	g2: "g2",
	h2: "h2",
	a3: "a3",
	b3: "b3",
	c3: "c3",
	d3: "d3",
	e3: "e3",
	f3: "f3",
	g3: "g3",
	h3: "h3",
	a4: "a4",
	b4: "b4",
	c4: "c4",
	d4: "d4",
	e4: "e4",
	f4: "f4",
	g4: "g4",
	h4: "h4",
	a5: "a5",
	b5: "b5",
	c5: "c5",
	d5: "d5",
	e5: "e5",
	f5: "f5",
	g5: "g5",
	h5: "h5",
	a6: "a6",
	b6: "b6",
	c6: "c6",
	d6: "d6",
	e6: "e6",
	f6: "f6",
	g6: "g6",
	h6: "h6",
	a7: "a7",
	b7: "b7",
	c7: "c7",
	d7: "d7",
	e7: "e7",
	f7: "f7",
	g7: "g7",
	h7: "h7",
	a8: "a8",
	b8: "b8",
	c8: "c8",
	d8: "d8",
	e8: "e8",
	f8: "f8",
	g8: "g8",
	h8: "h8",
}
