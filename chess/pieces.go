package chess

type Piece int8

const (
	Empty       Piece = iota
	WhitePawn         // 1
	WhiteKnight       // 2
	WhiteBishop       // 3
	WhiteRook         // 4
	WhiteQueen        // 5
	WhiteKing         // 6
)

const (
	BlackPawn   Piece = -1
	BlackKnight       = -2
	BlackBishop       = -3
	BlackRook         = -4
	BlackQueen        = -5
	BlackKing         = -6
)

var pieceToString = map[Piece]string{
	Empty:       "",
	WhitePawn:   "white pawn",
	WhiteKnight: "white knight",
	WhiteBishop: "white bishop",
	WhiteRook:   "white rook",
	WhiteQueen:  "white queen",
	WhiteKing:   "white king",
	BlackPawn:   "black pawn",
	BlackKnight: "black knight",
	BlackBishop: "black bishop",
	BlackRook:   "black rook",
	BlackQueen:  "black queen",
	BlackKing:   "black king",
}

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
