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
		return errors.New(fmt.Sprintf("bad move input, got: %v, good format: 'e2'", s))
	}
	sq2, err := b.getSquare(t)
	if err != nil {
		return errors.New(fmt.Sprintf("bad second input, got: %v, good format: 'e4'", t))
	}

	switch b.turn {
	case White:
		if b.board[sq1] < 0 {
			return errors.New("white's turn")
		}
	case Black:
		if b.board[sq1] > 0 {
			return errors.New("black's turn")
		}
	}

	availMoves := b.moves(sq1)
	if !inSlice(sq2, availMoves) {
		return errors.New(fmt.Sprintf("%s can't go to %s\n", pieceToString[b.board[sq1]], t))
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

func (b *Board) moves(s Square) []Square {
	p := b.board[s]
	var moves []Square
	switch p {
	case WhitePawn:
		moves = b.whitePawnMoves(s)
	case BlackPawn:
		moves = b.blackPawnMoves(s)
	case WhiteBishop:
		moves = b.bishopMoves(s)
	case BlackBishop:
		moves = b.bishopMoves(s)
	}
	return moves
}
