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
	table := []struct {
		moves     [][2]string
		bishopPos Square
		expected  []Move
	}{
		{
			moves: [][2]string{
				{"e2", "e4"},
			},
			bishopPos: f1,
			expected: []Move{
				makeMove(WhitePawn, f1, e2, Regular),
				makeMove(WhitePawn, f1, d3, Regular),
				makeMove(WhitePawn, f1, c4, Regular),
				makeMove(WhitePawn, f1, b5, Regular),
				makeMove(WhitePawn, f1, a6, Regular),
			},
		},
		{
			moves: [][2]string{
				{"e2", "e4"},
				{"d7", "d5"},
				{"f1", "c4"},
				{"d5", "d4"},
			},
			bishopPos: c4,
			expected: []Move{
				makeMove(WhitePawn, c4, f1, Regular),
				makeMove(WhitePawn, c4, e2, Regular),
				makeMove(WhitePawn, c4, b3, Regular),
				makeMove(WhitePawn, c4, d3, Regular),
				makeMove(WhitePawn, c4, b5, Regular),
				makeMove(WhitePawn, c4, d5, Regular),
				makeMove(WhitePawn, c4, a6, Regular),
				makeMove(WhitePawn, c4, e6, Regular),
				makeMove(WhitePawn, c4, f7, Capture),
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

		got := bishopMoves(row.bishopPos, b.board)
		if !sameAfterMoveSort(got, row.expected) {
			t.Errorf("got: %v, expected: %v for %s\n",
				printPrettyMoves(got), printPrettyMoves(row.expected), row.bishopPos)
		}
	}
}

//
//func TestRookMoves(t *testing.T) {
//	table := []struct {
//		whiteRook Square
//		blackRook Square
//		whitePawn Square
//		expected  []Square
//	}{
//		{
//			whiteRook: h1,
//			blackRook: a2,
//			whitePawn: d5,
//			expected:  []Square{a1, b1, c1, d1, e1, f1, g1, h2, h3, h4, h5, h6, h7, h8},
//		},
//		{
//			whiteRook: e1,
//			blackRook: e5,
//			whitePawn: d5,
//			expected:  []Square{a1, b1, c1, d1, e2, e3, e4, e5, f1, g1, h1},
//		},
//		{
//			whiteRook: e2,
//			blackRook: e5,
//			whitePawn: d2,
//			expected:  []Square{e1, e3, e4, e5, f2, g2, h2},
//		},
//	}
//
//	for _, row := range table {
//		b := NewEmptyMailBoxBoard()
//		b.board[row.whiteRook] = WhiteRook
//		b.board[row.whitePawn] = WhitePawn
//		b.board[row.blackRook] = BlackRook
//		got := rookMoves(row.whiteRook, b.board)
//
//		if !sameAfterSquareSort(got, row.expected) {
//			t.Errorf("got: %v, expected: %v for %s on %s\n",
//				printPrettySquares(got), printPrettySquares(row.expected), WhiteRook, squareToString[row.whiteRook])
//		}
//	}
//}
//
//func TestQueenMoves(t *testing.T) {
//	table := []struct {
//		whiteQueen Square
//		blackQueen Square
//		whitePawn  Square
//		expected   []Square
//	}{
//		{
//			whiteQueen: a8,
//			blackQueen: a1,
//			whitePawn:  a3,
//			expected:   []Square{a4, a5, a6, a7, b8, c8, d8, e8, f8, g8, h8, b7, c6, d5, e4, f3, g2, h1},
//		},
//	}
//
//	for _, row := range table {
//		b := NewEmptyMailBoxBoard()
//		b.board[row.whiteQueen] = WhiteQueen
//		b.board[row.whitePawn] = WhitePawn
//		b.board[row.blackQueen] = BlackQueen
//		got := queenMoves(row.whiteQueen, b.board)
//
//		if !sameAfterSquareSort(got, row.expected) {
//			t.Errorf("got: %v, expected: %v for %s on %s\n",
//				printPrettySquares(got), printPrettySquares(row.expected), WhiteQueen, squareToString[row.whiteQueen])
//		}
//	}
//}
//
//func TestBlackQueenMovesBeginningPosition(t *testing.T) {
//	table := []struct {
//		expected []Square
//	}{
//		{
//			expected: []Square{},
//		},
//	}
//
//	for _, row := range table {
//		b := NewMailBoxBoard()
//		got := queenMoves(d8, b.board)
//
//		if !sameAfterSquareSort(got, row.expected) {
//			t.Errorf("got: %v, expected: %v for %s on %s\n",
//				printPrettySquares(got), printPrettySquares(row.expected), WhiteQueen, squareToString[d8])
//		}
//	}
//}
//
//func TestKnightMoves(t *testing.T) {
//	table := []struct {
//		whiteKnight Square
//		blackKnight Square
//		expected    []Square
//		fullBoard   bool
//	}{
//		{
//			whiteKnight: a1,
//			blackKnight: a2,
//			expected:    []Square{b3, c2},
//			fullBoard:   false,
//		},
//		{
//			whiteKnight: e3,
//			blackKnight: f5,
//			expected:    []Square{d5, f5, g2, g4, d1, f1, c2, c4},
//			fullBoard:   false,
//		},
//		{
//			whiteKnight: g1,
//			blackKnight: g8,
//			expected:    []Square{f3, h3},
//			fullBoard:   true,
//		},
//	}
//	var b *MailBoxBoard
//	for _, row := range table {
//		if row.fullBoard {
//			b = NewMailBoxBoard()
//		} else {
//			b = NewEmptyMailBoxBoard()
//		}
//
//		b.board[row.whiteKnight] = WhiteKnight
//		b.board[row.blackKnight] = BlackKnight
//		got := knightMoves(row.whiteKnight, b.board)
//
//		if !sameAfterSquareSort(got, row.expected) {
//			t.Errorf("got: %v, expected: %v for %s on %s\n",
//				printPrettySquares(got), printPrettySquares(row.expected), WhiteKnight, squareToString[row.whiteKnight])
//		}
//	}
//}
//
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
		if (a[i].toSquare != b[i].toSquare) || a[i].fromSquare != b[i].fromSquare || a[i].moveType != b[i].moveType {
			return false
		}
	}
	return true
}
