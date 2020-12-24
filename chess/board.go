package chess

import (
	"errors"
	"fmt"
)

type Context struct {
	State               State
	PlayersTurn         Color
	Winner              Color
	Score               string // 1-0, 0-1, 1/2-1/2
	whiteCanCastleRight bool
	whiteCanCastleLeft  bool
	blackCanCastleRight bool
	blackCanCastleLeft  bool
	enPassantSquare     Square
	fullMove            int
	halfMove            int
}

func (c Context) String() string {
	return fmt.Sprintf("%s/%s/%s/%s/%v/%v/%v/%v/%s/%d/%d",
		c.State,
		c.PlayersTurn,
		c.Winner,
		c.Score,
		c.whiteCanCastleRight,
		c.whiteCanCastleLeft,
		c.blackCanCastleRight,
		c.blackCanCastleLeft,
		c.enPassantSquare,
		c.halfMove,
		c.fullMove,
	)
}

type Board struct {
	board   [64]Piece
}


var playerToFen = map[Color]string{
	White: "w",
	Black: "b",
}
var pieceToFen = map[Piece]string{
	WhitePawn:   "P",
	WhiteBishop: "B",
	WhiteKnight: "N",
	WhiteRook:   "R",
	WhiteQueen:  "Q",
	WhiteKing:   "K",
	BlackPawn:   "p",
	BlackBishop: "b",
	BlackKnight: "n",
	BlackRook:   "r",
	BlackQueen:  "q",
	BlackKing:   "k",
}

//CLI repr of board
func (b *Board) BoardMap() map[string]string {
	board := map[string]string{}
	for square := a1; square <= h8; square++ {
		board[squareToString[square]] = pieceToUnicode[b.board[square]]
	}

	return board
}




func ValidMoves(b *Board, p Color, c Context ) ([]string, error) {
	if c.State != Playing && c.State != Check {
		return nil, errors.New("not in playing state")
	}
	moves := validMovesForPlayer(p, b.board, c)
	var strMoves []string
	for _, move := range moves {
		strMoves = append(strMoves, move.fromSquare.String()+move.toSquare.String())
	}
	return strMoves, nil
}

func isDraw(player Color, board [64]Piece, ctx Context) bool {
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

func inCheck(kingSquare Square, board [64]Piece) bool {

	var opponent Color
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

	var hero, opponent Color
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
		tmpKingSquare = getKingSquareMust(hero, board)
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

func (b *Board) getOpponent(p Color) Color {
	switch p {
	case White:
		return Black
	case Black:
		return White
	}

	panic("must be black or white")
}

func getPieces(p Color, board [64]Piece) []Square {
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



var fenToPiece = map[byte]Piece{
	'P': WhitePawn,
	'B': WhiteBishop,
	'N': WhiteKnight,
	'R': WhiteRook,
	'Q': WhiteQueen,
	'K': WhiteKing,
	'p': BlackPawn,
	'b': BlackBishop,
	'n': BlackKnight,
	'r': BlackRook,
	'q': BlackQueen,
	'k': BlackKing,
	'1': Empty,
	'2': Empty,
	'3': Empty,
	'4': Empty,
	'5': Empty,
	'6': Empty,
	'7': Empty,
	'8': Empty,
}
