package chess

import "time"

type Player struct {
	color     Color
	moves     []Move
	timeSpent time.Duration
}

func getOpponent(pl []Player, c Color) Player {
	var opp Player
	for _, p := range pl {
		if p.color != c {
			opp = p
		}
	}
	return opp
}
