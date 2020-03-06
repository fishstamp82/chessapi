package chess

import (
	"testing"
)

func TestWhitesTurn(t *testing.T) {
	table := []struct {
		moves    [][2]string
		expected Player
	}{
		{
			moves: [][2]string{
				{"e2", "e4"},
				{"e7", "e5"},
				//{"a2", "a3"},
				//{"d8", "h4"},
				//{"b2", "b3"},
				//{"h4", "e4"},
			},
			expected: White,
		},
	}
	for _, row := range table {
		b := NewMailBoxBoard()
		for _, val := range row.moves {
			s, t := val[0], val[1]
			_, _ = b.Move(s, t)
		}
		if b.context.playersTurn != row.expected {
			t.Errorf("expected: %v, got: %v\n", row.expected, b.context.playersTurn)
		}
	}
}

func TestBishopMovesAfterMoves(t *testing.T) {
	table := []struct {
		moves    [][2]string
		pieceAt  Square
		expected []Square
	}{
		{
			moves: [][2]string{
				{"e2", "e4"},
				{"d7", "d5"},
			},
			pieceAt:  f8,
			expected: []Square{},
		},
	}

	var bm []Square
	for _, row := range table {
		b := NewMailBoxBoard()
		for _, val := range row.moves {
			s, t := val[0], val[1]
			_, _ = b.Move(s, t)
		}
		bm = bishopMoves(row.pieceAt, b.board)
		if !sameAfterSquareSort(bm, row.expected) {
			t.Errorf("expected: %v, got: %v\n", row.expected, printPrettySquares(bm))
		}
	}
}
