package chess

import "testing"

func TestPieces(t *testing.T) {
	type piecePos struct {
		pos   Square
		piece Piece
	}

	table := []struct {
		pieces   []piecePos
		player   Player
		expected []Square
	}{
		{
			pieces: []piecePos{
				{a1, WhiteKnight},
				{b4, WhiteKing},
				{b7, BlackKing},
			},
			player:   White,
			expected: []Square{a1, b4},
		},
		{
			pieces: []piecePos{
				{a1, BlackKing},
				{b4, WhiteKing},
				{b7, BlackQueen},
			},
			player:   Black,
			expected: []Square{a1, b7},
		},
	}

	for ind, row := range table {
		b := NewEmptyBoard()
		for _, val := range row.pieces {
			b.board[val.pos] = val.piece
		}

		got := getPieces(row.player, b.board)
		if !sameAfterSquareSort(got, row.expected) {
			t.Errorf("got: %v, expected: %v for testcase: %d\n",
				printPrettySquares(got), printPrettySquares(row.expected), ind+1)
		}
	}
}

func TestKingSquare(t *testing.T) {
	type piecePos struct {
		pos   Square
		piece Piece
	}

	table := []struct {
		pieces   []piecePos
		player   Player
		expected Square
	}{
		{
			pieces: []piecePos{
				{a1, WhiteKnight},
				{b4, WhiteKing},
				{b7, BlackKing},
			},
			player:   White,
			expected: b4,
		},
		{
			pieces: []piecePos{
				{a1, BlackKing},
				{b4, WhiteKing},
				{b7, BlackQueen},
			},
			player:   Black,
			expected: a1,
		},
		{
			pieces: []piecePos{
				{h8, BlackKing},
				{b4, WhiteKing},
				{b7, BlackQueen},
			},
			player:   Black,
			expected: h8,
		},
	}

	for ind, row := range table {
		b := NewEmptyBoard()
		for _, val := range row.pieces {
			b.board[val.pos] = val.piece
		}

		got := getKingSquareMust(row.player, b.board)
		if got != row.expected {
			t.Errorf("got: %v, expected: %v for testcase: %d\n",
				got, row.expected, ind+1)
		}
	}
}
