package main

import (
	"bufio"
	"chessapi/chess"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sort"
	"time"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	var reader *bufio.Reader
	_ = reader
	var move string
	var err error
	var moves [][2]string
	var context chess.Context

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

	b := chess.NewMailBoxBoard()
	for {
		if err == nil {
			fmt.Println(pretty(b.BoardMap()))
		}
		fmt.Printf("%s's turn\nmake a move...\n", b.Context.PlayersTurn)
		//reader = bufio.NewReader(os.Stdin)
		//move, _ = reader.ReadString('\n')
		//move = strings.TrimSuffix(move, "\n")
		//fmt.Printf("move : %s\n", move)
		move = pickRandom()
		context, err = b.Move(move)

		if err != nil {
			fmt.Println(err)
		} else {
			moves = append(moves, [2]string{move})
		}
		fmt.Println("state: " + context.State.String())
		if context.State == chess.CheckMate {
			winner := context.Winner
			fmt.Printf("game over, %s won", winner)
			fmt.Println(pretty(b.BoardMap()))
			break
		}
		fmt.Println(pretty(b.BoardMap()))
	}
	for _, val := range moves {
		fmt.Println(val)
	}
}

func pickRandom() string {
	rand.Seed(time.Now().UnixNano())
	lane := [8]string{"a", "b", "c", "d", "e", "f", "g", "h"}
	rank := [8]string{"1", "2", "3", "4", "5", "6", "7", "8"}
	moveFrom := lane[rand.Intn(8)] + rank[rand.Intn(8)]
	moveTo := lane[rand.Intn(8)] + rank[rand.Intn(8)]
	return moveFrom + moveTo
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
