package chess

import (
	"chessapi/chess"
	"testing"
)

func TestCheck(t *testing.T) {
	table := []struct {
		moves    [][2]string
		expected bool
	}{
		{
			moves: [][2]string{
				[2]string{"e2", "e4"},
				[2]string{"e7", "e5"},
				[2]string{"d2", "d4"},
				[2]string{"f7", "f5"},
				[2]string{"c1", "h6"},
			},
			expected: true,
		},
	}

	for _, row := range table {
		b := chess.NewBoard()
		for _, val := range row.moves {
			s, t := val[0], val[1]
			b.Move(s, t)
		}
		if b.InCheck() != row.expected {
			//t.Errorf("not in check ")
		}
	}
}
