package chess

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

const (
	gameUpdateInterval = 100 * time.Millisecond
)

var (
	ErrNotPlaying = errors.New("not in a playing GameState")

	ErrAlreadyPlaying = errors.New("player already seated")
	ErrColorTaken     = errors.New("such Color already taken")
)

type Game struct {
	Board        *Board  `json:"board"`
	Context      Context `json:"context"`
	Players      []Player `json:"players"`
	moves        []Move
	StartingTime time.Duration
	startedAt    int64
}

func (g *Game) Start() func() {
	for _, p := range g.Players {
		p.TimeLeft = g.StartingTime
	}
	g.Context.State = Playing
	g.startedAt = timeNow()
	exit := make(chan bool)
	ticker := time.NewTicker(gameUpdateInterval)
	cleanup := func() {
		exit <- true
		ticker.Stop()
	}
	runAsync := func() {
		for {
			select {
			case <-exit:
				g.End()
			case <-ticker.C:
				p := g.getPlayer(g.Context.ColorsTurn)
				p.TimeLeft -= gameUpdateInterval
				if p.TimeLeft < 0 {
					opp := g.getOpponent(p)
					g.Context.WinningPlayer = &opp
					g.Context.State = Over
				}
			}
		}
	}
	go runAsync()

	return cleanup
}

func (g *Game) End() {
	g.Context.State = Idle
}

// Move gets squares in human readable form, and performs a move
// error is nil on successful move
// arguments are two squares : "e2e4"
func (g *Game) Move(moveStr string) error {
	if g.Context.State != Playing && g.Context.State != Check {
		return fmt.Errorf("not in playing state")
	}
	fromSquare, toSquare, err := g.Board.getSquare(moveStr)
	if err != nil {
		return err
	}

	return g.move(fromSquare, toSquare)
}

