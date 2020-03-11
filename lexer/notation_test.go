package lexer

//func TestReadNotation(t *testing.T) {
//	table := []struct {
//		player   chess.Player
//		notation string
//		expected chess.Move
//	}{
//		{
//			player:   chess.White,
//			notation: "e4",
//			expected: chess.Move{
//				piece:      chess.WhitePawn,
//				fromSquare: chess.e2,
//				toSquare:   chess.e4,
//				piecePositions: []chess.piecePosition{
//					{chess.WhitePawn, chess.e4},
//					{chess.Empty, chess.e2},
//				},
//				moveType: chess.Regular,
//				reverseMove: &chess.Move{
//					piece:      chess.WhitePawn,
//					fromSquare: chess.e4,
//					toSquare:   chess.e2,
//					piecePositions: []chess.piecePosition{
//						{chess.WhitePawn, chess.e2},
//						{chess.Empty, chess.e4},
//					},
//					moveType: chess.Regular,
//				},
//
//			},
//		},
//	}
//	for _, row := range table {
//		board := chess.NewMailBoxBoard()
//		move, err := readAlgebraicNotation(row.player, row.notation, board.board, board.Context)
//		if err != nil {
//			t.Error(err)
//		}
//		if ! moveEquals(move, row.expected){
//			t.Errorf("got: %v, expected: %v\n", move, row.expected)
//		}
//		if ! moveEquals(*move.reverseMove, *row.expected.reverseMove){
//			t.Errorf("got: %v, expected: %v\n", move, row.expected)
//		}
//	}
//}
//
//func moveEquals(m1, m2 chess.Move) bool {
//	if m1.piece != m2.piece {
//		return false
//	}
//	if m1.moveType != m2.moveType {
//		return false
//	}
//	if m1.fromSquare != m2.fromSquare {
//		return false
//	}
//	if m1.toSquare != m2.toSquare {
//		return false
//	}
//	if m1.toSquare != m2.toSquare {
//		return false
//	}
//	if len(m1.piecePositions) != len(m2.piecePositions){
//		return false
//	}
//	for i := range m1.piecePositions {
//		if m1.piecePositions[i] != m2.piecePositions[i] {
//			return false
//		}
//	}
//
//	return true
//}
