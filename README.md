![Go](https://github.com/fishstamp82/chessapi/workflows/Go/badge.svg?branch=master)

# Chess API
A library in go to play chess.

## API

The chessboard only has a method, which you use to
play a game of chess.

There are 3 functions of the Board interface
which are used.

```go

type Board interface {
    Move(string, string) error
    IsCheckMate() bool
    Won() (string. error) //who won, if game is over
}
```