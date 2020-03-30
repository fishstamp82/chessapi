package chess

import (
	"errors"
	"fmt"
	"strconv"
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

func NewPBoard() PlayBoard {

	startingFen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	var err error
	splitted := strings.Split(startingFen, " ")
	board := splitted[0]
	turn := splitted[1]
	castle := splitted[2]
	enPassant := splitted[3]
	halfMove := splitted[4]
	fullMove := splitted[5]
	ranks := strings.Split(board, "/")

	finalBoard := map[Square]Piece{}
	var i, j, row, col, toSkip int
	var boardIdx Square
	for i = 0; i < len(ranks); i++ {
		row = 7 - i
		col = 0
		for j = 0; j < len(ranks[i]); j++ {
			boardIdx = Square(row*8 + col)
			switch piece := fenToPiece[ranks[i][j]]; {
			case piece == Empty:
				toSkip, _ = strconv.Atoi(ranks[i][j : j+1])
				col += toSkip
			default:
				finalBoard[boardIdx] = piece
				col += 1
			}
		}
	}
	eb := &PBoard{
		Context: Context{
			State:               Playing,
			PlayersTurn:         White,
			enPassantSquare:     none,
			whiteCanCastleLeft:  true,
			whiteCanCastleRight: true,
			blackCanCastleRight: true,
			blackCanCastleLeft:  true,
			fullMove:            1,
		},
	}
	for key, val := range finalBoard {
		eb.board[key] = val
	}
	switch turn {
	case "w":
		eb.Context.PlayersTurn = White
	case "b":
		eb.Context.PlayersTurn = Black
	}

	eb.Context.whiteCanCastleLeft = false
	eb.Context.whiteCanCastleRight = false
	eb.Context.blackCanCastleRight = false
	eb.Context.blackCanCastleLeft = false
	for _, b := range castle {
		switch b {
		case 'K':
			eb.Context.whiteCanCastleRight = true
		case 'Q':
			eb.Context.whiteCanCastleLeft = true
		case 'k':
			eb.Context.blackCanCastleRight = true
		case 'q':
			eb.Context.blackCanCastleLeft = true
		}
	}

	switch sq := enPassant; {
	case sq == "-":
		eb.Context.enPassantSquare = none
	default:
		eb.Context.enPassantSquare = stringToSquare[sq]
	}

	var halfMoveInt, fullMoveInt int
	halfMoveInt, err = strconv.Atoi(halfMove)
	if err != nil {
		panic(err)
	}
	eb.Context.halfMove = halfMoveInt
	fullMoveInt, err = strconv.Atoi(fullMove)
	if err != nil {
		panic(err)
	}
	eb.Context.fullMove = fullMoveInt

	return eb
}
