package chess

import (
	"errors"
	"fmt"
	"strings"
)

type context struct {
	playersTurn         Player
	winner              Player
	pawnPromotionSquare Square
	whiteCanCastleRight bool
	whiteCanCastleLeft  bool
	blackCanCastleRight bool
	blackCanCastleLeft  bool
	enPassantSquare     Square
}

type MailBoxBoard struct {
	board   [64]Piece
	state   State
	context context
	winner  Player
	score   string // 1 - 0, 0-1, \u00BD
}

func (b *MailBoxBoard) GetScore() string {
	if !(b.state == Over) {
		return "0-0"
	}
	switch b.context.winner {
	case White:
		return "1-0"
	case Black:
		return "0-1"
	case Noone:
		return "\u00BD-\u00BD"
	}
	panic("cant call this without a valid player")
}

func (b *MailBoxBoard) CheckMate() bool {
	if b.state == Over {
		return true
	}
	return false
}

func (b *MailBoxBoard) IsDraw() bool {
	if b.state == Draw {
		return true
	}
	return false
}

func (b *MailBoxBoard) Won() (string, error) {
	if b.state != Over {
		return "", errors.New("game not over")
	}
	switch b.context.winner {
	case White:
		return "white", nil
	case Black:
		return "black", nil
	}
	return "", errors.New("no clear winner, bug")
}

func (b *MailBoxBoard) PlayersTurn() string {
	p := b.context.playersTurn
	if p == White {
		return "white"
	} else if p == Black {
		return "black"
	}
	panic("neither white nor blacks turn")
}

//CLI repr of board
func (b *MailBoxBoard) BoardMap() map[string]string {
	board := map[string]string{}
	for square := a1; square <= h8; square++ {
		board[squareToString[square]] = pieceToUnicode[b.board[square]]
	}

	return board
}

// Move gets squares in human readable form, and performs a move
// error is nil on successful move
// arguments are algebraic chess notation 'e2' -> 'e4'
func (b *MailBoxBoard) Move(s, t string) (State, error) {
	if b.state != Playing {
		return b.state, errors.New("not in playing state")
	}
	fromSquare, err := b.getSquare(s)
	if err != nil {
		return b.state, errors.New(fmt.Sprintf("bad move input, got: %v, good format: 'e2'", s))
	}
	toSquare, err := b.getSquare(t)
	if err != nil {
		return b.state, errors.New(fmt.Sprintf("bad second input, got: %v, good format: 'e4'", t))
	}

	return b.move(fromSquare, toSquare)

}

//func (b *MailBoxBoard) Promote(p Piece) (State, error) {
//	if !validPromotion(p, b.context.playersTurn) {
//		return b.state, errors.New(fmt.Sprintf("%s not a valid piece \n", pieceToUnicode[p]))
//	}
//	b.board[b.context.pawnPromotionSquare] = p
//
//if b.isCheckMated(b.getOpponent(b.context.playersTurn)) {
//	b.state = Over
//	return b.state, nil
//}
//
//	b.state = Playing
//	b.switchTurn()
//	return Playing, nil
//}

func (b *MailBoxBoard) move(fromSquare, toSquare Square) (State, error) {

	var opponent Player
	switch b.context.playersTurn {
	case White:
		if b.board[fromSquare] < 0 {
			return b.state, errors.New("white's turn")
		}
		opponent = Black
	case Black:
		if b.board[fromSquare] > 0 {
			return b.state, errors.New("black's turn")
		}
		opponent = White
	}

	availMoves := validMoves(fromSquare, b.board, b.context)

	availSquares := getSquares(availMoves)
	if !inSquares(toSquare, availSquares) {
		return b.state, errors.New(fmt.Sprintf("%s can't go to %s\n", b.board[fromSquare], squareToString[toSquare]))
	}

	//
	m := Move{}
	for _, move := range availMoves {
		if move.fromSquare == fromSquare && move.toSquare == toSquare {
			m = move
		}
	}
	if m.toSquare == none {
		return b.state, &NoMoveError{Move: strings.Join([]string{fromSquare.String(), toSquare.String()}, "")}
	}

	// Make Move
	b.board = makeMove(m, b.board)

	//is check mate for opponent
	opponentsKing := getKingSquare(opponent, b.board)
	if isCheckMated(opponentsKing, b.board) {
		b.state = Over
		b.context.winner = b.context.playersTurn
		return b.state, nil
	}

	//castles
	b.abortCastling(m)
	b.context.enPassantSquare = b.getEnPassantSquare(m)

	// Switch to other player
	b.switchTurn()
	return b.state, nil
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

func (b *MailBoxBoard) abortCastling(m Move) {

	switch m.fromSquare {
	case a1:
		b.context.whiteCanCastleLeft = false
	case h1:
		b.context.whiteCanCastleRight = false
	case a8:
		b.context.blackCanCastleLeft = false
	case h8:
		b.context.blackCanCastleRight = false
	case e1:
		b.context.whiteCanCastleLeft = false
		b.context.whiteCanCastleRight = false
	case e8:
		b.context.blackCanCastleLeft = false
		b.context.blackCanCastleRight = false
	}

	switch m.toSquare {
	case a1:
		b.context.whiteCanCastleLeft = false
	case h1:
		b.context.whiteCanCastleRight = false
	case a8:
		b.context.blackCanCastleLeft = false
	case h8:
		b.context.blackCanCastleRight = false
	}

}

// Given human readable string input "e2", return string
// repr of piece. if none, return "-"
func (b *MailBoxBoard) stringRepr(s string) (string, error) {
	sq, err := b.getSquare(s)
	if err != nil {
		return "", errors.New("bad move input, good format should be: 'e2', 'd3', etc")
	}
	p := b.board[sq]
	return fmt.Sprintf("%s", p), nil
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

//Check if piece on square s check mates player p
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
		for _, target := range validMoves(source, board, context{}) {
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

func (b *MailBoxBoard) getOpponent(p Player) Player {
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

func (b *MailBoxBoard) getSquare(s string) (Square, error) {
	if len(s) != 2 {
		return 0, errors.New("wrong length")
	}
	sq, found := stringToSquare[s]
	if !found {
		return 0, errors.New(fmt.Sprintf("no such square: %s", s))
	}
	return sq, nil
}

func (b *MailBoxBoard) getEnPassantSquare(m Move) Square {
	if m.piece != WhitePawn && m.piece != BlackPawn {
		b.context.enPassantSquare = none
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
func (b *MailBoxBoard) switchTurn() {
	if b.context.playersTurn == White {
		b.context.playersTurn = Black
	} else {
		b.context.playersTurn = White
	}
}

func NewMailBoxBoard() *MailBoxBoard {
	b := &MailBoxBoard{
		state: Playing,
		context: context{
			playersTurn:         White,
			winner:              0,
			pawnPromotionSquare: 0,
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

func NewEmptyMailBoxBoard() *MailBoxBoard {
	b := &MailBoxBoard{
		state: Playing,
		context: context{
			playersTurn: White,
		},
	}
	return b
}
