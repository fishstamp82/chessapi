package main

import (
	"bufio"
	"chessapi/chess"
	"fmt"
	"os"
	"strings"
)

type UserReader struct {
	Reader *bufio.Reader
}

func (ur *UserReader) readMove() string {
	ur.Reader = bufio.NewReader(os.Stdin)
	move, _ := ur.Reader.ReadString('\n')
	move = strings.TrimSuffix(move, "\n")
	fmt.Printf("move : %s\n", move)
	return ""
}

func NewUserReader() *UserReader {
	return &UserReader{}
}

func playPgn(board chess.Board) {
	var err error
	reader := NewUserReader()
	move := reader.readMove()
	for board.Context.State != chess.CheckMate && board.Context.State != chess.Draw {
		_, err = board.Move(move)
		if err != nil {
			fmt.Println(err)
		}
	}
}
