package chess

import (
	"fmt"
	"sort"
	"testing"
)

func TestPawnMoves(t *testing.T) {
	var piece = WhitePawn
	b := NewBoard()
	be4 := NewBoard()
	be4.board[e4] = piece
	table := []struct {
		moves    []string
		pawnPos  Square
		expected []Move
	}{
		{
			moves: []string{
				"e2e4",
				"e7e5",
			},
			pawnPos: d2,
			expected: []Move{
				createMove(b.board, d2, d3, Regular),
				createMove(b.board, d2, d4, Regular),
			},
		},
		{
			moves: []string{
				"e2e4",
				"d7d5",
			},
			pawnPos: e4,
			expected: []Move{
				createMove(be4.board, e4, e5, Regular),
				createMove(be4.board, e4, d5, Capture),
			},
		},
		{
			moves: []string{
				"g2g4",
				"d7d5",
				"g4g5",
				"d5d4",
				"g5g6",
				"d4d3",
				"g6h7",
				"a7a6",
			},
			pawnPos:  h7,
			expected: createPawnPromotionMoves(White, h7, g8, BlackBishop, CapturePromotion),
		},
	}

	var err error
	for _, row := range table {
		b := NewBoard()
		for _, move := range row.moves {
			_, err = b.Move(move)
			if err != nil {
				t.Error(err)
			}
		}

		got := pawnMoves(row.pawnPos, b.board, b.Context.enPassantSquare)
		if !isMovesEqual(got, row.expected) {
			t.Errorf("got: %v, expected: %v for %s\n",
				printPrettyMoves(got), printPrettyMoves(row.expected), row.pawnPos)
		}
	}
}

func TestBishopMoves(t *testing.T) {
	var piece = WhiteBishop
	bf1 := NewBoard()
	bc4 := NewBoard()
	bf1.board[f1] = piece
	bc4.board[c4] = piece
	table := []struct {
		moves    []string
		pos      Square
		piece    Piece
		expected []Move
	}{
		{
			moves: []string{
				"e2e4",
			},
			pos:   f1,
			piece: piece,
			expected: []Move{
				createMove(bf1.board, f1, e2, Regular),
				createMove(bf1.board, f1, d3, Regular),
				createMove(bf1.board, f1, c4, Regular),
				createMove(bf1.board, f1, b5, Regular),
				createMove(bf1.board, f1, a6, Regular),
			},
		},
		{
			moves: []string{
				"e2e4",
				"d7d5",
				"f1c4",
				"d5d4",
			},
			pos:   c4,
			piece: piece,
			expected: []Move{
				createMove(bc4.board, c4, f1, Regular),
				createMove(bc4.board, c4, e2, Regular),
				createMove(bc4.board, c4, b3, Regular),
				createMove(bc4.board, c4, d3, Regular),
				createMove(bc4.board, c4, b5, Regular),
				createMove(bc4.board, c4, d5, Regular),
				createMove(bc4.board, c4, a6, Regular),
				createMove(bc4.board, c4, e6, Regular),
				createMove(bc4.board, c4, f7, Capture),
			},
		},
	}

	var err error
	for _, row := range table {
		b := NewBoard()
		for _, move := range row.moves {
			_, err = b.Move(move)
			if err != nil {
				t.Error(err)
			}
		}

		got := bishopMoves(row.pos, b.board)
		if !isMovesEqual(got, row.expected) {
			t.Errorf("got: %v, expected: %v for %s\n",
				printPrettyMoves(got), printPrettyMoves(row.expected), row.pos)
		}
	}
}

func TestKnightMoves(t *testing.T) {
	var piece = WhiteKnight
	var fun = knightMoves
	b := NewBoard()
	table := []struct {
		moves    []string
		pos      Square
		piece    Piece
		expected []Move
	}{
		{
			moves: []string{},
			pos:   g1,
			piece: piece,
			expected: []Move{
				createMove(b.board, g1, f3, Regular),
				createMove(b.board, g1, h3, Regular),
			},
		},
		{
			moves: []string{
				"e2e4",
			},
			pos:   g1,
			piece: piece,
			expected: []Move{
				createMove(b.board, g1, f3, Regular),
				createMove(b.board, g1, h3, Regular),
				createMove(b.board, g1, e2, Regular),
			},
		},
		{
			moves: []string{
				"f2f4",
				"e7e5",
				"f4f5",
				"e5e4",
				"f5f6",
			},
			pos:   g8,
			piece: BlackKnight,
			expected: []Move{
				createMove(b.board, g8, f6, Capture),
				createMove(b.board, g8, h6, Regular),
				createMove(b.board, g8, e7, Regular),
			},
		},
	}

	var err error
	for _, row := range table {
		b := NewBoard()
		for _, move := range row.moves {
			_, err = b.Move(move)
			if err != nil {
				t.Error(err)
			}
		}

		got := fun(row.pos, b.board)
		if !isMovesEqual(got, row.expected) {
			t.Errorf("got: %v, expected: %v for %s\n",
				printPrettyMoves(got), printPrettyMoves(row.expected), row.pos)
		}
	}
}

