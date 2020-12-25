package chess

import (
	"testing"
)

func TestWhitesTurn(t *testing.T) {
	table := []struct {
		moves    []string
		expected Color
	}{
		{
			moves: []string{
				"e2e4",
				"e7e5",
			},
			expected: White,
		},
	}
	for _, row := range table {
		g := NewGame()
		g.Context.State = Playing
		g.Players = []Player{
			{
				Color: White,
			},
			{
				Color: Black,
			},
		}
		for _, move := range row.moves {
			err := g.Move(move)
			if err != nil {
				t.Error(err)
			}
		}
		if g.Context.ColorsTurn != row.expected {
			t.Errorf("expected: %v, got: %v\n", row.expected, g.Context.ColorsTurn)
		}
	}
}
