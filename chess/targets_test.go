package chess

import (
	"testing"
)

func TestPawnTargets(t *testing.T) {
	table := []struct {
		whitePawn  Square
		blackPawns []Square
		expected   []Square
	}{
		{
			whitePawn:  a2,
			blackPawns: []Square{b3},
			expected:   []Square{b3},
		},
		{
			whitePawn:  b2,
			blackPawns: []Square{a3, c3},
			expected:   []Square{a3, c3},
		},
		{
			whitePawn:  c6,
			blackPawns: []Square{a2, d7},
			expected:   []Square{d7},
		},
	}

	for _, row := range table {
		b := NewEmptyMailBoxBoard()
		b.board[row.whitePawn] = WhitePawn
		for _, val := range row.blackPawns {
			b.board[val] = BlackPawn
		}

		got := b.pawnTargets(row.whitePawn)
		if !sameAfterSort(got, row.expected) {
			t.Errorf("got: %v, expected: %v for %s\n",
				printPrettySquares(got), printPrettySquares(row.expected), squareToString[row.whitePawn])
		}
	}
}

func TestBishopTargets(t *testing.T) {
	table := []struct {
		whiteBishop  Square
		blackBishops []Square
		expected     []Square
	}{
		{
			whiteBishop:  a1,
			blackBishops: []Square{h8},
			expected:     []Square{h8},
		},
		{
			whiteBishop:  e4,
			blackBishops: []Square{d5, d3, f5, f3},
			expected:     []Square{d3, d5, f3, f5},
		},
		{
			whiteBishop:  c6,
			blackBishops: []Square{a2, d7},
			expected:     []Square{d7},
		},
	}

	for _, row := range table {
		b := NewEmptyMailBoxBoard()
		b.board[row.whiteBishop] = WhiteBishop
		for _, val := range row.blackBishops {
			b.board[val] = BlackBishop
		}

		got := b.bishopTargets(row.whiteBishop)
		if !sameAfterSort(got, row.expected) {
			t.Errorf("got: %v, expected: %v for %s\n",
				printPrettySquares(got), printPrettySquares(row.expected), squareToString[row.whiteBishop])
		}
	}
}

func TestKnightTargets(t *testing.T) {
	table := []struct {
		whiteKnight  Square
		blackKnights []Square
		expected     []Square
	}{
		{
			whiteKnight:  a1,
			blackKnights: []Square{b3, c2},
			expected:     []Square{b3, c2},
		},
		{
			whiteKnight:  e4,
			blackKnights: []Square{d6, d2},
			expected:     []Square{d6, d2},
		},
		{
			whiteKnight:  c6,
			blackKnights: []Square{d4, d8, e5, e7},
			expected:     []Square{d4, d8, e5, e7},
		},
	}

	for _, row := range table {
		b := NewEmptyMailBoxBoard()
		b.board[row.whiteKnight] = WhiteKnight
		for _, val := range row.blackKnights {
			b.board[val] = BlackKnight
		}

		got := b.knightTargets(row.whiteKnight)
		if !sameAfterSort(got, row.expected) {
			t.Errorf("got: %v, expected: %v for %s\n",
				printPrettySquares(got), printPrettySquares(row.expected), squareToString[row.whiteKnight])
		}
	}
}

func TestRookTargets(t *testing.T) {
	table := []struct {
		whiteRook  Square
		blackRooks []Square
		expected   []Square
	}{
		{
			whiteRook:  a1,
			blackRooks: []Square{f1, a7},
			expected:   []Square{f1, a7},
		},
		{
			whiteRook:  e4,
			blackRooks: []Square{h4},
			expected:   []Square{h4},
		},
		{
			whiteRook:  a1,
			blackRooks: []Square{},
			expected:   []Square{},
		},
	}

	for _, row := range table {
		b := NewEmptyMailBoxBoard()
		b.board[row.whiteRook] = WhiteRook
		for _, val := range row.blackRooks {
			b.board[val] = BlackRook
		}

		got := b.rookTargets(row.whiteRook)
		if !sameAfterSort(got, row.expected) {
			t.Errorf("got: %v, expected: %v for %s\n",
				printPrettySquares(got), printPrettySquares(row.expected), squareToString[row.whiteRook])
		}
	}
}

func TestQueenTargets(t *testing.T) {
	table := []struct {
		whiteQueen  Square
		whitePawns  []Square
		blackQueens []Square
		expected    []Square
	}{
		{
			whiteQueen:  a1,
			blackQueens: []Square{a8, h8, h1},
			whitePawns:  []Square{},
			expected:    []Square{a8, h8, h1},
		},
		{
			whiteQueen:  a1,
			whitePawns:  []Square{a2, b2, b1},
			blackQueens: []Square{a8, h8, h1},
			expected:    []Square{},
		},
	}

	for _, row := range table {
		b := NewEmptyMailBoxBoard()
		b.board[row.whiteQueen] = WhiteQueen
		for _, val := range row.blackQueens {
			b.board[val] = BlackQueen
		}
		for _, val := range row.whitePawns {
			b.board[val] = WhitePawn
		}

		got := b.queenTargets(row.whiteQueen)
		if !sameAfterSort(got, row.expected) {
			t.Errorf("got: %v, expected: %v for %s\n",
				printPrettySquares(got), printPrettySquares(row.expected), squareToString[row.whiteQueen])
		}
	}
}

func TestKingTargets(t *testing.T) {
	table := []struct {
		whiteKing   Square
		blackBishop []Square
		expected    []Square
	}{
		{
			whiteKing:   a1,
			blackBishop: []Square{a2, b2, b1},
			expected:    []Square{a2, b2, b1},
		},
	}

	for _, row := range table {
		b := NewEmptyMailBoxBoard()
		b.board[row.whiteKing] = WhiteKing
		for _, val := range row.blackBishop {
			b.board[val] = BlackKing
		}

		got := b.kingTargets(row.whiteKing)
		if !sameAfterSort(got, row.expected) {
			t.Errorf("got: %v, expected: %v for %s\n",
				printPrettySquares(got), printPrettySquares(row.expected), squareToString[row.whiteKing])
		}
	}
}
