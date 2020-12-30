![Go](https://github.com/fishstamp82/chessapi/workflows/Go/badge.svg?branch=master)

# Chess API
A library in GO to play chess.

## API

The library exposes a main struct, Game, that can be used to play a game of chess.

The entry-point is the Start() method, which runs the chess game
asynchronously, and provides a cleanup function to shut if off.

The interface to setting up a game is via the Handle* functions.

```go
package chessapi

// Before start
func (g *Game) HandlePick(uid string, cstr string) error
func (g *Game) HandleSetTime(t time.Duration) error
func (g *Game) HandleLeave(uid string) error

// During play
func (g *Game) HandleSetMove(move string) error
func (g *Game) HandleResign(uid string)
```

# Development

```sh
make test
```