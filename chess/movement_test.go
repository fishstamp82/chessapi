package chess

import (
	"sort"
	"testing"
)

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

func TestBishopMoves(t *testing.T) {
	table := []struct {
		whiteBishop Square
		blackBishop Square
		expected    []Square
	}{
		{
			whiteBishop: a1,
			blackBishop: a2,
			expected:    []Square{b2, c3, d4, e5, f6, g7, h8},
		},
		{
			whiteBishop: a1,
			blackBishop: b2,
			expected:    []Square{b2},
		},
		{
			whiteBishop: e4,
			blackBishop: d5,
			expected:    []Square{d3, b1, c2, d5, f5, g6, h7, f3, g2, h1},
		},
		{
			whiteBishop: b1,
			blackBishop: b2,
			expected:    []Square{a2, c2, d3, e4, f5, g6, h7},
		},
		{
			whiteBishop: d4,
			blackBishop: h8,
			expected:    []Square{c3, b2, a1, c5, b6, a7, e3, f2, g1, e5, f6, g7, h8},
		},
	}

	for _, row := range table {
		b := NewEmptyBoard()
		b.board[row.whiteBishop] = WhiteBishop
		b.board[row.blackBishop] = BlackBishop
		got := b.bishopMoves(row.whiteBishop)

		if !sameAfterSort(got, row.expected) {
			t.Errorf("got: %v, expected: %v for %s on %s\n",
				pretty(got), pretty(row.expected), pieceToString[WhiteBishop], squareToString[row.whiteBishop])
		}
	}
}

func TestRookMoves(t *testing.T) {
	table := []struct {
		whiteRook Square
		blackRook Square
		expected  []Square
	}{
		{
			whiteRook: h1,
			blackRook: a2,
			expected:  []Square{a1, b1, c1, d1, e1, f1, g1},
		},
		{
			whiteRook: e1,
			blackRook: a2,
			expected:  []Square{a1, b1, c1, d1, f1, g1, h1},
		},
	}

	for _, row := range table {
		b := NewEmptyBoard()
		b.board[row.whiteRook] = WhiteRook
		b.board[row.blackRook] = BlackRook
		got := b.rookMoves(row.whiteRook)

		if !sameAfterSort(got, row.expected) {
			t.Errorf("got: %v, expected: %v for %s on %s\n",
				pretty(got), pretty(row.expected), pieceToString[WhiteRook], squareToString[row.whiteRook])
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
	var str []string
	for i := 0; i < len(s); i++ {
		str = append(str, squareToString[s[i]])
	}
	return str
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
