package main

import (
	"bufio"
	"chessapi/chess"
	"fmt"
	"os"
	"strings"
)

func main() {
	var reader *bufio.Reader
	var m1, m2 string
	var err error
	b := chess.NewChessBoard()
	for {
		if err == nil {
			fmt.Println(b.StrRepr())
		}
		fmt.Printf("%s's turn\nmake a move...\n", b.PlayersTurn())
		reader = bufio.NewReader(os.Stdin)
		m1, _ = reader.ReadString('\n')
		m1 = strings.TrimSuffix(m1, "\n")
		fmt.Printf("move from: %s\nto:", m1)
		reader = bufio.NewReader(os.Stdin)
		m2, _ = reader.ReadString('\n')
		m2 = strings.TrimSuffix(m2, "\n")
		err = b.Move(m1, m2)
		if err != nil {
			fmt.Println(err)
		}
		if b.CheckMate() {
			winner, _ := b.Won()
			fmt.Printf("game over, %s won", winner)
			break
		}
	}
}
