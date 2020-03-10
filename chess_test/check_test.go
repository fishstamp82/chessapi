package chess

import (
	"chessapi/chess"
	"testing"
)

func TestCheck(t *testing.T) {
	table := []struct {
		moves    [][2]string
		expected string
	}{
		{
			moves: [][2]string{
				{"e2", "e4"},
				{"e7", "e5"},
				{"f2", "f4"},
				{"d8", "h4"},
			},
			expected: "Check",
		},
		{
			moves: [][2]string{
				{"e2", "e4"},
				{"e7", "e5"},
				{"f2", "f4"},
				{"d8", "g5"},
			},
			expected: "Playing",
		},
	}

	var ctx chess.Context
	var err error
	for _, row := range table {
		b := chess.NewMailBoxBoard()
		for _, val := range row.moves {
			s, toSquare := val[0], val[1]
			ctx, err = b.Move(s, toSquare)
			if err != nil {
				t.Errorf("error: %s\n", err)
			}

		}
		if ctx.State.String() != row.expected {
			t.Errorf("not in check ")
		}
	}
}

func TestCheckMate(t *testing.T) {
	table := []struct {
		moves    [][2]string
		expected string
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
			expected: "CheckMate",
			won:      "White",
		},
		{
			moves: [][2]string{
				{"e2", "e4"},
				{"e7", "e5"},
				{"f2", "f4"},
				{"d8", "g5"},
			},
			expected: "Playing",
			won:      "Noone",
		},
	}

	var err error
	for ind, row := range table {
		b := chess.NewMailBoxBoard()
		for _, val := range row.moves {
			fromSquare, toSquare := val[0], val[1]
			_, err = b.Move(fromSquare, toSquare)
			if err != nil {
				t.Error("error")
			}
		}
		if b.Context.State.String() != row.expected {
			t.Errorf("not check mate for case: %d\n", ind+1)
		}
		if won := b.Context.Winner; won.String() != row.won {
			t.Errorf("expected: %s, got: %s for case %d\n", row.won, won, ind+1)
		}
	}
}
