package chess

//func TestDraw(t *testing.T) {
//	table := []struct {
//		moves    [][2]string
//		expected string
//	}{
//		{
//			moves: [][2]string{
//				{"e2", "e4"},
//				{"e7", "e5"},
//				{"f2", "f4"},
//				{"d8", "h4"},
//			},
//			expected: "Draw",
//		},
//	}
//
//	var ctx chess.Context
//	var err error
//	for _, row := range table {
//		b := chess.NewMailBoxBoard()
//		for _, val := range row.moves {
//			s, toSquare := val[0], val[1]
//			ctx, err = b.Move(s, toSquare)
//			if err != nil {
//				t.Errorf("error: %s\n", err)
//			}
//
//		}
//		if ctx.State.String() != row.expected {
//			t.Errorf("not in %s\n", row.expected)
//		}
//	}
//}
