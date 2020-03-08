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
				{"e2", "e4"},
				{"e7", "e5"},
				{"f2", "f4"},
				{"d8", "h4"},
			},
			expected: true,
		},
		{
			moves: [][2]string{
				{"e2", "e4"},
				{"e7", "e5"},
				{"f2", "f4"},
				{"d8", "g5"},
			},
			expected: false,
		},
	}

	for _, row := range table {
		b := chess.NewBoard()
		for _, val := range row.moves {
			s, t := val[0], val[1]
			_, _ = b.Move(s, t)
		}
		if b.IsCheck() != row.expected {
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
				{"e2", "e4"},
				{"e7", "e5"},
				{"d1", "f3"},
				{"a7", "a6"},
				{"f1", "c4"},
				{"b7", "b6"},
				{"f3", "f7"},
			},
			expected: true,
			won:      "white",
		},
		{
			moves: [][2]string{
				{"e2", "e4"},
				{"e7", "e5"},
				{"f2", "f4"},
				{"d8", "g5"},
			},
			expected: false,
			won:      "",
		},
	}

	var err error
	for ind, row := range table {
		b := chess.NewBoard()
		for _, val := range row.moves {
			fromSquare, toSquare := val[0], val[1]
			_, err = b.Move(fromSquare, toSquare)
			if err != nil {
				t.Error("error")
			}
		}
		if b.CheckMate() != row.expected {
			t.Errorf("not check mate for case: %d\n", ind+1)
		}
		if won, _ := b.Won(); won != row.won {
			t.Errorf("expected: %s, got: %s for case %d\n", row.won, won, ind+1)
		}
	}
}
