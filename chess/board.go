package chess

import (
	"errors"
	"fmt"
	"os"
	"runtime/debug"
	"strings"
)

type Context struct {
	State               State
	PlayersTurn         Player
	Winner              Player
	Score               string // 1-0, 0-1, 1/2-1/2
	moves               []Move
	whiteCanCastleRight bool
	whiteCanCastleLeft  bool
	blackCanCastleRight bool
	blackCanCastleLeft  bool
	enPassantSquare     Square
}

type Board struct {
	board   [64]Piece
	Context Context
}

//CLI repr of board
func (b *Board) BoardMap() map[string]string {
	board := map[string]string{}
	for square := a1; square <= h8; square++ {
		board[squareToString[square]] = pieceToUnicode[b.board[square]]
	}

	return board
}

// Move gets squares in human readable form, and performs a move
// error is nil on successful move
// arguments are two squares : "e2e4"
func (b *Board) Move(moveStr string) (Context, error) {
	if b.Context.State != Playing && b.Context.State != Check {
		return b.Context, errors.New("not in playing state")
	}
	fromSquare, toSquare, err := b.getSquare(moveStr)
	if err != nil {
		return b.Context, err
	}

	return b.move(fromSquare, toSquare)
}

func (b *Board) ValidMoves() ([]string, error) {
	if b.Context.State != Playing && b.Context.State != Check {
		return nil, errors.New("not in playing state")
	}
	moves := validMovesForPlayer(b.Context.PlayersTurn, b.board, b.Context)
	var strMoves []string
	for _, move := range moves {
		strMoves = append(strMoves, move.fromSquare.String()+move.toSquare.String())
	}
	return strMoves, nil
}

//func (b *Board) Promote(p Piece) (State, error) {
//	if !validPromotion(p, b.Context.PlayersTurn) {
//		return b.state, errors.New(fmt.Sprintf("%s not a valid piece \n", pieceToUnicode[p]))
//	}
//	b.board[b.Context.PawnPromotionSquare] = p
//
//if b.isCheckMated(b.getOpponent(b.Context.PlayersTurn)) {
//	b.state = CheckMate
//	return b.state, nil
//}
//
//	b.state = Playing
//	b.switchTurn()
//	return Playing, nil
//}

func (b *Board) move(fromSquare, toSquare Square) (Context, error) {

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

	defer func() {
		// runtime error and move sequence dump
		if r := recover(); r != nil {
			fmt.Println(r)
			fmt.Printf("%s\n", debug.Stack())
			fmt.Printf("Recovered in move, %s\n", m)
			for _, mov := range b.Context.moves {
				fmt.Printf("\"%s%s\"\n", mov.fromSquare, mov.toSquare)
			}
			os.Exit(0)
		}
	}()
	b.board = makeMove(m, b.board)

	opponentsKing := getKingSquare(opponent, b.board)
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

	b.abortCastling(m)
	b.Context.enPassantSquare = b.getEnPassantSquare(m)

	b.Context.moves = append(b.Context.moves, m)
	// Switch to other player
	b.switchTurn()
	return b.Context, nil
}

func isDraw(player Player, board [64]Piece, ctx Context) bool {
	moves := validMovesForPlayer(player, board, ctx)
	if moves == nil {
		return true
	}
	if twoKings(board) {
		return true
	}
	return false
}

func twoKings(b [64]Piece) bool {
	for i := a1; i <= h8; i++ {
		switch b[i] {
		case WhitePawn, WhiteBishop, WhiteKnight, WhiteRook, WhiteQueen, BlackPawn, BlackBishop, BlackKnight, BlackRook, BlackQueen:
			return false
		}
	}
	return true
}

func makeMove(m Move, b [64]Piece) [64]Piece {
	for _, pp := range m.piecePositions {
		b[pp.position] = pp.piece
	}
	return b
}

func getSquares(m []Move) []Square {
	var s []Square
	for _, m := range m {
		s = append(s, m.toSquare)
	}
	return s
}

func (b *Board) abortCastling(m Move) {

	switch m.fromSquare {
	case a1:
		b.Context.whiteCanCastleLeft = false
	case h1:
		b.Context.whiteCanCastleRight = false
	case a8:
		b.Context.blackCanCastleLeft = false
	case h8:
		b.Context.blackCanCastleRight = false
	case e1:
		b.Context.whiteCanCastleLeft = false
		b.Context.whiteCanCastleRight = false
	case e8:
		b.Context.blackCanCastleLeft = false
		b.Context.blackCanCastleRight = false
	}

	switch m.toSquare {
	case a1:
		b.Context.whiteCanCastleLeft = false
	case h1:
		b.Context.whiteCanCastleRight = false
	case a8:
		b.Context.blackCanCastleLeft = false
	case h8:
		b.Context.blackCanCastleRight = false
	}

}

