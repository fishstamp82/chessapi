package chess

import (
	"fmt"
	"sort"
	"testing"
)

func TestPawnMoves(t *testing.T) {
	var piece = WhitePawn
	b := NewMailBoxBoard()
	be4 := NewMailBoxBoard()
	be4.board[e4] = piece
	table := []struct {
		moves    [][2]string
		pawnPos  Square
		expected []Move
	}{
		{
			moves: [][2]string{
				{"e2", "e4"},
				{"e7", "e5"},
			},
			pawnPos: d2,
			expected: []Move{
				createMove(b.board, d2, d3, Regular),
				createMove(b.board, d2, d4, Regular),
			},
		},
		{
			moves: [][2]string{
				{"e2", "e4"},
				{"d7", "d5"},
			},
			pawnPos: e4,
			expected: []Move{
				createMove(be4.board, e4, e5, Regular),
				createMove(be4.board, e4, d5, Capture),
			},
		},
		{
			moves: [][2]string{
				{"g2", "g4"},
				{"d7", "d5"},
				{"g4", "g5"},
				{"d5", "d4"},
				{"g5", "g6"},
				{"d4", "d3"},
				{"g6", "h7"},
				{"a7", "a6"},
			},
			pawnPos:  h7,
			expected: createPawnPromotionMoves(White, h7, g8, BlackBishop, CapturePromotion),
		},
	}

	var err error
	for _, row := range table {
		b := NewMailBoxBoard()
		for _, val := range row.moves {
			s, to := val[0], val[1]
			_, err = b.Move(s, to)
			if err != nil {
				t.Error(err)
			}
		}

		got := pawnMoves(row.pawnPos, b.board, b.Context.enPassantSquare)
		if !sameAfterMoveSort(got, row.expected) {
			t.Errorf("got: %v, expected: %v for %s\n",
				printPrettyMoves(got), printPrettyMoves(row.expected), row.pawnPos)
		}
	}
}

func TestBishopMoves(t *testing.T) {
	var piece = WhiteBishop
	bf1 := NewMailBoxBoard()
	bc4 := NewMailBoxBoard()
	bf1.board[f1] = piece
	bc4.board[c4] = piece
	table := []struct {
		moves    [][2]string
		pos      Square
		piece    Piece
		expected []Move
	}{
		{
			moves: [][2]string{
				{"e2", "e4"},
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
			moves: [][2]string{
				{"e2", "e4"},
				{"d7", "d5"},
				{"f1", "c4"},
				{"d5", "d4"},
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
		b := NewMailBoxBoard()
		for _, val := range row.moves {
			s, to := val[0], val[1]
			_, err = b.Move(s, to)
			if err != nil {
				t.Error(err)
			}
		}

		got := bishopMoves(row.pos, b.board)
		if !sameAfterMoveSort(got, row.expected) {
			t.Errorf("got: %v, expected: %v for %s\n",
				printPrettyMoves(got), printPrettyMoves(row.expected), row.pos)
		}
	}
}

func TestKnightMoves(t *testing.T) {
	var piece = WhiteKnight
	var fun = knightMoves
	b := NewMailBoxBoard()
	table := []struct {
		moves    [][2]string
		pos      Square
		piece    Piece
		expected []Move
	}{
		{
			moves: [][2]string{},
			pos:   g1,
			piece: piece,
			expected: []Move{
				createMove(b.board, g1, f3, Regular),
				createMove(b.board, g1, h3, Regular),
			},
		},
		{
			moves: [][2]string{
				{"e2", "e4"},
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
			moves: [][2]string{
				{"f2", "f4"},
				{"e7", "e5"},
				{"f4", "f5"},
				{"e5", "e4"},
				{"f5", "f6"},
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
		b := NewMailBoxBoard()
		for _, val := range row.moves {
			s, to := val[0], val[1]
			_, err = b.Move(s, to)
			if err != nil {
				t.Error(err)
			}
		}

		got := fun(row.pos, b.board)
		if !sameAfterMoveSort(got, row.expected) {
			t.Errorf("got: %v, expected: %v for %s\n",
				printPrettyMoves(got), printPrettyMoves(row.expected), row.pos)
		}
	}
}

func TestRookMoves(t *testing.T) {
	var piece = WhiteRook
	var fun = rookMoves
	b := NewMailBoxBoard()
	table := []struct {
		moves    [][2]string
		pos      Square
		piece    Piece
		expected []Move
	}{
		{
			moves:    [][2]string{},
			pos:      a1,
			piece:    piece,
			expected: []Move{},
		},
		{
			moves: [][2]string{
				{"a2", "a4"},
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
		b := NewMailBoxBoard()
		for _, val := range row.moves {
			s, to := val[0], val[1]
			_, err = b.Move(s, to)
			if err != nil {
				t.Error(err)
			}
		}

		got := fun(row.pos, b.board)
		if !sameAfterMoveSort(got, row.expected) {
			t.Errorf("got: %v, expected: %v for %s\n",
				printPrettyMoves(got), printPrettyMoves(row.expected), row.pos)
		}
	}
}

func TestQueenMoves(t *testing.T) {
	var piece = WhiteQueen
	var fun = queenMoves
	b := NewMailBoxBoard()
	table := []struct {
		moves    [][2]string
		pos      Square
		piece    Piece
		expected []Move
	}{
		{
			moves:    [][2]string{},
			pos:      d1,
			piece:    piece,
			expected: []Move{},
		},
		{
			moves: [][2]string{
				{"e2", "e4"},
				{"h7", "h5"},
				{"d2", "d4"},
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
		b := NewMailBoxBoard()
		for _, val := range row.moves {
			s, to := val[0], val[1]
			_, err = b.Move(s, to)
			if err != nil {
				t.Error(err)
			}
		}

		got := fun(row.pos, b.board)
		if !sameAfterMoveSort(got, row.expected) {
			t.Errorf("got: %v, expected: %v for %s\n",
				printPrettyMoves(got), printPrettyMoves(row.expected), row.pos)
		}
	}
}

func TestKingMoves(t *testing.T) {
	var piece = WhiteKing
	var fun1 = kingMoves
	var fun2 = castleMoves
	b := NewMailBoxBoard()
	table := []struct {
		moves    [][2]string
		pos      Square
		piece    Piece
		expected []Move
	}{
		{
			moves:    [][2]string{},
			pos:      e1,
			piece:    piece,
			expected: []Move{},
		},
		{
			moves: [][2]string{
				{"e2", "e4"},
				{"e7", "e5"},
				{"d2", "d4"},
			},
			pos:   e1,
			piece: piece,
			expected: []Move{
				createMove(b.board, e1, e2, Regular),
				createMove(b.board, e1, d2, Regular),
			},
		},
		{
			moves: [][2]string{
				{"e2", "e4"},
				{"e7", "e5"},
				{"f1", "c4"},
				{"d7", "d5"},
				{"g1", "f3"},
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
			moves: [][2]string{
				{"d2", "d4"},
				{"d7", "d5"},
				{"c1", "f4"},
				{"e7", "e5"},
				{"b1", "c3"},
				{"a7", "a5"},
				{"d1", "d2"},
				{"a8", "a6"},
				{"a1", "b1"},
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
		b := NewMailBoxBoard()
		for _, val := range row.moves {
			s, to := val[0], val[1]
			_, err = b.Move(s, to)
			if err != nil {
				t.Error(err)
			}
		}

		got := fun1(row.pos, b.board)
		got = append(got, fun2(row.pos, b.board, b.Context)...)
		if !sameAfterMoveSort(got, row.expected) {
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

func sameAfterMoveSort(a, b []Move) bool {
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
