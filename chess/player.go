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
	var id string
	if p == nil {
		id = ""
	}
	switch p {
	case nil:
		id = ""
	default:
		id = p.ID
	}
	return fmt.Sprintf("%s\n", id)
}
