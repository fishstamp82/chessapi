package chess

import (
	"fmt"
	"sort"
	"testing"
)

func TestPawnMoves(t *testing.T) {
	table := []struct {
		moves    [][2]string
		pawnPos  Square
		expected []Move
	}{
		{
			moves: [][2]string{
				{"e2", "e4"},
			},
			pawnPos: d2,
			expected: []Move{
				makeMove(WhitePawn, d2, d3, Regular),
				makeMove(WhitePawn, d2, d4, Regular),
			},
		},
		{
			moves: [][2]string{
				{"e2", "e4"},
				{"d7", "d5"},
			},
			pawnPos: e4,
			expected: []Move{
				makeMove(WhitePawn, e4, e5, Regular),
				makeMove(WhitePawn, e4, d5, Capture),
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
			},
			pawnPos:  h7,
			expected: makePawnPromotionMoves(White, h7, g8, CapturePromotion),
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

		got, _ := pawnMoves(row.pawnPos, b.board)
		if !sameAfterMoveSort(got, row.expected) {
			t.Errorf("got: %v, expected: %v for %s\n",
				printPrettyMoves(got), printPrettyMoves(row.expected), row.pawnPos)
		}
	}
}

func TestBishopMoves(t *testing.T) {
	var piece = WhiteBishop
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
				makeMove(piece, f1, e2, Regular),
				makeMove(piece, f1, d3, Regular),
				makeMove(piece, f1, c4, Regular),
				makeMove(piece, f1, b5, Regular),
				makeMove(piece, f1, a6, Regular),
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
				makeMove(piece, c4, f1, Regular),
				makeMove(piece, c4, e2, Regular),
				makeMove(piece, c4, b3, Regular),
				makeMove(piece, c4, d3, Regular),
				makeMove(piece, c4, b5, Regular),
				makeMove(piece, c4, d5, Regular),
				makeMove(piece, c4, a6, Regular),
				makeMove(piece, c4, e6, Regular),
				makeMove(piece, c4, f7, Capture),
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
				makeMove(piece, g1, f3, Regular),
				makeMove(piece, g1, h3, Regular),
			},
		},
		{
			moves: [][2]string{
				{"e2", "e4"},
			},
			pos:   g1,
			piece: piece,
			expected: []Move{
				makeMove(piece, g1, f3, Regular),
				makeMove(piece, g1, h3, Regular),
				makeMove(piece, g1, e2, Regular),
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
				makeMove(BlackKnight, g8, f6, Capture),
				makeMove(BlackKnight, g8, h6, Regular),
				makeMove(BlackKnight, g8, e7, Regular),
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
				makeMove(piece, a1, a2, Regular),
				makeMove(piece, a1, a3, Regular),
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
				makeMove(piece, d1, d2, Regular),
				makeMove(piece, d1, d3, Regular),
				makeMove(piece, d1, e2, Regular),
				makeMove(piece, d1, f3, Regular),
				makeMove(piece, d1, g4, Regular),
				makeMove(piece, d1, h5, Capture),
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
	var fun = kingMoves
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
				makeMove(piece, e1, e2, Regular),
				makeMove(piece, e1, d2, Regular),
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
				makeMove(piece, e1, e2, Regular),
				makeMove(piece, e1, f1, Regular),
				//makeMove(piece, e1, g1, ShortCastle),
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

//func TestKingMoves(t *testing.T) {
//	table := []struct {
//		whiteKing Square
//		blackKing Square
//		expected  []Square
//		fullBoard bool
//	}{
//		{
//			whiteKing: a1,
//			blackKing: a2,
//			expected:  []Square{a2, b1, b2},
//			fullBoard: false,
//		},
//		{
//			whiteKing: e2,
//			blackKing: d8,
//			expected:  []Square{d1, d2, d3, e1, e3, f1, f2, f3},
//			fullBoard: false,
//		},
//	}
//	var b *MailBoxBoard
//	var got []Square
//	for _, row := range table {
//		if row.fullBoard {
//			b = NewMailBoxBoard()
//		} else {
//			b = NewEmptyMailBoxBoard()
//		}
//
//		b.board[row.whiteKing] = WhiteKing
//		b.board[row.blackKing] = BlackKing
//		got = whiteKingMoves(row.whiteKing, b.board)
//
//		if !sameAfterSquareSort(got, row.expected) {
//			t.Errorf("got: %v, expected: %v for %s on %s\n",
//				printPrettySquares(got), printPrettySquares(row.expected), WhiteKing, squareToString[row.whiteKing])
//		}
//	}
//}

//func TestWhiteKingCastle(t *testing.T) {
//	type piecePosition struct {
//		position Square
//		piece    Piece
//	}
//	table := []struct {
//		pieces   []piecePosition
//		expected []Square
//	}{
//		{
//			pieces: []piecePosition{
//				{e1, WhiteKing},
//				{h1, WhiteRook},
//				{a1, WhiteRook},
//			},
//			expected: []Square{c1, g1},
//		},
//	}
//	var b *MailBoxBoard
//	var got []Square
//	for _, row := range table {
//		b = NewEmptyMailBoxBoard()
//		for _, val := range row.pieces {
//			b.board[val.position] = val.piece
//		}
//		b.context.whiteCanCastleRight = true
//		b.context.whiteCanCastleLeft = true
//
//		got = append(got, whiteKingCastleMoves(e1, b.board, b.context)...)
//
//		if !sameAfterSquareSort(got, row.expected) {
//			t.Errorf("got: %v, expected: %v for %s on %s\n",
//				printPrettySquares(got), printPrettySquares(row.expected), WhiteKing, squareToString[e1])
//		}
//	}
//}

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
