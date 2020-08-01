package main

import (
	"bufio"
	"fmt"
	"github.com/fishstamp82/chessapi/chess"
	"github.com/nsf/termbox-go"
	"os"
	"os/exec"
	"strings"
)

var (
	keyboardEventsChan = make(chan keyboardEvent)
)

type UserReader struct {
	Reader *bufio.Reader
}

func (ur *UserReader) interactWithUser() {
	// disable input buffering
	_ = exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	_ = exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	var b = make([]byte, 1)
	var err error
	for {
		_, err = os.Stdin.Read(b)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("bytes:", b)
		fmt.Println("string", string(b))
	}
}

func (ur *UserReader) readMove() string {
	ur.Reader = bufio.NewReader(os.Stdin)
	move, _ := ur.Reader.ReadString('\n')
	move = strings.TrimSuffix(move, "\n")

	return move
}

func NewUserReader() *UserReader {
	return &UserReader{}
}

func explorePgn(board *chess.Board) {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()
	go listenToKeyboard(keyboardEventsChan)

	mainloop: for {
		select {
		case e := <-keyboardEventsChan:
			switch e.eventType {
			case MOVE:
				d := keyToDirection(e.key)
				switch d {
				case BACK:
					//board.GoBack()
					fmt.Print(pretty(board.BoardMap()))
				case FORWARD:
					//board.GoForward()
					fmt.Print(pretty(board.BoardMap()))
				}
			case END:
				break mainloop
			}
		}
	}
}
