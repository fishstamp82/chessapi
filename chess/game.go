package chess

const (
	White Player = iota + 1 // 1
	Black                   // 2
)

type Game struct {
	board     *Board
	state     string // playing or over
	inCheck   bool   // if in current state there is a check, for caching
	playerWon string // white or black if over
}
