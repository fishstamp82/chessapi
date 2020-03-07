package chess

var pieceToUnicode = map[Piece]string{
	Empty:       "\u0020",
	WhitePawn:   "\u2659",
	WhiteKnight: "\u2658",
	WhiteBishop: "\u2657",
	WhiteRook:   "\u2656",
	WhiteQueen:  "\u2655",
	WhiteKing:   "\u2654",
	BlackPawn:   "\u265F",
	BlackKnight: "\u265E",
	BlackBishop: "\u265D",
	BlackRook:   "\u265C",
	BlackQueen:  "\u265B",
	BlackKing:   "\u265A",
}
