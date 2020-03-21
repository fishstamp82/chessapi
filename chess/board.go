package chess

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Context struct {
	State               State
	PlayersTurn         Player
	Winner              Player
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
	Context Context
}

func (b *Board) fenString() string {
	var cnt int
	var board string
	var sq Square
	for i := 7; i >= 0; i-- {
		cnt = 0
		for j := 0; j < 8; j++ {
			sq = Square(i*8 + j)

			switch p := b.board[sq]; {
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

	toMove := playerToFen[b.Context.PlayersTurn]

	var castle string
	if b.Context.whiteCanCastleRight {
		castle += pieceToFen[WhiteKing]
	}
	if b.Context.whiteCanCastleLeft {
		castle += pieceToFen[WhiteQueen]
	}
	if b.Context.blackCanCastleRight {
		castle += pieceToFen[BlackKing]
	}
	if b.Context.whiteCanCastleRight {
		castle += pieceToFen[BlackQueen]
	}

	if castle == "" {
		castle = "-"
	}

	var enpassant string
	if b.Context.enPassantSquare >= a1 {
		enpassant = b.Context.enPassantSquare.String()
	} else {
		enpassant = "-"
	}
	fullMove := strconv.Itoa(b.Context.fullMove)
	halfMove := strconv.Itoa(b.Context.halfMove)
	return fmt.Sprintf("%s %s %s %s %s %s", board, toMove, castle, enpassant, halfMove, fullMove)
}

var playerToFen = map[Player]string{
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

func (b *Board) String() string {
	return b.fenString()
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

	//Increment full move if this was blacks move
	if b.Context.PlayersTurn == Black {
		b.Context.fullMove += 1
	}

	//Increment half move if this was not a pawn move and not a capture
	isPawnMove := false
	for _, moveType := range m.moveTypes {
		if moveType == PawnMove {
			isPawnMove = true
		}
	}
	if isPawnMove {
		b.Context.halfMove = 0
	} else {
		b.Context.halfMove += 1
	}

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

func NewBoard() *Board {
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

func NewEmptyBoard() *Board {
	b := &Board{
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
	return b
}

func NewFromFEN(fen string) *Board {
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
	eb := NewEmptyBoard()
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
