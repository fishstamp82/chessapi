package chess

import (
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
//func parseNotation(player Player, playerMove string, board [64]Piece) Move {
//	targetSquare := move[len(move)-2:]
//	isCapture := isCapture(playerMove)
//	fromSquare := findFromSquares(WhitePawn)
//	_, _ = targetSquare, fromSquare
//	return Move{}
//}

func isCapture(playerMove string) bool {
	for i := 0; i < len(playerMove); i++ {
		if playerMove[i] == 'x' {
			return true
		}
	}
	return false
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
