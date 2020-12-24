package chess

import (
	"testing"
)

func TestPGN(t *testing.T) {

	game := NewGame()
	table := []struct {
		pgnGame       string
		expectedMoves []Move
		expectedFen   string
	}{
		//		{
		//			pgnGame: `[Event "Lloyds Bank op"]
		//[Site "London"]
		//[Date "1984.??.??"]
		//[Round "1"]
		//[White "Adams, Michael"]
		//[Black "Sedgwick, David"]
		//[Result "1-0"]
		//[WhiteElo ""]
		//[BlackElo ""]
		//[ECO "C05"]
		//
		//1.e4 e6 2.d4 d5 3.Nd2 Nf6 4.e5 Nfd7 5.f4 c5 6.c3 Nc6 7.Ndf3 cxd4 8.cxd4 f6
		//9.Bd3 Bb4+ 10.Bd2 Qb6 11.Ne2 fxe5 12.fxe5 O-O 13.a3 Be7 14.Qc2 Rxf3 15.gxf3 Nxd4
		//16.Nxd4 Qxd4 17.O-O-O Nxe5 18.Bxh7+ Kh8 19.Kb1 Qh4 20.Bc3 Bf6 21.f4 Nc4 22.Bxf6 Qxf6
		//23.Bd3 b5 24.Qe2 Bd7 25.Rhg1 Be8 26.Rde1 Bf7 27.Rg3 Rc8 28.Reg1 Nd6 29.Rxg7 Nf5
		//30.R7g5 Rc7 31.Bxf5 exf5 32.Rh5+  1-0`,
		//			expectedMoves: []Move{
		//				//createMove()
		//			},
		//			expectedFen: "7k/p1r2b2/5q2/1p1p1p1R/5P2/P7/1P2Q2P/1K4R1 b - - 1 32",
		//		},
		{
			pgnGame: `1.e4 e6`,
			expectedMoves: []Move{
				createMove(game.Board.board, e2, e4, []MovementType{Regular, PawnMove}),
				createMove(game.Board.board, e7, e6, []MovementType{Regular, PawnMove}),
			},
			expectedFen: "rnbqkbnr/pppp1ppp/4p3/8/4P3/8/PPPP1PPP/RNBQKBNR w KQkq - 0 2",
		},
	}
	for _, row := range table {
		for _, move := range row.expectedMoves {
			_, _ = game.move(move.fromSquare, move.toSquare)
		}
		//got, _ := pgnParse(strings.NewReader(row.pgnGame))
		if game.FenString() != row.expectedFen {
			t.Errorf("got: %v, expected: %v\n", game.FenString(), row.expectedFen)
		}
	}
}

func TestFilterComments(t *testing.T) {
	table := []struct {
		pgnGame        string
		expectedString string
	}{
		{
			pgnGame: `[Event "Lloyds Bank op"]
[Site "London"]
[Date "1984.??.??"]
[Round "1"]
[White "Adams, Michael"]
[Black "Sedgwick, David"]
[Result "1-0"]
[WhiteElo ""]
[BlackElo ""]
[ECO "C05"]
[]

1.e4 e6 2.d4 d5 3.Nd2 Nf6 4.e5 Nfd7 5.f4 c5 6.c3 Nc6 7.Ndf3 cxd4 8.cxd4 f6
9.Bd3 Bb4+ 10.Bd2 Qb6 11.Ne2 fxe5 12.fxe5 O-O 13.a3 Be7 14.Qc2 Rxf3 15.gxf3 Nxd4
16.Nxd4 Qxd4 17.O-O-O Nxe5 18.Bxh7+ Kh8 19.Kb1 Qh4 20.Bc3 Bf6 21.f4 Nc4 22.Bxf6 Qxf6
23.Bd3 b5 24.Qe2 Bd7 25.Rhg1 Be8 26.Rde1 Bf7 27.Rg3 Rc8 28.Reg1 Nd6 29.Rxg7 Nf5
30.R7g5 Rc7 31.Bxf5 exf5 32.Rh5+  1-0`,
			expectedString: `1.e4 e6 2.d4 d5 3.Nd2 Nf6 4.e5 Nfd7 5.f4 c5 6.c3 Nc6 7.Ndf3 cxd4 8.cxd4 f6
9.Bd3 Bb4+ 10.Bd2 Qb6 11.Ne2 fxe5 12.fxe5 O-O 13.a3 Be7 14.Qc2 Rxf3 15.gxf3 Nxd4
16.Nxd4 Qxd4 17.O-O-O Nxe5 18.Bxh7+ Kh8 19.Kb1 Qh4 20.Bc3 Bf6 21.f4 Nc4 22.Bxf6 Qxf6
23.Bd3 b5 24.Qe2 Bd7 25.Rhg1 Be8 26.Rde1 Bf7 27.Rg3 Rc8 28.Reg1 Nd6 29.Rxg7 Nf5
30.R7g5 Rc7 31.Bxf5 exf5 32.Rh5+  1-0`,
		},
	}

	for _, row := range table {
		got := filterMoves(row.pgnGame)
		if got != row.expectedString {
			t.Errorf("got: %v, expected: %v\n", got, row.expectedString)
		}
	}
}

