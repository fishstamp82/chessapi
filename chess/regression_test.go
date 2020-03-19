package chess

import (
	"testing"
)

func TestWhitesTurn(t *testing.T) {
	table := []struct {
		moves    []string
		expected Player
	}{
		{
			moves: []string{
				"e2e4",
				"e7e5",
				//"a2a3",
				//"d8h4",
				//"b2b3",
				//"h4e4",
			},
			expected: White,
		},
	}
	var ctx Context
	for _, row := range table {
		b := NewBoard()
		for _, move := range row.moves {
			ctx, _ = b.Move(move)
		}
		if ctx.PlayersTurn != row.expected {
			t.Errorf("expected: %v, got: %v\n", row.expected, b.Context.PlayersTurn)
		}
	}
}
