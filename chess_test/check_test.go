package chess

import (
	"chessapi/chess"
	"testing"
)

func TestCheck(t *testing.T) {
	table := []struct {
		moves    []string
		expected string
	}{
		{
			moves: []string{
				"e2e4",
				"e7e5",
				"f2f4",
				"d8h4",
			},
			expected: "Check",
		},
		{
			moves: []string{
				"e2e4",
				"e7e5",
				"f2f4",
				"d8g5",
			},
			expected: "Playing",
		},
	}

	var ctx chess.Context
	var err error
	for _, row := range table {
		b := chess.NewMailBoxBoard()
		for _, move := range row.moves {
			ctx, err = b.Move(move)
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
		moves    []string
		expected string
		won      string
	}{
		{
			moves: []string{
				"e2e4",
				"e7e5",
				"d1f3",
				"a7a6",
				"f1c4",
				"b7b6",
				"f3f7",
			},
			expected: "CheckMate",
			won:      "White",
		},
		{
			moves: []string{
				"e2e4",
				"e7e5",
				"f2f4",
				"d8g5",
			},
			expected: "Playing",
			won:      "Noone",
		},
	}

	var err error
	for ind, row := range table {
		b := chess.NewMailBoxBoard()
		for _, move := range row.moves {
			_, err = b.Move(move)
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
