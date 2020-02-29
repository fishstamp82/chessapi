![Go](https://github.com/fishstamp82/chessapi/workflows/Go/badge.svg?branch=master)

# Chess API
A library in go to play chess.

## API

The chessboard only has a method, which you use to
play a game of chess.

The library exposes an interface which should have
functionality to play chess.

```go

type Board interface {
	CheckMate() bool
	Draw() bool
	Won() (string, error)
	InCheck() bool
	PlayersTurn() string
	BoardMap() map[string]string
	Move(s, t string) error
}
```