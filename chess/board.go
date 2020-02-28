package chess

import (
	"errors"
	"fmt"
)

type Player byte
type Square int
type State byte

const (
	White Player = iota + 1 // 1
	Black                   // 2
)

const (
	Playing State = iota
	Over
	Draw
)

// row goes from 0 to 7
func (s Square) row() Square {
	return s / 8
}

// col goes from 0 to 7
func (s Square) col() Square {
	return s % 8
}

type Board interface {
	CheckMate() bool
	Draw() bool
	Won() (string, error)
	InCheck() bool
	PlayersTurn() string
	BoardMap() map[string]string
	Move(s, t string) error
}

type ChessBoard struct {
	board  [64]Piece
	turn   Player
	state  State
	winner Player
}

func (b *ChessBoard) CheckMate() bool {
	if b.state == Over {
		return true
	}
	return false
}

func (b *ChessBoard) Draw() bool {
	if b.state == Draw {
		return true
	}
	return false
}

func (b *ChessBoard) Won() (string, error) {
	if b.state != Over {
		return "", errors.New("game not over")
	}
	switch b.winner {
	case White:
		return "white", nil
	case Black:
		return "black", nil
	}
	return "", errors.New("no clear winner, bug")
}

func (b *ChessBoard) InCheck() bool {
	if b.inCheck(White) {
		return true
	}
	if b.inCheck(Black) {
		return true
	}
	return false
}

func (b *ChessBoard) PlayersTurn() string {
	var p Player
	p = b.turn
	if p == White {
		return "white"
	} else if p == Black {
		return "black"
	}
	panic("neither white nor blacks turn")
}

//CLI repr of board
func (b *ChessBoard) BoardMap() map[string]string {
	board := map[string]string{}
	for square := a1; square <= h8; square++ {
		board[squareToString[square]] = pieceToUnicode[b.board[square]]
	}

	return board
}

// Move gets squares in human readable form, and performs a move
// error is nil on successful move
// arguments are algebraic chess notation 'e2' -> 'e4'
func (b *ChessBoard) Move(s, t string) error {
	if b.state == Over {
		return errors.New("game over, can't move")
	}
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

	// Make Move
	p := b.board[sq1]
	b.board[sq1] = 0
	b.board[sq2] = p

	//If we are in check after a move, this move is not allowed
	if b.inCheck(b.turn) {
		b.board[sq1] = p
		b.board[sq2] = 0
		return errors.New(fmt.Sprintf("%s can't go to %s, check exposed\n", pieceToString[b.board[sq1]], t))
	}

	//check for check mate on opponent
	if b.isCheckMateBySquare(sq2, b.opponent(b.turn)) {
		b.state = Over
		b.winner = b.turn
		return nil
	}

	// Switch to other player
	b.switchTurn()
	return nil
}

// Given human readable string input "e2", return string
// repr of piece. if none, return "-"
func (b *ChessBoard) get(s string) (string, error) {
	sq, err := b.getSquare(s)
	if err != nil {
		return "", errors.New("bad move input, good format should be: 'e2', 'd3', etc")
	}
	p := b.board[sq]
	return pieceToString[p], nil
}

func (b *ChessBoard) inCheck(player Player) bool {
	ourKingPos := b.kingSquare(player)

	for _, oppPiece := range b.pieces(b.opponent(player)) {
		if inSquares(ourKingPos, b.targets(oppPiece)) {
			return true
		}
	}
	return false
}

// See if player is checked by piece on square s
func (b *ChessBoard) inCheckBySquare(s Square, player Player) bool {
	kingPos := b.kingSquare(player)
	if inSquares(kingPos, b.targets(s)) {
		return true
	}
	return false
}

//Check if piece on square s check mates player p
func (b *ChessBoard) isCheckMateBySquare(s Square, p Player) bool {
	if !b.inCheckBySquare(s, p) {
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
	kingMoves := b.moves(kingSquare)
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

	//Must Block attack, see if any piece can move to any of the blocking squares
	blocks := b.blocks(s, kingSquare)
	for _, piece := range b.piecesWithoutKing(p) {
		for _, move := range b.moves(piece) {
			if inSquares(move, blocks) {
				return false
			}
		}
	}

	return true
}

func (b *ChessBoard) opponent(p Player) Player {
	switch p {
	case White:
		return Black
	case Black:
		return White
	}

	panic("must be black or white")
}

func (b *ChessBoard) pieces(p Player) []Square {
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

func (b *ChessBoard) piecesWithoutKing(p Player) []Square {
	var isWhite bool
	var piece Piece
	switch p {
	case White:
		isWhite = true
	case Black:
		isWhite = false
	}

	var pieces []Square
	for pos := a1; pos <= h8; pos += 1 {
		piece = b.board[p]
		if piece == WhiteKing || piece == BlackKing {
			continue
		}
		if piece > 0 && isWhite {
			pieces = append(pieces, pos)
		} else if piece < 0 && !isWhite {
			pieces = append(pieces, pos)

		}

	}
	return pieces
}

func (b *ChessBoard) kingSquare(p Player) Square {
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

func NewChessBoard() *ChessBoard {
	b := &ChessBoard{state: Playing}

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

func NewEmptyChessBoard() *ChessBoard {
	b := &ChessBoard{state: Playing}
	return b
}

func NewBoard() Board {
	b := &ChessBoard{state: Playing}

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

func (b *ChessBoard) getSquare(s string) (Square, error) {
	if len(s) != 2 {
		return 0, errors.New("wrong length")
	}
	sq, found := stringToSquare[s]
	if !found {
		return 0, errors.New(fmt.Sprintf("no such square: %s", s))
	}
	return sq, nil
}

func (b *ChessBoard) switchTurn() {
	if b.turn == White {
		b.turn = Black
	} else {
		b.turn = White
	}
}

func (b *ChessBoard) moves(s Square) []Square {
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

func (b *ChessBoard) targets(s Square) []Square {
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

func (b *ChessBoard) blocks(s, t Square) []Square {
	p := b.board[s]
	var blocks []Square
	switch p {
	case WhitePawn:
		blocks = b.pawnBlocks(s)
	case BlackPawn:
		blocks = b.pawnBlocks(s)
	case WhiteBishop:
		blocks = b.bishopBlocks(s, t)
	case BlackBishop:
		blocks = b.bishopBlocks(s, t)
	case WhiteKnight:
		blocks = b.knightBlocks(s)
	case BlackKnight:
		blocks = b.knightBlocks(s)
	case WhiteRook:
		blocks = b.rookBlocks(s, t)
	case BlackRook:
		blocks = b.rookBlocks(s, t)
	case WhiteQueen:
		blocks = b.queenBlocks(s, t)
	case BlackQueen:
		blocks = b.queenBlocks(s, t)
	}
	return blocks
}
