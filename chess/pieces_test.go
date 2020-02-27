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

		got := b.pieces(row.player)
		if !sameAfterSort(got, row.expected) {
			t.Errorf("got: %v, expected: %v for testcase: %d\n",
				pretty(got), pretty(row.expected), ind+1)
		}
	}
}
