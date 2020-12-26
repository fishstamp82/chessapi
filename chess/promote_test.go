package chess

//func TestPromotion(t *testing.T) {
//	type piecePosition struct {
//		position Square
//		piece    Piece
//	}
//	table := []struct {
//		initialBoard     []piecePosition
//		Moves            [][2]Square
//		promotion        Piece
//		expectedStateOne State
//		expectedStateTwo State
//		checkMate        bool
//	}{
//		{
//			initialBoard: []piecePosition{
//				{e7, WhitePawn},
//				{g6, WhiteKing},
//				{g8, BlackKing},
//			},
//			Moves: [][2]Square{
//				{e7, e8},
//			},
//			promotion:        WhiteRook,
//			expectedStateOne: Promo,
//			expectedStateTwo: CheckMate,
//			checkMate:        true,
//		},
//	}
//
//	var state State
//	var err error
//
//	for _, row := range table {
//		b := NewEmptyBoard()
//		for _, pp := range row.initialBoard {
//			b.board[pp.position] = pp.piece
//		}
//		for _, move := range row.Moves {
//			state, err = b.move(move[0], move[1])
//			if err != nil {
//				t.Error(err)
//			}
//		}
//		if state != row.expectedStateOne {
//			t.Errorf("expected: %v, got: %v\n", row.expectedStateOne, state)
//		}
//
//		if state != row.expectedStateTwo {
//			t.Error("did not get state: Playing after promotion")
//		}
//	}
//}
