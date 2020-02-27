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
	board   [64]Piece
	turn    Player
	inCheck bool
}

func (b *Board) InCheck() bool {
	return b.inCheck
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
	if !inSquares(sq2, availMoves) {
		return errors.New(fmt.Sprintf("%s can't go to %s\n", pieceToString[b.board[sq1]], t))
	}

	ourKingPos := b.kingSquare(b.turn)
	// Make Move
	p := b.board[sq1]
	b.board[sq1] = 0
	b.board[sq2] = p

	//If we are in check, revert the move and return error can't move
	for _, oppPiece := range b.pieces(b.opponent()) {
		if inSquares(ourKingPos, b.targets(oppPiece)) {
			b.board[sq1] = p
			b.board[sq2] = 0
			return errors.New("king will be in check")
		}
	}

	//b.is_check()
	//Calculate if we check the opponent, and update inCheck

	// Make other players turn
	b.switchTurn()
	return nil
}

func (b *Board) opponent() Player {
	switch b.turn {
	case White:
		return Black
	case Black:
		return White
	}

	panic("must be black or white")
}

func (b *Board) pieces(p Player) []Square {
	var isWhite bool
	switch p {
	case White:
		isWhite = true
	case Black:
		isWhite = false
	}

	var pieces []Square
	for pos := a1; pos <= h8; pos += 1 {
		if b.board[pos] > 0 && isWhite {
			pieces = append(pieces, pos)
		} else if b.board[pos] < 0 && !isWhite {
			pieces = append(pieces, pos)

		}

	}
	return pieces
}

func (b *Board) kingSquare(p Player) Square {
	var king Piece
	switch p {
	case White:
		king = WhiteKing
	case Black:
		king = BlackKing
	}

	for pos := a1; pos <= h8; pos += 1 {
		if b.board[pos] == king {
			return pos
		}
	}

	panic("must have a white King")
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
	case WhiteKnight:
		moves = b.knightMoves(s)
	case BlackKnight:
		moves = b.knightMoves(s)
	case WhiteRook:
		moves = b.rookMoves(s)
	case BlackRook:
		moves = b.rookMoves(s)
	case WhiteQueen:
		moves = b.queenMoves(s)
	case BlackQueen:
		moves = b.queenMoves(s)
	case WhiteKing:
		moves = b.kingMoves(s)
	case BlackKing:
		moves = b.kingMoves(s)
	}
	return moves
}

func (b *Board) targets(s Square) []Square {
	p := b.board[s]
	var targets []Square
	switch p {
	case WhitePawn:
		targets = b.pawnTargets(s)
	case BlackPawn:
		targets = b.pawnTargets(s)
	case WhiteBishop:
		targets = b.bishopTargets(s)
	case BlackBishop:
		targets = b.bishopTargets(s)
	case WhiteKnight:
		targets = b.knightTargets(s)
	case BlackKnight:
		targets = b.knightTargets(s)
	case WhiteRook:
		targets = b.rookTargets(s)
	case BlackRook:
		targets = b.rookTargets(s)
	case WhiteQueen:
		targets = b.queenTargets(s)
	case BlackQueen:
		targets = b.queenTargets(s)
	case WhiteKing:
		targets = b.kingTargets(s)
	case BlackKing:
		targets = b.kingTargets(s)
	}
	return targets
}
