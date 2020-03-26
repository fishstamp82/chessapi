package chess

import (
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"strings"
)

var fullMoveRegex = `\d+`
var notationRegex = `[a-hBNRQKOx\-]+[+O1-8]+`
var overRegex = `(1-0)|(0-1)|(1/2-1/2)`
var allmovesRegexp = regexp.MustCompile(fmt.Sprintf(`(%s)\.(%s)\s+(%s)?\s*(%s)?\s*`,
	fullMoveRegex,
	notationRegex,
	notationRegex,
	overRegex))
var gameOverRegexp = regexp.MustCompile(fmt.Sprintf(`(%s)\.(%s)\s+(%s)?\s*(%s)\s*`,
	fullMoveRegex,
	notationRegex,
	notationRegex,
	overRegex))
var moveRegexp = regexp.MustCompile(fmt.Sprintf(`(%s)\.(%s)\s+(%s)\s*`,
	fullMoveRegex,
	notationRegex,
	notationRegex))

func pgnParse(reader io.Reader) ([]Move, error) {
	var moves []Move
	_ = moves
	var pgnBytes []byte
	var err error

	pgnBytes, err = ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	pgnString := filterMoves(string(pgnBytes))
	movesStr := allmovesRegexp.FindAllString(pgnString, -1)
	moves = getMoves(movesStr)
	return moves, nil
}

func getMoves(allMoves []string) []Move {
	b := NewBoard()

	type move struct {
		player   Player
		notation string
	}
	var realMoves []Move
	var realMove Move
	var moves []move
	var groups []string

	for _, each := range allMoves {
		if gameOverRegexp.MatchString(each) {
			groups = gameOverRegexp.FindStringSubmatch(each)
			if len(groups) == 5 {
				moves = append(moves, move{
					player:   White,
					notation: groups[2],
				})
				moves = append(moves, move{
					player:   Black,
					notation: groups[3],
				})
			}
			continue
		}
		groups = moveRegexp.FindStringSubmatch(each)
		moves = append(moves, move{
			player:   White,
			notation: groups[2],
		})
		moves = append(moves, move{
			player:   Black,
			notation: groups[3],
		})
	}
	for _, m := range moves {
		realMove = parseNotation(m.player, m.notation, b.board, b.Context)
		realMoves = append(realMoves, realMove)
		b.board = makeMove(realMove, b.board)
	}
	return realMoves
}

func parseNotation(player Player, playerMove string, board [64]Piece, context Context) Move {
	var targetSquare Square

	var fromInformation string
	var promotion bool
	var piece, promoPiece Piece
	var movementTypes []MovementType
	var move Move

	if isCastle(playerMove) {
		switch player {
		case White:
			piece = WhiteKing
		case Black:
			piece = BlackKing
		}
		fromSquare, toSquare := decodeCastleMust(player, playerMove)
		movementTypes = append(movementTypes, Castle)
		return createCastleMove(piece, fromSquare, toSquare, movementTypes)
	}

	if isCheck(playerMove) {
		playerMove = playerMove[:len(playerMove)-1]
	}

	targetSquareString := playerMove[len(playerMove)-2:]
	targetSquare = stringToSquare[targetSquareString]

	if isPromotion(playerMove) {
		promotion = true
		bytePiece := playerMove[len(playerMove)-1]
		promoPiece = byteToPiece[bytePiece]
		playerMove = strings.Split(playerMove, "=")[0]
	}

	isPawnMove := isPawn(playerMove)
	if isPawnMove {
		switch player {
		case White:
			piece = WhitePawn
		case Black:
			piece = BlackPawn
		}
	} else {
		bytePiece := playerMove[0]
		piece = getPieceMust(byteToPiece[bytePiece], player)
	}

	if isCapture(playerMove) {
		movementTypes = append(movementTypes, Capture)
		fromInformation = strings.Split(playerMove, "x")[0]
	} else {
		fromInformation = playerMove[:len(playerMove)-2]
	}

	if !isPawnMove {
		movementTypes = append(movementTypes, Regular)
		fromInformation = fromInformation[1:]
	} else {
		movementTypes = append(movementTypes, PawnMove)
	}

	file, rank := getFileRank(fromInformation)
	fromSquares := findFromSquares(piece, targetSquare, board, context)
	fromSquare := disambiguateMust(fromSquares, file, rank)

	switch piece {
	case WhitePawn, BlackPawn:
		if promotion {
			move = createPawnPromotionMove(board, fromSquare, targetSquare, promoPiece, []MovementType{Promotion})
		} else {
			move = createPawnMove(piece, fromSquare, targetSquare, movementTypes)
		}
	default:
		move = createMove(board, fromSquare, targetSquare, movementTypes)
	}

	return move
}

