package chess

import (
	"fmt"
	"regexp"
	"strings"
)

var allmovesRegexp = regexp.MustCompile(`(\d+)\.(.*\d|\+) (.*\d|\+)?\s+(1-0|1\/2-1\/2|0-1)?`)
var gameoverRegexp = regexp.MustCompile(`.*(1-0|1\/2-1\/2|0-1)$`)
var regularmoveRegexp = regexp.MustCompile(`(\d+)\.(.*\d|\+) (.*\d|\+)`)

//
//func pgnParse(reader io.Reader) ([]Move, error) {
//	var moves []Move
//	_ = moves
//	var pgnBytes []byte
//	var err error
//
//	pgnBytes, err = ioutil.ReadAll(reader)
//	if err != nil {
//		return nil, err
//	}
//	pgnString := filterMoves(string(pgnBytes))
//	movesStr := allmovesRegexp.FindAllString(pgnString, -1)
//	_ = movesStr
//	_ = pgnString
//	moves = getMoves(movesStr)
//	return moves, nil
//}
//
//func getMoves(allMoves []string) []Move {
//	b := NewEmptyBoard()
//
//	type move struct {
//		white string
//		black string
//	}
//	var realMoves []Move
//	var moves []move
//	var groups []string
//
//	for _, each := range allMoves {
//		if gameoverRegexp.MatchString(each) {
//			// Handle edge case reading final
//			continue
//		}
//		groups = regularmoveRegexp.FindStringSubmatch(each)
//		moves = append(moves, move{
//			white: groups[1],
//			black: groups[2],
//		})
//	}
//	for _, m := range moves {
//		realMoves = append(realMoves, parseNotation(m.white, b.board))
//	}
//	return []Move{}
//}
//
//func parseNotation(player Player, playerMove string, board [64]Piece, context Context) Move {
//	var targetSquare Square
//	targetSquareString := playerMove[len(playerMove)-2:]
//	targetSquare = stringToSquare[targetSquareString]
//	var fromInformation string
//	var piece Piece
//	//var lane, rank byte
//
//	// handle pawns first
//	//if isPromotion(playerMove){
//	//
//	//}
//
//	isPawnMove := isPawn(playerMove)
//	if isPawnMove {
//		switch player {
//		case White :
//			piece = WhitePawn
//		case Black:
//			piece = BlackPawn
//		}
//	} else {
//		bytePiece := playerMove[0]
//		piece = getPieceMust(byteToPiece[bytePiece], player)
//	}
//
//	if isCastle(playerMove){
//		switch player {
//		case White:
//			piece = WhiteKing
//		case Black:
//			piece = BlackKing
//		}
//		fromSquare, toSquare := decodeCastleMust(player, playerMove)
//		return createCastleMove(piece, fromSquare, toSquare, []MovementType{Castle})
//	}
//	if isCapture(playerMove){
//		fromInformation = strings.Split(playerMove, "x")[0]
//	} else {
//		fromInformation = playerMove[:len(playerMove)-2]
//	}
//
//	if !isPawnMove {
//		fromInformation = fromInformation[1:]
//	}
//
//	lane, rank := getLaneRank(fromInformation)
//	fromSquare := findFromSquares(piece, targetSquare, board, context)
//	disambiguated := disambiguate(fromSquare, lane, rank )
//	_, _ = targetSquare, disambiguated
//	return Move{}
//}
//
//func disambiguate(square []Square, lane byte, rank byte) Square {
//	if
//	return none
//}

//Used for disambiguation
func getLaneRank(fromInformation string) (byte, byte) {
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
