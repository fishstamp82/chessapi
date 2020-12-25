package chess

import "time"

type Player struct {
	Color     Color
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
