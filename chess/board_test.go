package chess

import (
	"sort"
	"testing"
)

//
func TestWhitePawnMove(t *testing.T) {
	table := []struct {
		whitePawn Square
		blackPawn Square
		expected  []Square
	}{
		{
			whitePawn: a2,
			blackPawn: a3,
			expected:  []Square{},
		},
		{
			whitePawn: b2,
			blackPawn: a3,
			expected:  []Square{a3, b3, b4},
		},
		{
			whitePawn: a3,
			blackPawn: h8,
			expected:  []Square{a4},
		},
	}

	for _, row := range table {
		b := NewEmptyBoard()
		b.board[row.whitePawn] = WhitePawn
		b.board[row.blackPawn] = BlackPawn

		got := b.whitePawnMoves(row.whitePawn)
		if !sameAfterSort(got, row.expected) {
			t.Errorf("got: %v, expected: %v for %s\n",
				pretty(got), pretty(row.expected), squareToString[row.whitePawn])
		}
	}
}

func TestBlackPawnMove(t *testing.T) {
	table := []struct {
		whitePawn Square
		blackPawn Square
		expected  []Square
	}{
		{
			whitePawn: a2,
			blackPawn: a3,
			expected:  []Square{},
		},
		{
			whitePawn: b2,
			blackPawn: a3,
			expected:  []Square{b2, a2},
		},
		{
			whitePawn: g6,
			blackPawn: h7,
			expected:  []Square{h5, h6, g6},
		},
	}

	for _, row := range table {
		b := NewEmptyBoard()
		b.board[row.whitePawn] = WhitePawn
		b.board[row.blackPawn] = BlackPawn

		got := b.blackPawnMoves(row.blackPawn)
		if !sameAfterSort(got, row.expected) {
			t.Errorf("got: %v, expected: %v for %s\n",
				pretty(got), pretty(row.expected), squareToString[row.blackPawn])
		}
	}
}

func pretty(s []Square) []string {
	new := []string{}
	for i := 0; i < len(s); i++ {
		new = append(new, squareToString[s[i]])
	}
	return new
}
func sameAfterSort(a, b []Square) bool {
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