func isCheck(move string) bool {
	if move[len(move)-1] == '+' {
		return true
	}
	return false
}

func disambiguateMust(squares []Square, file byte, rank byte) Square {
	if len(squares) == 1 {
		return squares[0]
	}

	//disambiguate by file first
	if file != 0 && rank == 0 {
		for _, square := range squares {
			l, _ := getFileRank(square.String())
			if file == l {
				return square
			}
		}
	}
	//disambiguate by rank second
	if file == 0 && rank != 0 {
		for _, square := range squares {
			_, r := getFileRank(square.String())
			if rank == r {
				return square
			}
		}
	}

	if file != 0 && rank != 0 {
		for _, square := range squares {
			l, r := getFileRank(square.String())
			if file == l && rank == r {
				return square
			}
		}
	}

	panic(fmt.Sprintf("no unique square found from squares: %s file: %q, rank: %q\n", squares, file, rank))
}

//Used for disambiguation
func getFileRank(fromInformation string) (byte, byte) {
	if len(fromInformation) == 2 {
		return fromInformation[0], fromInformation[1]
	}
	if len(fromInformation) == 1 {
		switch fromInformation[0] {
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h':
			return fromInformation[0], 0
		case '1', '2', '3', '4', '5', '6', '7', '8':
			return 0, fromInformation[0]
		}
	}
	return 0, 0
}

func decodeCastleMust(player Player, move string) (Square, Square) {
	if player == White && move == "O-O" {
		return e1, g1
	}
	if player == White && move == "O-O-O" {
		return e1, c1
	}
	if player == Black && move == "O-O" {
		return e8, g8
	}
	if player == Black && move == "O-O-O" {
		return e8, c8
	}
	panic(fmt.Sprintf("called decode Castle with player: %v, move: %v\n", player, move))
}

func getPieceMust(piece Piece, player Player) Piece {
	if piece == Bishop && player == White {
		return WhiteBishop
	}
	if piece == Knight && player == White {
		return WhiteKnight
	}
	if piece == Rook && player == White {
		return WhiteRook
	}
	if piece == Queen && player == White {
		return WhiteQueen
	}
	if piece == King && player == White {
		return WhiteKing
	}
	if piece == Bishop && player == Black {
		return BlackBishop
	}
	if piece == Knight && player == Black {
		return BlackKnight
	}
	if piece == Rook && player == Black {
		return BlackRook
	}
	if piece == Queen && player == Black {
		return BlackQueen
	}
	if piece == King && player == Black {
		return BlackKing
	}
	panic("no valid piece player combo")

}

func isPromotion(playerMove string) bool {
	for i := range playerMove {
		if playerMove[i] == '=' {
			return true
		}
	}
	return false
}

func isCastle(playerMove string) bool {
	switch playerMove[0] {
	case 'O':
		return true
	}
	return false
}

func isCapture(playerMove string) bool {
	for i := 0; i < len(playerMove); i++ {
		if playerMove[i] == 'x' {
			return true
		}
	}
	return false
}

func isPawn(playerMove string) bool {
	switch playerMove[0] {
	case 'B', 'N', 'R', 'Q', 'K', 'O':
		return false
	}
	return true
}

func findFromSquares(piece Piece, target Square, board [64]Piece, ctx Context) []Square {
	fromSquares := getPieceSquares(piece, board)
	var moves []Move
	var returnMoves []Square
	for _, fromSquare := range fromSquares {
		moves = validMovesForSquare(fromSquare, board, ctx)
		for _, move := range moves {
			if move.toSquare == target {
				returnMoves = append(returnMoves, move.fromSquare)
			}
		}
	}
	return returnMoves
}

func filterMoves(s string) string {
	var keep []string
	for _, line := range strings.Split(s, "\n") {
		if strings.HasPrefix(line, "[") {
			continue
		}
		line = strings.Trim(line, "\n")
		if len(line) == 0 {
			continue
		}
		keep = append(keep, line)
	}
	return strings.Join(keep, "\n")
}

var byteToPiece = map[byte]Piece{
	'B': Bishop,
	'N': Knight,
	'R': Rook,
	'Q': Queen,
	'K': King,
}
