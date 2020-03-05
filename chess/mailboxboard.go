package chess

import (
	"errors"
	"fmt"
)

type context struct {
	playersTurn         Player
	winner              Player
	pawnPromotionSquare Square
	whiteCanCastleRight bool
	whiteCanCastleLeft  bool
	blackCanCastleRight bool
	blackCanCastleLeft  bool
}

type MailBoxBoard struct {
	board   [64]Piece
	state   State
	context context
	winner  Player
}

func (b *MailBoxBoard) CheckMate() bool {
	if b.state == Over {
		return true
	}
	return false
}

func (b *MailBoxBoard) Draw() bool {
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

func (b *MailBoxBoard) InCheck() bool {
	if b.inCheck(White) {
		return true
	}
	if b.inCheck(Black) {
		return true
	}
	return false
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

func (b *MailBoxBoard) Promote(p Piece) (State, error) {
	if !validPromotion(p, b.context.playersTurn) {
		return b.state, errors.New(fmt.Sprintf("%s not a valid piece \n", pieceToUnicode[p]))
	}
	b.board[b.context.pawnPromotionSquare] = p

	if b.isCheckMated(b.getOpponent(b.context.playersTurn)) {
		b.state = Over
		return b.state, nil
	}

	b.state = Playing
	b.switchTurn()
	return Playing, nil
}

func (b *MailBoxBoard) move(fromSquare, toSquare Square) (State, error) {

	switch b.context.playersTurn {
	case White:
		if b.board[fromSquare] < 0 {
			return b.state, errors.New("white's turn")
		}
	case Black:
		if b.board[fromSquare] > 0 {
			return b.state, errors.New("black's turn")
		}
	}

	availMoves := validMoves(fromSquare, b.board, b.context)
	if !inSquares(toSquare, availMoves) {
		return b.state, errors.New(fmt.Sprintf("%s can't go to %s\n", pieceToString[b.board[fromSquare]], squareToString[toSquare]))
	}

	// Make Move
	piece := b.board[fromSquare]
	tmpPiece := b.board[toSquare]
	b.board[fromSquare] = 0
	b.board[toSquare] = piece

	//If we are in check after a move, this move is not allowed
	if b.inCheck(b.context.playersTurn) {
		b.board[fromSquare] = piece
		b.board[toSquare] = tmpPiece
		return b.state, errors.New(fmt.Sprintf("%s can't go to %s, check exposed\n", pieceToString[b.board[fromSquare]], squareToString[toSquare]))
	}

	//If we reach pawn promotion, return
	if pawnFinalRank(piece, toSquare) {
		b.state = Promotion
		b.context.pawnPromotionSquare = toSquare
		return b.state, nil
	}

	//is check mate for opponent
	if b.isCheckMated(b.getOpponent(b.context.playersTurn)) {
		b.state = Over
		b.context.winner = b.context.playersTurn
		return b.state, nil
	}

	//castles
	b.moveRookIfCastle(fromSquare, toSquare, piece)
	b.abortCastling(fromSquare, toSquare)

	// Switch to other player
	b.switchTurn()
	return b.state, nil
}

func (b *MailBoxBoard) moveRookIfCastle(fromSquare, toSquare Square, p Piece) {
	if p != WhiteKing && p != BlackKing {
		return
	}

	if fromSquare == e1 && toSquare == g1 {
		b.board[h1] = Empty
		b.board[f1] = WhiteRook
	}
	if fromSquare == e1 && toSquare == c1 {
		b.board[a1] = Empty
		b.board[d1] = WhiteRook
	}
	if fromSquare == e8 && toSquare == g8 {
		b.board[h8] = Empty
		b.board[f8] = BlackRook
	}
	if fromSquare == e8 && toSquare == c8 {
		b.board[a8] = Empty
		b.board[d8] = BlackRook
	}
}

func (b *MailBoxBoard) abortCastling(fromSquare, toSquare Square) {

	switch fromSquare {
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

	switch toSquare {
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
	return pieceToString[p], nil
}

func (b *MailBoxBoard) inCheck(player Player) bool {
	ourKingPos := b.kingSquare(player)

	for _, oppPiece := range b.getPieces(b.getOpponent(player)) {
		if inSquares(ourKingPos, targets(oppPiece, b.board)) {
			return true
		}
	}
	return false
}

//Check if piece on square s check mates player p
func (b *MailBoxBoard) isCheckMated(p Player) bool {
	if !b.inCheck(p) {
		return false
	}

	var king Piece
	switch p {
	case White:
		king = WhiteKing
	case Black:
		king = BlackKing
	}

	//Check possible escapes by the king
	kingSquare := b.kingSquare(p)
	kingMoves := validMoves(kingSquare, b.board, b.context)
	for _, move := range kingMoves {
		tmpPiece := b.board[move]
		b.board[kingSquare] = Empty
		b.board[move] = king
		if !b.inCheck(p) {
			b.board[kingSquare] = king
			b.board[move] = tmpPiece
			return false
		}
		b.board[kingSquare] = king
		b.board[move] = tmpPiece
	}

	//Must Block all attacks from getOpponent in one move
	var toBlock []Square
	var trgts []Square
	for _, square := range b.squaresWithoutKing(b.getOpponent(p)) {
		trgts = targets(square, b.board)
		if !inSquares(kingSquare, trgts) {
			continue
		}
		for _, sq := range blocks(square, kingSquare, b.board) {
			toBlock = append(toBlock, sq)
		}
	}
	toBlock = uniqueSquares(toBlock)

	var s, t Piece
	//Must Block all attacks from getOpponent in one move
	for _, source := range b.squaresWithoutKing(p) {
		for _, target := range validMoves(source, b.board, b.context) {
			if inSquares(target, toBlock) {
				s = b.board[source]
				t = b.board[target]
				b.board[target] = s
				if !b.inCheck(p) {
					b.board[source] = s
					b.board[target] = t
					return false
				}
				b.board[source] = s
				b.board[target] = t
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

func (b *MailBoxBoard) getPieces(p Player) []Square {
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
