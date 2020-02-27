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
				[2]string{"f2", "f4"},
				[2]string{"d8", "h4"},
			},
			expected: true,
		},
		{
			moves: [][2]string{
				[2]string{"e2", "e4"},
				[2]string{"e7", "e5"},
				[2]string{"f2", "f4"},
				[2]string{"d8", "g5"},
			},
			expected: false,
		},
	}

	for _, row := range table {
		b := chess.NewBoard()
		for _, val := range row.moves {
			s, t := val[0], val[1]
			b.Move(s, t)
		}
		if b.InCheck() != row.expected {
			t.Errorf("not in check ")
		}
	}
}

func TestCheckMate(t *testing.T) {
	table := []struct {
		moves    [][2]string
		expected bool
		won      string
	}{
		{
			moves: [][2]string{
				[2]string{"e2", "e4"},
				[2]string{"e7", "e5"},
				[2]string{"d1", "f3"},
				[2]string{"a7", "a6"},
				[2]string{"f1", "c4"},
				[2]string{"b7", "b6"},
				[2]string{"f3", "f7"},
			},
			expected: true,
			won:      "white",
		},
		{
			moves: [][2]string{
				[2]string{"e2", "e4"},
				[2]string{"e7", "e5"},
				[2]string{"f2", "f4"},
				[2]string{"d8", "g5"},
			},
			expected: false,
			won:      "",
		},
	}

	for ind, row := range table {
		b := chess.NewBoard()
		for _, val := range row.moves {
			s, t := val[0], val[1]
			b.Move(s, t)
		}
		if b.CheckMate() != row.expected {
			t.Errorf("not check mate for case: %d\n", ind+1)
		}
		if won, _ := b.Won(); won != row.won {
			t.Errorf("expected: %s, got: %s for case %d\n", row.won, won, ind+1)
		}
	}
}