func TestFindFromSquare(t *testing.T) {
	startFen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	table := []struct {
		piece               Piece
		toSquare            Square
		game                *Game
		expectedFromSquares []Square
	}{
		{
			WhitePawn,
			e4,
			NewGameFromFEN(startFen),
			[]Square{e2},
		},
		{
			WhiteRook,
			d1,
			NewGameFromFEN("8/8/8/8/8/8/7K/R6R w - - 0 1"),
			[]Square{a1, h1},
		},
	}

	for _, row := range table {
		got := findFromSquares(row.piece, row.toSquare, row.game.Board.board, row.game.Context)
		if !areSquaresEqual(got, row.expectedFromSquares) {
			t.Errorf("got: %v, expected: %v\n", got, row.expectedFromSquares)
		}
	}
}

func TestIsCapture(t *testing.T) {
	table := []struct {
		playerMove string
		expected   bool
	}{
		{
			"exd5",
			true,
		},
		{
			"Ng3",
			false,
		},
	}

	for _, row := range table {
		got := isCapture(row.playerMove)
		if got != row.expected {
			t.Errorf("got: %v, expected: %v\n", got, row.expected)
		}
	}
}

func TestIsPawn(t *testing.T) {
	table := []struct {
		playerMove string
		expected   bool
	}{
		{
			"exd5",
			true,
		},
		{
			"Ng3",
			false,
		},
		{
			"f3",
			true,
		},
		{
			"O-O",
			false,
		},
	}

	for _, row := range table {
		got := isPawn(row.playerMove)
		if got != row.expected {
			t.Errorf("got: %v, expected: %v\n", got, row.expected)
		}
	}
}

func TestIsCastle(t *testing.T) {
	table := []struct {
		playerMove string
		expected   bool
	}{
		{
			"exd5",
			false,
		},
		{
			"Ng3",
			false,
		},
		{
			"f3",
			false,
		},
		{
			"dxe8=Q",
			false,
		},
		{
			"O-O",
			true,
		},
	}

	for _, row := range table {
		got := isCastle(row.playerMove)
		if got != row.expected {
			t.Errorf("got: %v, expected: %v\n", got, row.expected)
		}
	}
}

func TestIsPromotion(t *testing.T) {
	table := []struct {
		playerMove string
		expected   bool
	}{
		{
			"exd5",
			false,
		},
		{
			"Ng3",
			false,
		},
		{
			"a8=N",
			true,
		},
		{
			"f3",
			false,
		},
		{
			"dxe8=Q",
			true,
		},
		{
			"O-O",
			false,
		},
	}

	for _, row := range table {
		got := isPromotion(row.playerMove)
		if got != row.expected {
			t.Errorf("got: %v, expected: %v\n", got, row.expected)
		}
	}
}

func TestDecodeCastleMust(t *testing.T) {
	table := []struct {
		player     Color
		playerMove string
		expected   [2]Square
	}{
		{
			White,
			"O-O",
			[2]Square{e1, g1},
		},
		{
			White,
			"O-O-O",
			[2]Square{e1, c1},
		},
		{
			Black,
			"O-O",
			[2]Square{e8, g8},
		},
		{
			Black,
			"O-O-O",
			[2]Square{e8, c8},
		},
	}

	for _, row := range table {
		got1, got2 := decodeCastleMust(row.player, row.playerMove)
		if [2]Square{got1, got2} != row.expected {
			t.Errorf("got: %v, expected: %v\n", [2]Square{got1, got2}, row.expected)
		}
	}
}

func TestGetPieceMust(t *testing.T) {
	table := []struct {
		piece         Piece
		player        Color
		expectedPiece Piece
	}{
		{
			Bishop,
			Black,
			BlackBishop,
		},
		{
			Knight,
			Black,
			BlackKnight,
		},
		{
			Rook,
			Black,
			BlackRook,
		},
		{
			Queen,
			Black,
			BlackQueen,
		},
		{
			King,
			Black,
			BlackKing,
		},
		{
			Bishop,
			White,
			WhiteBishop,
		},
		{
			Knight,
			White,
			WhiteKnight,
		},
		{
			Rook,
			White,
			WhiteRook,
		},
		{
			Queen,
			White,
			WhiteQueen,
		},
		{
			King,
			White,
			WhiteKing,
		},
	}

	for _, row := range table {
		got := getPieceMust(row.piece, row.player)
		if got != row.expectedPiece {
			t.Errorf("got: %v, expected: %v\n", got, row.expectedPiece)
		}
	}
}