func inCheck(kingSquare Square, board [64]Piece) bool {

	var opponent Player
	switch board[kingSquare] {
	case WhiteKing:
		opponent = Black
	case BlackKing:
		opponent = White
	}

	for _, oppPiece := range getPieces(opponent, board) {
		if inSquares(kingSquare, getSquares(getTargets(oppPiece, board))) {
			return true
		}
	}
	return false
}

//CheckMove if piece on square s check mates player p
func isCheckMated(kingSquare Square, board [64]Piece) bool {
	if !inCheck(kingSquare, board) {
		return false
	}

	var hero, opponent Player
	switch board[kingSquare] {
	case WhiteKing:
		hero = White
		opponent = Black
	case BlackKing:
		hero = Black
		opponent = White
	default:
		panic("called without a king piece")
	}

	var toBlock []Square
	var targets []Move

	//Try to escape with king
	var tmpKingSquare Square
	for _, move := range kingMoves(kingSquare, board) {
		board = makeMove(move, board)
		tmpKingSquare = getKingSquare(hero, board)
		if !inCheck(tmpKingSquare, board) {
			return false
		}
		board = makeMove(*move.reverseMove, board)
	}

	//get all attacks from opponent to block
	for _, square := range squaresWithoutKing(opponent, board) {
		targets = getTargets(square, board)
		if !inSquares(kingSquare, getSquares(targets)) {
			continue
		}
		for _, sq := range getBlocks(square, kingSquare, board) {
			toBlock = append(toBlock, sq)
		}
	}
	toBlock = uniqueSquares(toBlock)

	//Must Block all attacks from opponent in single move
	for _, source := range squaresWithoutKing(hero, board) {
		for _, target := range validMovesForSquare(source, board, Context{}) {
			if inSquares(target.toSquare, toBlock) {
				board = makeMove(target, board)
				if !inCheck(kingSquare, board) {
					return false
				}
				board = makeMove(*target.reverseMove, board)
			}
		}
	}

	return true
}

func (b *Board) getOpponent(p Player) Player {
	switch p {
	case White:
		return Black
	case Black:
		return White
	}

	panic("must be black or white")
}

func getPieces(p Player, board [64]Piece) []Square {
	var isWhite bool
	switch p {
	case White:
		isWhite = true
	case Black:
		isWhite = false
	}

	var pieces []Square
	for pos := a1; pos <= h8; pos += 1 {
		if board[pos] > 0 && isWhite {
			pieces = append(pieces, pos)
		} else if board[pos] < 0 && !isWhite {
			pieces = append(pieces, pos)

		}

	}
	return pieces
}

func (b *Board) getSquare(s string) (Square, Square, error) {
	if len(s) != 4 {
		return none, none, errors.New("wrong length")
	}
	sq1, found := stringToSquare[s[:2]]
	if !found {
		return none, none, errors.New(fmt.Sprintf("no such move: %s", s))
	}
	sq2, found := stringToSquare[s[2:]]
	if !found {
		return none, none, errors.New(fmt.Sprintf("no such move: %s", s))
	}
	return sq1, sq2, nil
}

func (b *Board) getEnPassantSquare(m Move) Square {
	if m.piece != WhitePawn && m.piece != BlackPawn {
		b.Context.enPassantSquare = none
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
func (b *Board) switchTurn() {
	if b.Context.PlayersTurn == White {
		b.Context.PlayersTurn = Black
	} else {
		b.Context.PlayersTurn = White
	}
}

func NewMailBoxBoard() *Board {
	b := &Board{
		Context: Context{
			State:               Playing,
			PlayersTurn:         White,
			Winner:              0,
			whiteCanCastleRight: true,
			whiteCanCastleLeft:  true,
			blackCanCastleRight: true,
			blackCanCastleLeft:  true,
		},
	}

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

func NewEmptyMailBoxBoard() *Board {
	b := &Board{
		Context: Context{
			State:       Playing,
			PlayersTurn: White,
		},
	}
	return b
}
