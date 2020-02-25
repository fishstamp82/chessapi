package chess_test

import (
	"chessapi/chess"
	"testing"
)

func TestInitialization(t *testing.T) {
	b := chess.NewBoard()
	if b.PlayersTurn() != "white" {
		t.Error("not white's turn on start of the game")
	}

	table := []struct {
		square string
		piece  string
	}{
		{"a2", "white pawn"},
		{"b2", "white pawn"},
		{"c2", "white pawn"},
		{"d2", "white pawn"},
		{"e2", "white pawn"},
		{"f2", "white pawn"},
		{"g2", "white pawn"},
		{"h2", "white pawn"},
		{"a1", "white rook"},
		{"b1", "white knight"},
		{"c1", "white bishop"},
		{"d1", "white queen"},
		{"e1", "white king"},
		{"f1", "white bishop"},
		{"g1", "white knight"},
		{"h1", "white rook"},
		{"a7", "black pawn"},
		{"b7", "black pawn"},
		{"c7", "black pawn"},
		{"d7", "black pawn"},
		{"e7", "black pawn"},
		{"f7", "black pawn"},
		{"g7", "black pawn"},
		{"h7", "black pawn"},
		{"a8", "black rook"},
		{"b8", "black knight"},
		{"c8", "black bishop"},
		{"d8", "black queen"},
		{"e8", "black king"},
		{"f8", "black bishop"},
		{"g8", "black knight"},
		{"h8", "black rook"},
	}
	_ = table

	for _, testCase := range table {
		p := testCase.piece
		s := testCase.square
		if piece, _ := b.Get(s); piece != p {
			t.Errorf("expected: %s on %s, got: %s\n", p, s, piece)
		}
	}

}

func helper(b *chess.Board, t *testing.T) {

}
