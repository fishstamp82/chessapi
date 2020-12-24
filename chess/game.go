package chess

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

type Game struct {
	Board     *Board
	Context   Context
	Players   []Player
	moves     []Move
	startedAt int64
}

func (g *Game) Start() {
	g.startedAt = time.Now().UTC().UnixNano()
}

// Move gets squares in human readable form, and performs a move
// error is nil on successful move
// arguments are two squares : "e2e4"
func (game *Game) Move(moveStr string) (Context, error) {
	if game.Context.State != Playing && game.Context.State != Check {
		return game.Context, errors.New("not in playing state")
	}
	fromSquare, toSquare, err := game.Board.getSquare(moveStr)
	if err != nil {
		return game.Context, err
	}

	return game.move(fromSquare, toSquare)
}

// Move gets squares in human readable form, and performs a move
// error is nil on successful move
// arguments are two squares : "e2e4"
func (g *Game) MoveNotation(move Move) (Context, error) {
	if g.Context.State != Playing && g.Context.State != Check {
		return g.Context, errors.New("not in playing state")
	}
	fromSquare, toSquare := move.fromSquare, move.toSquare
	return g.move(fromSquare, toSquare)
}

func NewEmptyGame() *Game {
	b := &Board{}
	return &Game{
		Board: b,
		Context: Context{
			State:               Idle,
			ColorsTurn:          White,
			enPassantSquare:     none,
			whiteCanCastleLeft:  true,
			whiteCanCastleRight: true,
			blackCanCastleRight: true,
			blackCanCastleLeft:  true,
			fullMove:            1,
		},
	}
}

