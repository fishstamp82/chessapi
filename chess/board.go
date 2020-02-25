package chess

import (
	"errors"
	"fmt"
)

type Player byte
type Square int

// row goes from 0 to 7
func (s Square) row() Square {
	return s / 8
}

// col goes from 0 to 7
func (s Square) col() Square {
	return s % 8
}

const (
	a1 Square = iota
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

type Board struct {
	board [64]Piece
	turn  Player
}

func (b *Board) PlayersTurn() string {
	var p Player
	p = b.turn
	if p == White {
		return "white"
	} else if p == Black {
		return "black"
	}
	panic("neither white nor blacks turn")
}

// Given human readable string input "e2", return string
// repr of piece. if none, return "-"
func (b *Board) Get(s string) (string, error) {
	sq, err := b.getSquare(s)
	if err != nil {
		return "", errors.New("bad move input, good format should be: 'e2', 'd3', etc")
	}
	p := b.board[sq]
	return pieceToString[p], nil
}

//CLI repr of board
func (b *Board) CliStrRepr() string {
	var board string
	for row := 7; row >= 0; row-- {
		board += "\n-----------------\n|"
		for col := 0; col <= 7; col++ {
			idx := row*8 + col
			board += pieceToUnicode[b.board[idx]] + "|"
		}
	}

	return board
}

// Move gets squares in human readable form, and performs a move
// error is nil on successful move
// arguments are algebraic chess notation 'e2' -> 'e4'
func (b *Board) Move(s, t string) error {
	sq1, err := b.getSquare(s)
	if err != nil {
		return errors.New("bad move input, good format should be: 'e2', 'd3', etc")
	}
	sq2, err := b.getSquare(t)
	if err != nil {
		return errors.New("bad move input, good format should be: 'e2', 'd3', etc")
	}

	availMoves := b.moves(sq1)
	if !inSlice(sq2, availMoves) {
		return errors.New(fmt.Sprintf("%s can't go to %s\n", pieceToString[b.board[sq1]], t))
	}
	for i := 0; i < len(availMoves); i++ {

	}
	// Mave Move
	p := b.board[sq1]
	b.board[sq1] = 0
	b.board[sq2] = p

	// Make other players turn
	b.switchTurn()
	return nil
}

func NewBoard() *Board {
	b := &Board{}

	b.turn = White
	//Pawns
	for _, s := range []Square{a2, b2, c2, d2, e2, f2, g2, h2} {
		b.board[s] = WhitePawn
	}
	for _, s := range []Square{a7, b7, c7, d7, e7, f7, g7, h7} {
		b.board[s] = BlackPawn
	}
	b.board[a1] = WhiteRook
	b.board[h1] = WhiteRook

	b.board[a8] = BlackRook
	b.board[h8] = BlackRook

	b.board[b1] = WhiteKnight
	b.board[g1] = WhiteKnight

	b.board[b8] = BlackKnight
	b.board[g8] = BlackKnight

	b.board[c1] = WhiteBishop
	b.board[f1] = WhiteBishop

	b.board[c8] = BlackBishop
	b.board[f8] = BlackBishop

	b.board[d1] = WhiteQueen
	b.board[e1] = WhiteKing

	b.board[d8] = BlackQueen
	b.board[e8] = BlackKing

	return b
}

func NewEmptyBoard() *Board {

	b := &Board{}
	return b

}
func (b *Board) getSquare(s string) (Square, error) {
	if len(s) != 2 {
		return 0, errors.New("wrong length")
	}
	sq, found := stringToSquare[s]
	if !found {
		return 0, errors.New(fmt.Sprintf("no such square: %s", s))
	}
	return sq, nil
}

func (b *Board) switchTurn() {
	if b.turn == White {
		b.turn = Black
	} else {
		b.turn = White
	}

}

func (b *Board) whitePawnMoves(s Square) []Square {
	var moves []Square
	var t Square
	var first, second Square
	col := s.col()
	row := s.row()
	pos := row*8 + col
	if row == 1 {
		first = pos + 8   // one square move
		second = pos + 16 // two square move
		if b.board[first] == 0 {
			moves = append(moves, first)
		}
		if b.board[second] == 0 && b.board[first] == 0 {
			moves = append(moves, second)
		}
	} else {
		first = pos + 8 // one square move
		if b.board[first] == 0 {
			moves = append(moves, first)
		}
	}

	upperRight := func(s []Square) []Square {
		t = pos + 9 // attack upper right
		if b.board[t] < 0 {
			s = append(s, t)
		}
		return s
	}

	upperLeft := func(s []Square) []Square {
		t = pos + 7 // attack upper right
		if b.board[t] < 0 {
			s = append(s, t)
		}
		return s
	}
	if col == 0 {
		moves = upperRight(moves)
	} else if col == 7 {
		moves = upperLeft(moves)
	} else {
		moves = upperRight(moves)
		moves = upperLeft(moves)
	}
	return moves
}

func (b *Board) blackPawnMoves(s Square) []Square {
	var moves []Square
	var t Square
	var first, second Square
	col := s.col()
	row := s.row()
	pos := row*8 + col
	if row == 6 {
		first = pos - 8   // one square move
		second = pos - 16 // two square move
		if b.board[first] == 0 {
			moves = append(moves, first)
		}
		if b.board[second] == 0 && b.board[first] == 0 {
			moves = append(moves, second)
		}
	} else {
		first = pos - 8 // one square move
		if b.board[first] == 0 {
			moves = append(moves, first)
		}
	}

	lowerRight := func(s []Square) []Square {
		t = pos - 7 // attack lower right
		if b.board[t] > 0 {
			s = append(s, t)
		}
		return s
	}

	lowerLeft := func(s []Square) []Square {
		t = pos - 9 // attack lower left
		if b.board[t] > 0 {
			s = append(s, t)
		}
		return s
	}
	if col == 0 {
		moves = lowerRight(moves)
	} else if col == 7 {
		moves = lowerLeft(moves)
	} else {
		moves = lowerRight(moves)
		moves = lowerLeft(moves)
	}
	return moves
}

func (b *Board) moves(s Square) []Square {
	p := b.board[s]
	var moves []Square
	switch p {
	case WhitePawn:
		moves = b.whitePawnMoves(s)
	}
	return moves
}
