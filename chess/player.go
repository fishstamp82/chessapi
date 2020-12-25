package chess

import (
	"fmt"
	"time"
)

type Player struct {
	Color     Color
	ID        string
	moves     []Move
	timeSpent time.Duration
}

func getOpponent(pl []Player, c Color) Player {
	var opp Player
	for _, p := range pl {
		if p.Color != c {
			opp = p
		}
	}
	return opp
}

func (p *Player) String() string {
	return fmt.Sprintf("%s\n", p.ID)
}
