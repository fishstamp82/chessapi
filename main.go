package main

import (
	"bufio"
	"chessapi/chess"
	"fmt"
	"os"
	"os/signal"
	"sort"
	"strings"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	var reader *bufio.Reader
	var m1, m2 string
	var err error
	var moves [][2]string
	var state chess.State

	go func(moves *[][2]string) {
		for _ = range c {
			fmt.Printf("starting dump for reg test\n")
			fmt.Printf("moves: [][2]string{\n")
			for _, each := range *moves {
				fmt.Printf("\t{\"%s\", \"%s\"},\n", each[0], each[1])
			}
			fmt.Printf("},\n")

			os.Exit(0)
		}
	}(&moves)

	b := chess.NewBoard()
	for {

		if err == nil {
			fmt.Println(pretty(b.BoardMap()))
		}
		fmt.Printf("%s's turn\nmake a move...\n", b.PlayersTurn())
		reader = bufio.NewReader(os.Stdin)
		m1, _ = reader.ReadString('\n')
		m1 = strings.TrimSuffix(m1, "\n")
		fmt.Printf("move from: %s\nto:", m1)
		reader = bufio.NewReader(os.Stdin)
		m2, _ = reader.ReadString('\n')
		m2 = strings.TrimSuffix(m2, "\n")
		state, err = b.Move(m1, m2)

		if state == chess.Promotion {
			// read input from player
			continue
		}

		if err != nil {
			fmt.Println(err)
		} else {
			moves = append(moves, [2]string{m1, m2})
		}
		if b.CheckMate() {
			winner, _ := b.Won()
			fmt.Printf("game over, %s won", winner)
			break
		}
	}
	for _, val := range moves {
		fmt.Println(val)
	}
}

func pretty(m map[string]string) string {
	if len(m) != 64 {
		panic("not correct boardmap")
	}
	var s [][2]string
	var finalString string
	for key, val := range m {
		s = append(s, [2]string{key, val})
	}
	sort.Slice(s, func(i, j int) bool {
		return lookup(s[i][0], s[j][0])
	})
	var ind int
	for row := 7; row >= 0; row-- {
		finalString += "\n-----------------\n|"
		for col := 0; col <= 7; col++ {
			ind = row*8 + col
			finalString += s[ind][1] + "|"
		}
	}
	return finalString
}

func lookup(a, b string) bool {
	m := map[string]int{
		"a1": 0,
		"b1": 1,
		"c1": 2,
		"d1": 3,
		"e1": 4,
		"f1": 5,
		"g1": 6,
		"h1": 7,

		"a2": 8,
		"b2": 9,
		"c2": 10,
		"d2": 11,
		"e2": 12,
		"f2": 13,
		"g2": 14,
		"h2": 15,

		"a3": 16,
		"b3": 17,
		"c3": 18,
		"d3": 19,
		"e3": 20,
		"f3": 21,
		"g3": 22,
		"h3": 23,

		"a4": 24,
		"b4": 25,
		"c4": 26,
		"d4": 27,
		"e4": 28,
		"f4": 29,
		"g4": 30,
		"h4": 31,

		"a5": 32,
		"b5": 33,
		"c5": 34,
		"d5": 35,
		"e5": 36,
		"f5": 37,
		"g5": 38,
		"h5": 39,

		"a6": 40,
		"b6": 41,
		"c6": 42,
		"d6": 43,
		"e6": 44,
		"f6": 45,
		"g6": 46,
		"h6": 47,

		"a7": 48,
		"b7": 49,
		"c7": 50,
		"d7": 51,
		"e7": 52,
		"f7": 53,
		"g7": 54,
		"h7": 55,

		"a8": 56,
		"b8": 57,
		"c8": 58,
		"d8": 59,
		"e8": 60,
		"f8": 61,
		"g8": 62,
		"h8": 63,
	}
	return m[a] < m[b]
}