// Move gets squares in human readable form, and performs a move
// error is nil on successful move
// arguments are two squares : "e2e4"
func (g *Game) MoveNotation(move Move) error {
	if g.Context.State != Playing && g.Context.State != Check {
		return fmt.Errorf("not in playing state")
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
				toSkip, err = strconv.Atoi(ranks[i][j : j+1])
				if err != nil {
					panic(err)
				}
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

func (g *Game) HandleSetMove(move string) error {
	if g.Context.State != Playing && g.Context.State != Check {
		return ErrNotPlaying
	}
	err := g.Move(move)
	return err
}

func (g *Game) HandleSetTime(t time.Duration) error {
	if !(g.Context.State == Idle){
		return ErrAlreadyPlaying
	}
	g.StartingTime = t
	return nil
}

func (g *Game) HandleResign(uid string) {
	for _, ps := range g.Players {
		if (ps.ID == uid) && (ps.Color == White) {
			won := g.getPlayer(Black)
			g.Context.WinningPlayer = &won
		}
		if (ps.ID == uid) && (ps.Color == Black) {
			won := g.getPlayer(White)
			g.Context.WinningPlayer = &won
		}
	}
}

// Enable a player to leave a game before it starts
func (g *Game) HandleLeave(uid string) error {
	if g.Context.State != Idle {
		return fmt.Errorf("can't leave in-progress game")
	}
	var toDelete int
	for i, ps := range g.Players {
		if ps.ID == uid {
			toDelete = i
		}
	}
	g.Players[toDelete] = g.Players[len(g.Players)-1]
	g.Players[len(g.Players)-1] = Player{}
	g.Players = g.Players[:len(g.Players)-1]
	return nil
}

func (g *Game) HandlePick(uid string, cstr string) error {
	c := getColor(cstr)
	for _, ps := range g.Players {
		if ps.Color == c {
			return ErrColorTaken
		}
		if ps.ID == uid {
			return ErrAlreadyPlaying
		}
	}
	g.Players = append(g.Players, Player{ID: uid, Color: c})
	return nil
}

func getColor(cstr string) Color {
	switch strings.ToLower(cstr) {
	case "white":
		return White
	case "black":
		return Black
	}
	panic(fmt.Errorf("no such color: %s\n", cstr))
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

func (g *Game) FenString() string {
	var cnt int
	var board string
	var sq Square
	for i := 7; i >= 0; i-- {
		cnt = 0
		for j := 0; j < 8; j++ {
			sq = Square(i*8 + j)

			switch p := g.Board.board[sq]; {
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

	toMove := playerToFen[g.Context.ColorsTurn]

	var castle string
	if g.Context.whiteCanCastleRight {
		castle += pieceToFen[WhiteKing]
	}
	if g.Context.whiteCanCastleLeft {
		castle += pieceToFen[WhiteQueen]
	}
	if g.Context.blackCanCastleRight {
		castle += pieceToFen[BlackKing]
	}
	if g.Context.whiteCanCastleRight {
		castle += pieceToFen[BlackQueen]
	}

	if castle == "" {
		castle = "-"
	}

	var enpassant string
	if g.Context.enPassantSquare >= a1 {
		enpassant = g.Context.enPassantSquare.String()
	} else {
		enpassant = "-"
	}
	halfMove := strconv.Itoa(g.Context.halfMove)
	fullMove := strconv.Itoa(g.Context.fullMove)
	return fmt.Sprintf("%s %s %s %s %s %s", board, toMove, castle, enpassant, halfMove, fullMove)
}

func (g *Game) move(fromSquare, toSquare Square) error {

	var opponent Color
	switch g.Context.ColorsTurn {
	case White:
		if g.Board.board[fromSquare] < 0 {
			return fmt.Errorf("white's turn\n")
		}
		opponent = Black
	case Black:
		if g.Board.board[fromSquare] > 0 {
			return fmt.Errorf("black's turn\n")
		}
		opponent = White
	}

	availMoves := validMovesForSquare(fromSquare, g.Board.board, g.Context)

	availSquares := getSquares(availMoves)
	if !inSquares(toSquare, availSquares) {
		return fmt.Errorf("%s can't move to %s\n", g.Board.board[fromSquare], squareToString[toSquare])
	}

	//todo: replace with function thate uses chess algebraic notation
	var m Move
	for _, move := range availMoves {
		if move.fromSquare == fromSquare && move.toSquare == toSquare {
			m = move
		}
	}
	if m.toSquare == none {
		return fmt.Errorf("target square %s is 'none'\n", squareToString[toSquare])
	}

	// Commit the move to the board, update timers
	g.Board.board = makeMove(m, g.Board.board)
	p := g.getPlayer(g.Context.ColorsTurn)
	p.moves = append(p.moves, m)

	opponentsKing := getKingSquareMust(opponent, g.Board.board)
	if inCheck(opponentsKing, g.Board.board) {
		g.Context.State = Check
	} else {
		g.Context.State = Playing
	}

	if isCheckMated(opponentsKing, g.Board.board) {
		g.Context.State = CheckMate
		g.Context.WinningPlayer = &p
		return nil
	}

	if isDraw(opponent, g.Board.board, g.Context) {
		g.Context.State = Draw
		return nil
	}

	// Invalidate castling rules if move prevents castling
	g.abortCastling(m)

	// Set possible enPassantSquare as available move for next move
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

	// Switch next turn to other player
	g.switchTurn()
	return nil
}

func (g *Game) getEnPassantSquare(m Move) Square {
	if m.piece != WhitePawn && m.piece != BlackPawn {
		g.Context.enPassantSquare = none
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

func (g *Game) switchTurn() {
	if g.Context.ColorsTurn == White {
		g.Context.ColorsTurn = Black
	} else {
		g.Context.ColorsTurn = White
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

func (g *Game) getPlayer(c Color) Player {
	for _, p := range g.Players {
		if p.Color == c {
			return p
		}
	}
	panic(fmt.Sprintf("no player with Color %s in game", c.String()))
}

func (g *Game) getOpponent(p Player) Player {
	for _, ps := range g.Players {
		if ps.ID != p.ID {
			return ps
		}
	}
	panic(fmt.Sprintf("can't find opponent to %v in game", p))
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

	return &Game{Board: b,
		Context: Context{
			State:               Idle,
			ColorsTurn:          White,
			WinningPlayer:       nil,
			whiteCanCastleRight: true,
			whiteCanCastleLeft:  true,
			blackCanCastleRight: true,
			blackCanCastleLeft:  true,
			halfMove:            0,
			fullMove:            1,
		},
		Players: nil,
	}
}