func TestRookMoves(t *testing.T) {
	var piece = WhiteRook
	var fun = rookMoves
	b := NewBoard()
	table := []struct {
		moves    []string
		pos      Square
		piece    Piece
		expected []Move
	}{
		{
			moves:    []string{},
			pos:      a1,
			piece:    piece,
			expected: []Move{},
		},
		{
			moves: []string{
				"a2a4",
			},
			pos:   a1,
			piece: piece,
			expected: []Move{
				createMove(b.board, a1, a2, Regular),
				createMove(b.board, a1, a3, Regular),
			},
		},
	}

	var err error
	for _, row := range table {
		b := NewBoard()
		for _, move := range row.moves {
			_, err = b.Move(move)
			if err != nil {
				t.Error(err)
			}
		}

		got := fun(row.pos, b.board)
		if !isMovesEqual(got, row.expected) {
			t.Errorf("got: %v, expected: %v for %s\n",
				printPrettyMoves(got), printPrettyMoves(row.expected), row.pos)
		}
	}
}

func TestQueenMoves(t *testing.T) {
	var piece = WhiteQueen
	var fun = queenMoves
	b := NewBoard()
	table := []struct {
		moves    []string
		pos      Square
		piece    Piece
		expected []Move
	}{
		{
			moves:    []string{},
			pos:      d1,
			piece:    piece,
			expected: []Move{},
		},
		{
			moves: []string{
				"e2e4",
				"h7h5",
				"d2d4",
			},
			pos:   d1,
			piece: piece,
			expected: []Move{
				createMove(b.board, d1, d2, Regular),
				createMove(b.board, d1, d3, Regular),
				createMove(b.board, d1, e2, Regular),
				createMove(b.board, d1, f3, Regular),
				createMove(b.board, d1, g4, Regular),
				createMove(b.board, d1, h5, Capture),
			},
		},
	}

	var err error
	for _, row := range table {
		b := NewBoard()
		for _, move := range row.moves {
			_, err = b.Move(move)
			if err != nil {
				t.Error(err)
			}
		}

		got := fun(row.pos, b.board)
		if !isMovesEqual(got, row.expected) {
			t.Errorf("got: %v, expected: %v for %s\n",
				printPrettyMoves(got), printPrettyMoves(row.expected), row.pos)
		}
	}
}

func TestKingMoves(t *testing.T) {
	var piece = WhiteKing
	var fun1 = kingMoves
	var fun2 = castleMoves
	b := NewBoard()
	table := []struct {
		moves    []string
		pos      Square
		piece    Piece
		expected []Move
	}{
		{
			moves:    []string{},
			pos:      e1,
			piece:    piece,
			expected: []Move{},
		},
		{
			moves: []string{
				"e2e4",
				"e7e5",
				"d2d4",
			},
			pos:   e1,
			piece: piece,
			expected: []Move{
				createMove(b.board, e1, e2, Regular),
				createMove(b.board, e1, d2, Regular),
			},
		},
		{
			moves: []string{
				"e2e4",
				"e7e5",
				"f1c4",
				"d7d5",
				"g1f3",
			},
			pos:   e1,
			piece: piece,
			expected: []Move{
				createMove(b.board, e1, e2, Regular),
				createMove(b.board, e1, f1, Regular),
				createMove(b.board, e1, g1, ShortCastle),
			},
		},
		{
			moves: []string{
				"d2d4",
				"d7d5",
				"c1f4",
				"e7e5",
				"b1c3",
				"a7a5",
				"d1d2",
				"a8a6",
				"a1b1",
			},
			pos:   e1,
			piece: piece,
			expected: []Move{
				createMove(b.board, e1, d1, Regular),
			},
		},
	}

	var err error
	for _, row := range table {
		b := NewBoard()
		for _, move := range row.moves {
			_, err = b.Move(move)
			if err != nil {
				t.Error(err)
			}
		}

		got := fun1(row.pos, b.board)
		got = append(got, fun2(row.pos, b.board, b.Context)...)
		if !isMovesEqual(got, row.expected) {
			t.Errorf("got: %v, expected: %v for %s\n",
				printPrettyMoves(got), printPrettyMoves(row.expected), row.pos)
		}
	}
}

func printPrettySquares(s []Square) []string {
	var str []string
	for i := 0; i < len(s); i++ {
		str = append(str, squareToString[s[i]])
	}
	return str
}

func printPrettyMoves(s []Move) []string {
	var str []string
	for i := 0; i < len(s); i++ {
		str = append(str, fmt.Sprintf("%s%s;%s", s[i].fromSquare, s[i].toSquare, s[i].moveType))
	}
	return str
}

func sameAfterSquareSort(a, b []Square) bool {
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

func isMovesEqual(a, b []Move) bool {
	sort.Slice(a, func(i, j int) bool { return a[i].toSquare < a[j].toSquare })
	sort.Slice(b, func(i, j int) bool { return b[i].toSquare < b[j].toSquare })
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if (a[i].piece != b[i].piece || a[i].toSquare != b[i].toSquare) || a[i].fromSquare != b[i].fromSquare || a[i].moveType != b[i].moveType {
			return false
		}
	}
	return true
}
