package chess

type PlayBoard interface {
	ValidMoves(string) map[string]Move //key is algebraic notation or e2e4 type
	Move(Move) Context
	BoardMap() map[Square]Piece
}

type PBoard struct {
	board   [64]Piece
	moves   []Move
	Context Context
}

func (b *PBoard) Move(m Move) Context {
	return b.move(m)
}

func (b *PBoard) move(m Move) Context {
	var opponent Player
	b.board = makeMove(m, b.board)

	opponentsKing := getKingSquareMust(opponent, b.board)
	if inCheck(opponentsKing, b.board) {
		b.Context.State = Check
	} else {
		b.Context.State = Playing
	}

	if isCheckMated(opponentsKing, b.board) {
		b.Context.State = CheckMate
		b.Context.Winner = b.Context.PlayersTurn
		return b.Context
	}

	if isDraw(opponent, b.board, b.Context) {
		b.Context.State = Draw
		b.Context.Winner = Both
		return b.Context
	}

	abortCastling(m, b.Context)
	b.Context.enPassantSquare = getEnPassantSquare(m)
	if b.Context.PlayersTurn == Black {
		b.Context.fullMove += 1
	}
	isPawnMove := false
	for _, moveType := range m.moveTypes {
		if moveType == PawnMove {
			isPawnMove = true
		}
	}
	if isPawnMove {
		b.Context.halfMove = 0
	} else {
		b.Context.halfMove += 1
	}
	b.Context.PlayersTurn = b.Context.PlayersTurn
	return b.Context
}

func NewPBoard(fen string) *PBoard {
	return &PBoard{}
}
