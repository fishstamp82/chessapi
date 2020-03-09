package chess

type Board interface {
	CheckMate() bool
	IsCheck() bool
	IsDraw() bool
	Won() (string, error)
	PlayersTurn() string
	BoardMap() map[string]string
	Move(s, t string) (State, error)
}

func (b *MailBoxBoard) IsCheck() bool {
	kingSquare := getKingSquare(b.context.playersTurn, b.board)
	return inCheck(kingSquare, b.board)
}

func NewBoard() Board {
	b := &MailBoxBoard{state: Playing}

	b.context.playersTurn = White
	b.context.whiteCanCastleLeft = true
	b.context.whiteCanCastleRight = true
	b.context.blackCanCastleLeft = true
	b.context.blackCanCastleRight = true
	//Pawns
	for _, s := range []Square{a2, b2, c2, d2, e2, f2, g2, h2} {
		b.board[s] = WhitePawn
	}
	for _, s := range []Square{a7, b7, c7, d7, e7, f7, g7, h7} {
		b.board[s] = BlackPawn
	}
	b.board[a1] = WhiteRook
	b.board[h1] = WhiteRook

	b.board[a8] = BlackRook
	b.board[h8] = BlackRook

	b.board[b1] = WhiteKnight
	b.board[g1] = WhiteKnight

	b.board[b8] = BlackKnight
	b.board[g8] = BlackKnight

	b.board[c1] = WhiteBishop
	b.board[f1] = WhiteBishop

	b.board[c8] = BlackBishop
	b.board[f8] = BlackBishop

	b.board[d1] = WhiteQueen
	b.board[e1] = WhiteKing

	b.board[d8] = BlackQueen
	b.board[e8] = BlackKing

	return b
}