func TestGetFileRank(t *testing.T) {
	table := []struct {
		fromInformation string
		expectedFile    byte
		expectedRank    byte
	}{
		{
			"e1",
			'e',
			'1',
		},
		{
			"d",
			'd',
			0,
		},
		{
			"8",
			0,
			'8',
		},
		{
			"",
			0,
			0,
		},
	}

	for _, row := range table {
		gotFile, gotRank := getFileRank(row.fromInformation)
		if gotFile != row.expectedFile || gotRank != row.expectedRank {
			t.Errorf("got: %v %v, expected: %v %v\n", gotFile, gotRank, row.expectedFile, row.expectedRank)
		}
	}
}

func TestDisambiguateMust(t *testing.T) {
	table := []struct {
		squares        []Square
		file           byte
		rank           byte
		expectedSquare Square
	}{
		{
			squares:        []Square{e3, c3},
			file:           'c',
			rank:           0,
			expectedSquare: c3,
		},
		{
			squares:        []Square{a1, h1},
			file:           'h',
			rank:           0,
			expectedSquare: h1,
		},
		{
			squares:        []Square{a1, a8},
			file:           0,
			rank:           '8',
			expectedSquare: a8,
		},
		{
			squares:        []Square{e4, h4, h1},
			file:           'h',
			rank:           '4',
			expectedSquare: h4,
		},
	}

	for _, row := range table {
		gotSquare := disambiguateMust(row.squares, row.file, row.rank)
		if gotSquare != row.expectedSquare {
			t.Errorf("got: %s, expected: %s\n", gotSquare, row.expectedSquare)
		}
	}
}

func TestParseNotation(t *testing.T) {
	fen0 := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	fen1 := "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w KQkq - 0 1"
	fen3 := "3r3r/b2k4/3b4/R7/2K1Q2Q/8/8/R6Q w KQkq - 0 1"

	table := []struct {
		player       Color
		notation     string
		board        [64]Piece
		context      Context
		expectedMove Move
	}{
		{
			player:       White,
			notation:     "e4",
			board:        NewGameFromFEN(fen0).Board.board,
			context:      NewGameFromFEN(fen0).Context,
			expectedMove: createPawnMove(WhitePawn, e2, e4, []MovementType{PawnMove}),
		},
		{
			player:       Black,
			notation:     "Nf6",
			board:        NewGameFromFEN(fen1).Board.board,
			context:      NewGameFromFEN(fen1).Context,
			expectedMove: createMove(NewGameFromFEN(fen1).Board.board, g8, f6, []MovementType{Regular}),
		},
		{
			player:       Black,
			notation:     "Bdb8",
			board:        NewGameFromFEN(fen3).Board.board,
			context:      NewGameFromFEN(fen3).Context,
			expectedMove: createMove(NewGameFromFEN(fen3).Board.board, d6, b8, []MovementType{Regular}),
		},
		{
			player:       Black,
			notation:     "Rdf8",
			board:        NewGameFromFEN(fen3).Board.board,
			context:      NewGameFromFEN(fen3).Context,
			expectedMove: createMove(NewGameFromFEN(fen3).Board.board, d8, f8, []MovementType{Regular}),
		},
		{
			player:       White,
			notation:     "R1a3",
			board:        NewGameFromFEN(fen3).Board.board,
			context:      NewGameFromFEN(fen3).Context,
			expectedMove: createMove(NewGameFromFEN(fen3).Board.board, a1, a3, []MovementType{Regular}),
		},
		{
			player:       White,
			notation:     "Qh4e1",
			board:        NewGameFromFEN(fen3).Board.board,
			context:      NewGameFromFEN(fen3).Context,
			expectedMove: createMove(NewGameFromFEN(fen3).Board.board, h4, e1, []MovementType{Regular}),
		},
	}

	for _, row := range table {
		gotMove := parseNotation(row.player, row.notation, row.board, row.context)
		if !isMoveEqual(gotMove, row.expectedMove) {
			t.Errorf("got: %s, expected: %s\n", gotMove, row.expectedMove)
		}
	}
}

func areSquaresEqual(a, b []Square) bool {
	if len(a) != len(b) {
		return false

	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
