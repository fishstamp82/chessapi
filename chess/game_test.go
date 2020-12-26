package chess

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGame_Start(t *testing.T) {
	tests := []struct {
		name      string
		game      *Game
		moves     []string
		wantScore string
	}{
		{
			name:      "white lose on time",
			game:      NewGame(),
			moves:     []string{},
			wantScore: "0 - 1",
		},
		{
			name:      "black lose on time after white makes a move",
			game:      NewGame(),
			moves:     []string{"e2e4"},
			wantScore: "1 - 0",
		},
	}
	for _, tt := range tests {
		tt.game.Players = []*Player{
			{Color: Black, ID: "black"},
			{Color: White, ID: "white"},
		}
		tt.game.startedAt = 0
		tt.game.StartingTime = gameUpdateInterval / 2
		cleanup := tt.game.Start()
		// Make all Moves:

		for _, m := range tt.moves {
			err := tt.game.Move(m)
			if err != nil {
				t.Fatal(err)
			}
		}
		defer cleanup()
		time.Sleep(gameUpdateInterval + 5*time.Millisecond)
		assert.Equal(t, tt.wantScore, tt.game.Context.Score(), "Score should be same")
	}
}
