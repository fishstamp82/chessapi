package chess

type Piece int8

const (
	WhitePawn   Piece = iota + 1 // 1
	WhiteKnight                  // 2
	WhiteBishop                  // 3
	WhiteRook                    // 4
	WhiteQueen                   // 5
	WhiteKing                    // 6
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
