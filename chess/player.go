package chess

import (
	"fmt"
	"time"
)

type Player struct {
	Color    Color
	ID       string
	TimeLeft time.Duration
	moves    []Move
}

func (p *Player) String() string {
	var id string
	if p == nil {
		return ""
	}
	switch p {
	case nil:
		id = ""
	default:
		id = p.ID
	}
	return fmt.Sprintf("%s\n", id)
}
