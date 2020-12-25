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

func (p *Player) String() string {
	return fmt.Sprintf("%s\n", p.ID)
}
