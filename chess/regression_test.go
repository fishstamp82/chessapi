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
				[2]string{"e2", "e4"},
				[2]string{"e7", "e5"},
				[2]string{"a2", "a3"},
				[2]string{"d8", "h4"},
				[2]string{"b2", "b3"},
				[2]string{"h4", "e4"},
			},
			expected: White,
		},
	}
	for _, row := range table {
		b := NewBoard()
		for _, val := range row.moves {
			s, t := val[0], val[1]
			b.Move(s, t)
		}
		if b.turn != row.expected {
			t.Errorf("expected: %v, got: %v\n", row.expected, b.turn)
		}
	}
}
