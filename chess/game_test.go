package chess

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGame_Start(t *testing.T) {
	tests := []struct {
		name   string
		game *Game
		want   chan<- bool
	}{
		{
			name: "white lose on time",
			game: NewGame(),
			want: make(chan<- bool),
		},
	}
	for _, tt := range tests {
		tt.game.Players = []Player{
			{Color: Black, ID: "black"},
			{Color: White, ID: "white"},
		}
		tt.game.startedAt = 0
		tt.game.startingTime = gameUpdateInterval / 2
		cleanup := tt.game.Start()
		defer cleanup()
		time.Sleep(gameUpdateInterval + 5 * time.Millisecond)
		assert.Equal(t, "0 - 1", tt.game.Context.Score(), "Score should be same")
	}
}
