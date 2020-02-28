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
				{"a2", "a3"},
				{"d8", "h4"},
				{"b2", "b3"},
				{"h4", "e4"},
			},
			expected: White,
		},
	}
	for _, row := range table {
		b := NewChessBoard()
		for _, val := range row.moves {
			s, t := val[0], val[1]
			_ = b.Move(s, t)
		}
		if b.turn != row.expected {
			t.Errorf("expected: %v, got: %v\n", row.expected, b.turn)
		}
	}
}
