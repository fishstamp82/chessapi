package chess

import (
	"errors"
	"fmt"
	"strings"
)

type PlayBoard interface {
	Move(string) (Context, error) //Move piece from e2 to e4, error if not possible
	Board() map[Square]Piece      //For human visualization of board
}

type PBoard struct {
	board   [64]Piece
	moves   []Move
	Context Context
}

//CLI repr of board
func (b *PBoard) Board() map[Square]Piece {
	board := map[Square]Piece{}
	for square := a1; square <= h8; square++ {
		board[square] = b.board[square]
	}

	return board
}

// Move gets squares in human readable form, and performs a move
// error is nil on successful move
// arguments are two squares : "e2e4"
func (b *PBoard) Move(moveStr string) (Context, error) {
	if b.Context.State != Playing && b.Context.State != Check {
		return b.Context, errors.New("not in playing state")
	}
	fromSquare, toSquare, err := getSquare(moveStr)
	if err != nil {
		return b.Context, err
	}

	return b.move(fromSquare, toSquare)
}

func (b *PBoard) move(fromSquare, toSquare Square) (Context, error) {
	var opponent Player
	switch b.Context.PlayersTurn {
	case White:
		if b.board[fromSquare] < 0 {
			return b.Context, errors.New("white's turn")
		}
		opponent = Black
	case Black:
		if b.board[fromSquare] > 0 {
			return b.Context, errors.New("black's turn")
		}
		opponent = White
	}

	availMoves := validMovesForSquare(fromSquare, b.board, b.Context)

	availSquares := getSquares(availMoves)
	if !inSquares(toSquare, availSquares) {
		return b.Context, errors.New(fmt.Sprintf("%s can't go to %s\n", b.board[fromSquare], squareToString[toSquare]))
	}

	//todo: replace with function thate uses chess algebraic notation
	m := Move{}
	for _, move := range availMoves {
		if move.fromSquare == fromSquare && move.toSquare == toSquare {
			m = move
		}
	}
	if m.toSquare == none {
		return b.Context, &NoMoveError{Move: strings.Join([]string{fromSquare.String(), toSquare.String()}, "")}
	}

	b.board = makeMove(m, b.board)

	opponentsKing := getKingSquareMust(opponent, b.board)
	if inCheck(opponentsKing, b.board) {
		b.Context.State = Check
	} else {
		b.Context.State = Playing
	}

	if isCheckMated(opponentsKing, b.board) {
		b.Context.State = CheckMate
		b.Context.Winner = b.Context.PlayersTurn
		return b.Context, nil
	}

	if isDraw(opponent, b.board, b.Context) {
		b.Context.State = Draw
		b.Context.Winner = Both
		return b.Context, nil
	}

	updateCastling(m, b.Context)
	b.Context.enPassantSquare = getEnPassantSquare(m)

	//Increment full move if this was blacks move
	if b.Context.PlayersTurn == Black {
		b.Context.fullMove += 1
	}

	//Increment half move if this was not a pawn move and not a capture
	isPawnMove, isCaptureMove := false, false
	for _, moveType := range m.moveTypes {
		if moveType == PawnMove {
			isPawnMove = true
		}
		if moveType == Capture {
			isCaptureMove = true
		}
	}
	if isPawnMove || isCaptureMove {
		b.Context.halfMove = 0
	} else {
		b.Context.halfMove += 1
	}

	// Switch to other player
	b.Context.PlayersTurn = switchTurn(b.Context.PlayersTurn)
	return b.Context, nil
}

func NewPBoard(fen string) PlayBoard {
	return &PBoard{}
}
