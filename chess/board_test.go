package chess

import (
	"testing"
)

func TestInitialization(t *testing.T) {
	b := NewMailBoxBoard()
	if b.PlayersTurn() != "white" {
		t.Error("not white's turn on start of the game")
	}

	table := []struct {
		square Square
		piece  Piece
	}{
		{a2, WhitePawn},
		{b2, WhitePawn},
		{c2, WhitePawn},
		{d2, WhitePawn},
		{e2, WhitePawn},
		{f2, WhitePawn},
		{g2, WhitePawn},
		{h2, WhitePawn},
		{a1, WhiteRook},
		{b1, WhiteKnight},
		{c1, WhiteBishop},
		{d1, WhiteQueen},
		{e1, WhiteKing},
		{f1, WhiteBishop},
		{g1, WhiteKnight},
		{h1, WhiteRook},
		{a7, BlackPawn},
		{b7, BlackPawn},
		{c7, BlackPawn},
		{d7, BlackPawn},
		{e7, BlackPawn},
		{f7, BlackPawn},
		{g7, BlackPawn},
		{h7, BlackPawn},
		{a8, BlackRook},
		{b8, BlackKnight},
		{c8, BlackBishop},
		{d8, BlackQueen},
		{e8, BlackKing},
		{f8, BlackBishop},
		{g8, BlackKnight},
		{h8, BlackRook},
		{a5, Empty},
	}

	for _, row := range table {
		p := row.piece
		s := row.square
		if piece := b.board[s]; piece != p {
			t.Errorf("expected: %s on %s, got: %s\n", p, s, piece)
		}
	}

}

//
//func TestEnPassant(t *testing.T) {
//	table := []struct {
//		moves         [][2]string
//		expectedMoves []Square
//	}{
//		{
//			moves: [][2]string{
//				{"e2", "e4"},
//				{"d7", "d5"},
//				{"e4", "e5"},
//				{"f7", "f5"},
//			},
//			expectedMoves: []Square{e6, f6},
//		},
//	}
//
//	for _, row := range table {
//		b := NewMailBoxBoard()
//
//		for _, val := range row.moves {
//			s, t := val[0], val[1]
//			_, _ = b.Move(s, t)
//		}
//		moves := whitePawnMoves(e5 ,b.board)
//		if !sameAfterSquareSort(moves, row.expectedMoves){
//			t.Errorf("got: %s, expected: %s\n", moves, row.expectedMoves)
//		}
//	}
//
//}
