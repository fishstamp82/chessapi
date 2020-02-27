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
		whitePawn Square
		expected  []Square
	}{
		{
			whiteRook: h1,
			blackRook: a2,
			whitePawn: d5,
			expected:  []Square{a1, b1, c1, d1, e1, f1, g1, h2, h3, h4, h5, h6, h7, h8},
		},
		{
			whiteRook: e1,
			blackRook: e5,
			whitePawn: d5,
			expected:  []Square{a1, b1, c1, d1, e2, e3, e4, e5, f1, g1, h1},
		},
		{
			whiteRook: e2,
			blackRook: e5,
			whitePawn: d2,
			expected:  []Square{e1, e3, e4, e5, f2, g2, h2},
		},
	}

	for _, row := range table {
		b := NewEmptyBoard()
		b.board[row.whiteRook] = WhiteRook
		b.board[row.whitePawn] = WhitePawn
		b.board[row.blackRook] = BlackRook
		got := b.rookMoves(row.whiteRook)

		if !sameAfterSort(got, row.expected) {
			t.Errorf("got: %v, expected: %v for %s on %s\n",
				pretty(got), pretty(row.expected), pieceToString[WhiteRook], squareToString[row.whiteRook])
		}
	}
}

func TestQueenMoves(t *testing.T) {
	table := []struct {
		whiteQueen Square
		blackQueen Square
		whitePawn  Square
		expected   []Square
	}{
		{
			whiteQueen: a8,
			blackQueen: a1,
			whitePawn:  a3,
			expected:   []Square{a4, a5, a6, a7, b8, c8, d8, e8, f8, g8, h8, b7, c6, d5, e4, f3, g2, h1},
		},
	}

	for _, row := range table {
		b := NewEmptyBoard()
		b.board[row.whiteQueen] = WhiteQueen
		b.board[row.whitePawn] = WhitePawn
		b.board[row.blackQueen] = BlackQueen
		got := b.queenMoves(row.whiteQueen)

		if !sameAfterSort(got, row.expected) {
			t.Errorf("got: %v, expected: %v for %s on %s\n",
				pretty(got), pretty(row.expected), pieceToString[WhiteQueen], squareToString[row.whiteQueen])
		}
	}
}

func TestKnightMoves(t *testing.T) {
	table := []struct {
		whiteKnight Square
		blackKnight Square
		expected    []Square
		fullBoard   bool
	}{
		{
			whiteKnight: a1,
			blackKnight: a2,
			expected:    []Square{b3, c2},
			fullBoard:   false,
		},
		{
			whiteKnight: e3,
			blackKnight: f5,
			expected:    []Square{d5, f5, g2, g4, d1, f1, c2, c4},
			fullBoard:   false,
		},
		{
			whiteKnight: g1,
			blackKnight: g8,
			expected:    []Square{f3, h3},
			fullBoard:   true,
		},
	}
	var b *Board
	for _, row := range table {
		if row.fullBoard {
			b = NewBoard()
		} else {
			b = NewEmptyBoard()
		}

		b.board[row.whiteKnight] = WhiteKnight
		b.board[row.blackKnight] = BlackKnight
		got := b.knightMoves(row.whiteKnight)

		if !sameAfterSort(got, row.expected) {
			t.Errorf("got: %v, expected: %v for %s on %s\n",
				pretty(got), pretty(row.expected), pieceToString[WhiteKnight], squareToString[row.whiteKnight])
		}
	}
}

func TestKingMoves(t *testing.T) {
	table := []struct {
		whiteKing Square
		blackKing Square
		expected  []Square
		fullBoard bool
	}{
		{
			whiteKing: a1,
			blackKing: a2,
			expected:  []Square{a2, b1, b2},
			fullBoard: false,
		},
		{
			whiteKing: e2,
			blackKing: d8,
			expected:  []Square{d1, d2, d3, e1, e3, f1, f2, f3},
			fullBoard: false,
		},
	}
	var b *Board
	for _, row := range table {
		if row.fullBoard {
			b = NewBoard()
		} else {
			b = NewEmptyBoard()
		}

		b.board[row.whiteKing] = WhiteKing
		b.board[row.blackKing] = BlackKing
		got := b.kingMoves(row.whiteKing)

		if !sameAfterSort(got, row.expected) {
			t.Errorf("got: %v, expected: %v for %s on %s\n",
				pretty(got), pretty(row.expected), pieceToString[WhiteKing], squareToString[row.whiteKing])
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