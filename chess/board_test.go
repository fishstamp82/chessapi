package chess

import (
	"testing"
)

func TestInitialization(t *testing.T) {
	game := NewGame()
	if game.Context.ColorsTurn != White {
		t.Error("not white's turn on start of the game")
	}

	table := []struct {
		square Square
		piece  Piece
	}{
		{a2, WhitePawn},
		{b2, WhitePawn},
		{c2, WhitePawn},
		{d2, WhitePawn},
		{e2, WhitePawn},
		{f2, WhitePawn},
		{g2, WhitePawn},
		{h2, WhitePawn},
		{a1, WhiteRook},
		{b1, WhiteKnight},
		{c1, WhiteBishop},
		{d1, WhiteQueen},
		{e1, WhiteKing},
		{f1, WhiteBishop},
		{g1, WhiteKnight},
		{h1, WhiteRook},
		{a7, BlackPawn},
		{b7, BlackPawn},
		{c7, BlackPawn},
		{d7, BlackPawn},
		{e7, BlackPawn},
		{f7, BlackPawn},
		{g7, BlackPawn},
		{h7, BlackPawn},
		{a8, BlackRook},
		{b8, BlackKnight},
		{c8, BlackBishop},
		{d8, BlackQueen},
		{e8, BlackKing},
		{f8, BlackBishop},
		{g8, BlackKnight},
		{h8, BlackRook},
		{a5, Empty},
	}

	for _, row := range table {
		p := row.piece
		s := row.square
		if piece := game.Board.board[s]; piece != p {
			t.Errorf("expected: %s on %s, got: %s\n", p, s, piece)
		}
	}

}

func TestEnPassant(t *testing.T) {
	var piece = WhitePawn
	table := []struct {
		moves         []string
		expectedMoves []Move
	}{
		{
			moves: []string{
				"e2e4",
				"d7d5",
				"e4e5",
				"f7f5",
			},
			expectedMoves: []Move{
				createPawnMove(piece, e5, e6, []MovementType{Regular, PawnMove}),
				createPawnEnPassantMove(piece, e5, f6, []MovementType{CaptureEnPassant, PawnMove}),
			},
		},
	}

	for _, row := range table {
		game := NewGame()
		game.Context.State = Playing

		for _, move := range row.moves {
			_, err := game.Move(move)
			if err != nil {
				t.Error(err)
			}
		}
		moves := pawnMoves(e5, game.Board.board, game.Context.enPassantSquare)
		if !isMovesEqual(moves, row.expectedMoves) {
			t.Errorf("got: %s, expected: %s\n", printPrettyMoves(moves), printPrettyMoves(row.expectedMoves))
		}
	}
}

func TestNewFromFEN(t *testing.T) {
	startingFen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	type piecePosition struct {
		position Square
		piece    Piece
	}
	table := []struct {
		fen            string
		piecePositions []piecePosition
		context        Context
	}{
		{
			fen: startingFen,
			piecePositions: []piecePosition{
				{a2, WhitePawn},
				{b2, WhitePawn},
				{c2, WhitePawn},
				{d2, WhitePawn},
				{e2, WhitePawn},
				{f2, WhitePawn},
				{g2, WhitePawn},
				{h2, WhitePawn},
				{a1, WhiteRook},
				{b1, WhiteKnight},
				{c1, WhiteBishop},
				{d1, WhiteQueen},
				{e1, WhiteKing},
				{f1, WhiteBishop},
				{g1, WhiteKnight},
				{h1, WhiteRook},
				{a7, BlackPawn},
				{b7, BlackPawn},
				{c7, BlackPawn},
				{d7, BlackPawn},
				{e7, BlackPawn},
				{f7, BlackPawn},
				{g7, BlackPawn},
				{h7, BlackPawn},
				{a8, BlackRook},
				{b8, BlackKnight},
				{c8, BlackBishop},
				{d8, BlackQueen},
				{e8, BlackKing},
				{f8, BlackBishop},
				{g8, BlackKnight},
				{h8, BlackRook},
			},
			context: Context{
				State:               Playing,
				ColorsTurn:          White,
				Winner:              Noone,
				Score:               "",
				whiteCanCastleRight: true,
				whiteCanCastleLeft:  true,
				blackCanCastleRight: true,
				blackCanCastleLeft:  true,
				enPassantSquare:     none,
				halfMove:            0,
				fullMove:            1,
			},
		},
		{
			fen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNB1KBNR w KQkq - 0 1",
			piecePositions: []piecePosition{
				{a2, WhitePawn},
				{b2, WhitePawn},
				{c2, WhitePawn},
				{d2, WhitePawn},
				{e2, WhitePawn},
				{f2, WhitePawn},
				{g2, WhitePawn},
				{h2, WhitePawn},
				{a1, WhiteRook},
				{b1, WhiteKnight},
				{c1, WhiteBishop},
				{e1, WhiteKing},
				{f1, WhiteBishop},
				{g1, WhiteKnight},
				{h1, WhiteRook},
				{a7, BlackPawn},
				{b7, BlackPawn},
				{c7, BlackPawn},
				{d7, BlackPawn},
				{e7, BlackPawn},
				{f7, BlackPawn},
				{g7, BlackPawn},
				{h7, BlackPawn},
				{a8, BlackRook},
				{b8, BlackKnight},
				{c8, BlackBishop},
				{d8, BlackQueen},
				{e8, BlackKing},
				{f8, BlackBishop},
				{g8, BlackKnight},
				{h8, BlackRook},
			},
			context: Context{
				State:               Playing,
				ColorsTurn:          White,
				Winner:              Noone,
				Score:               "",
				whiteCanCastleRight: true,
				whiteCanCastleLeft:  true,
				blackCanCastleRight: true,
				blackCanCastleLeft:  true,
				enPassantSquare:     none,
				halfMove:            0,
				fullMove:            1,
			},
		},
		{
			fen: "8/8/8/8/8/8/8/7R w KQkq - 0 1",
			piecePositions: []piecePosition{
				{h1, WhiteRook},
			},
			context: Context{
				State:               Playing,
				ColorsTurn:          White,
				Winner:              Noone,
				Score:               "",
				whiteCanCastleRight: true,
				whiteCanCastleLeft:  true,
				blackCanCastleRight: true,
				blackCanCastleLeft:  true,
				enPassantSquare:     none,
				halfMove:            0,
				fullMove:            1,
			},
		},
		{
			fen: "8/8/8/8/4P3/8/8/8 w - e3 55 60",
			piecePositions: []piecePosition{
				{e4, WhitePawn},
			},
			context: Context{
				State:               Playing,
				ColorsTurn:          White,
				Winner:              Noone,
				Score:               "",
				whiteCanCastleRight: false,
				whiteCanCastleLeft:  false,
				blackCanCastleRight: false,
				blackCanCastleLeft:  false,
				enPassantSquare:     e3,
				halfMove:            55,
				fullMove:            60,
			},
		},
	}

	for _, row := range table {
		game := NewGameFromFEN(row.fen)
		game.Context.State = Playing
		expectedGame := NewEmptyGame()
		for _, pp := range row.piecePositions {
			expectedGame.Board.board[pp.position] = pp.piece
		}

		if game.Context != row.context {
			t.Errorf("got: %s, expected: %s\n", game.Context, row.context)
		}

		if game.Board.board != expectedGame.Board.board {
			t.Errorf("got: %s, expected: %s\n", game.FenString(), expectedGame.FenString())

		}
	}
}
