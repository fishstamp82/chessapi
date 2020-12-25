package chess

import (
	"github.com/fishstamp82/chessapi/chess"
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

	for _, row := range table {
		game := chess.NewGame()
		game.Context.State = chess.Playing
		game.Players = []chess.Player{
			{
				Color: chess.White,
			},
			{
				Color: chess.Black,
			},
		}
		for _, move := range row.moves {
			err := game.Move(move)
			if err != nil {
				t.Errorf("error: %s\n", err)
			}

		}
		if game.Context.State.String() != row.expected {
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
		g := chess.NewGame()
		g.Context.State = chess.Playing
		g.Players = []chess.Player{
			{
				Color: chess.White,
			},
			{
				Color: chess.Black,
			},
		}
		for _, move := range row.moves {
			err = g.Move(move)
			if err != nil {
				t.Error(err)
			}
		}
		if g.Context.State.String() != row.expected {
			t.Errorf("not check mate for case: %d\n", ind+1)
		}
		if won := g.Context.Winner; won.String() != row.won {
			t.Errorf("expected: %s, got: %s for case %d\n", row.won, won, ind+1)
		}
	}
}
