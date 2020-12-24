package chess

import (
	"testing"
)

func TestHelper(t *testing.T) {
	table := []struct {
		square   Square
		squares  []Square
		expected bool
	}{
		{
			square:   2,
			squares:  []Square{4, 5},
			expected: false,
		},
		{
			square:   2,
			squares:  []Square{2, 4, 5},
			expected: true,
		},
		{
			square:   2,
			squares:  []Square{2},
			expected: true,
		},
	}
	for _, row := range table {
		got := inSquares(row.square, row.squares)
		if got != row.expected {
			t.Errorf("got: %t, expected: %t for %d in %v\n",
				got, row.expected, row.square, row.squares)
		}
	}
}

func TestValidPromotion(t *testing.T) {
	table := []struct {
		player   Color
		piece    Piece
		expected bool
	}{
		{
			player:   White,
			piece:    WhiteKing,
			expected: false,
		},
		{
			player:   Black,
			piece:    BlackKnight,
			expected: true,
		},
	}
	for _, row := range table {
		got := validPromotion(row.piece, row.player)
		if got != row.expected {
			t.Errorf("got: %t, expected: %t for %d in %v\n",
				got, row.expected, row.piece, row.player)
		}
	}
}