func NewGameFromFEN(fen string) *Game {
	var err error
	splitted := strings.Split(fen, " ")
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
	eb := NewEmptyGame()
	for key, val := range finalBoard {
		eb.Board.board[key] = val
	}
	switch turn {
	case "w":
		eb.Context.ColorsTurn = White
	case "b":
		eb.Context.ColorsTurn = Black
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

func GameFromPGN(reader io.Reader) *Game {
	g := NewGame()
	moves, err := pgnParse(reader)
	if err != nil {
		panic(err)
	}

	g.moves = moves

	return g
}

func (game *Game) FenString() string {
	var cnt int
	var board string
	var sq Square
	for i := 7; i >= 0; i-- {
		cnt = 0
		for j := 0; j < 8; j++ {
			sq = Square(i*8 + j)

			switch p := game.Board.board[sq]; {
			case p == Empty:
				cnt += 1
			default:
				if cnt > 0 {
					board += strconv.Itoa(cnt)
				}
				cnt = 0
				board += pieceToFen[p]
			}
			if j == 7 {
				if cnt == 0 {
					board += "/"
					continue
				}
				board += strconv.Itoa(cnt) + "/"
				cnt = 0
			}
		}
	}
	board = strings.TrimSuffix(board, "/")

	toMove := playerToFen[game.Context.ColorsTurn]

	var castle string
	if game.Context.whiteCanCastleRight {
		castle += pieceToFen[WhiteKing]
	}
	if game.Context.whiteCanCastleLeft {
		castle += pieceToFen[WhiteQueen]
	}
	if game.Context.blackCanCastleRight {
		castle += pieceToFen[BlackKing]
	}
	if game.Context.whiteCanCastleRight {
		castle += pieceToFen[BlackQueen]
	}

	if castle == "" {
		castle = "-"
	}

	var enpassant string
	if game.Context.enPassantSquare >= a1 {
		enpassant = game.Context.enPassantSquare.String()
	} else {
		enpassant = "-"
	}
	halfMove := strconv.Itoa(game.Context.halfMove)
	fullMove := strconv.Itoa(game.Context.fullMove)
	return fmt.Sprintf("%s %s %s %s %s %s", board, toMove, castle, enpassant, halfMove, fullMove)
}

func (g *Game) move(fromSquare, toSquare Square) (Context, error) {

	var opponent Color
	switch g.Context.ColorsTurn {
	case White:
		if g.Board.board[fromSquare] < 0 {
			return g.Context, errors.New("white's turn")
		}
		opponent = Black
	case Black:
		if g.Board.board[fromSquare] > 0 {
			return g.Context, errors.New("black's turn")
		}
		opponent = White
	}

	availMoves := validMovesForSquare(fromSquare, g.Board.board, g.Context)

	availSquares := getSquares(availMoves)
	if !inSquares(toSquare, availSquares) {
		return g.Context, fmt.Errorf("%s can't go to %s\n", g.Board.board[fromSquare], squareToString[toSquare])
	}

	//todo: replace with function thate uses chess algebraic notation
	m := Move{}
	for _, move := range availMoves {
		if move.fromSquare == fromSquare && move.toSquare == toSquare {
			m = move
		}
	}
	if m.toSquare == none {
		return g.Context, &NoMoveError{Move: strings.Join([]string{fromSquare.String(), toSquare.String()}, "")}
	}

	g.Board.board = makeMove(m, g.Board.board)

	opponentsKing := getKingSquareMust(opponent, g.Board.board)
	if inCheck(opponentsKing, g.Board.board) {
		g.Context.State = Check
	} else {
		g.Context.State = Playing
	}

	if isCheckMated(opponentsKing, g.Board.board) {
		g.Context.State = CheckMate
		g.Context.Winner = g.Context.ColorsTurn
		return g.Context, nil
	}

	if isDraw(opponent, g.Board.board, g.Context) {
		g.Context.State = Draw
		g.Context.Winner = Both
		return g.Context, nil
	}

	g.abortCastling(m)
	g.Context.enPassantSquare = g.getEnPassantSquare(m)

	//Increment full move if this was blacks move
	if g.Context.ColorsTurn == Black {
		g.Context.fullMove += 1
	}

	//Increment half move if this was not a pawn move and not a capture
	isPawnMove := false
	for _, moveType := range m.moveTypes {
		if moveType == PawnMove {
			isPawnMove = true
		}
	}
	if isPawnMove {
		g.Context.halfMove = 0
	} else {
		g.Context.halfMove += 1
	}

	// Switch to other player
	g.switchTurn()
	return g.Context, nil
}

func (game *Game) getEnPassantSquare(m Move) Square {
	if m.piece != WhitePawn && m.piece != BlackPawn {
		game.Context.enPassantSquare = none
		return none
	}
	if m.fromSquare.rank() == 2 && m.toSquare.rank() == 4 && m.piece == WhitePawn {
		return m.fromSquare + 8
	} else if m.fromSquare.rank() == 7 && m.toSquare.rank() == 5 && m.piece == BlackPawn {
		return m.fromSquare - 8
	} else {
		return none
	}
}

func (game *Game) switchTurn() {
	if game.Context.ColorsTurn == White {
		game.Context.ColorsTurn = Black
	} else {
		game.Context.ColorsTurn = White
	}
}

func (g *Game) abortCastling(m Move) {

	switch m.fromSquare {
	case a1:
		g.Context.whiteCanCastleLeft = false
	case h1:
		g.Context.whiteCanCastleRight = false
	case a8:
		g.Context.blackCanCastleLeft = false
	case h8:
		g.Context.blackCanCastleRight = false
	case e1:
		g.Context.whiteCanCastleLeft = false
		g.Context.whiteCanCastleRight = false
	case e8:
		g.Context.blackCanCastleLeft = false
		g.Context.blackCanCastleRight = false
	}

	switch m.toSquare {
	case a1:
		g.Context.whiteCanCastleLeft = false
	case h1:
		g.Context.whiteCanCastleRight = false
	case a8:
		g.Context.blackCanCastleLeft = false
	case h8:
		g.Context.blackCanCastleRight = false
	}

}

func NewGame() *Game {
	b := &Board{}

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

	return &Game{Board: b, Context: Context{
		State:               Idle,
		ColorsTurn:          White,
		Winner:              0,
		whiteCanCastleRight: true,
		whiteCanCastleLeft:  true,
		blackCanCastleRight: true,
		blackCanCastleLeft:  true,
		halfMove:            0,
		fullMove:            1,
	}}
}
