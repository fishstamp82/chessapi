package chess

func (b *ChessBoard) targets(s Square) []Square {
	p := b.board[s]
	var targets []Square
	switch p {
	case WhitePawn:
		targets = b.pawnTargets(s)
	case BlackPawn:
		targets = b.pawnTargets(s)
	case WhiteBishop:
		targets = b.bishopTargets(s)
	case BlackBishop:
		targets = b.bishopTargets(s)
	case WhiteKnight:
		targets = b.knightTargets(s)
	case BlackKnight:
		targets = b.knightTargets(s)
	case WhiteRook:
		targets = b.rookTargets(s)
	case BlackRook:
		targets = b.rookTargets(s)
	case WhiteQueen:
		targets = b.queenTargets(s)
	case BlackQueen:
		targets = b.queenTargets(s)
	case WhiteKing:
		targets = b.kingTargets(s)
	case BlackKing:
		targets = b.kingTargets(s)
	}
	return targets
}

func (b *ChessBoard) pawnTargets(s Square) []Square {
	var isWhite bool
	switch b.board[s] > 0 {
	case true:
		isWhite = true
	case false:
		isWhite = false
	}

	var moves []Square
	var targets []Square

	p := b.board[s]
	switch p {
	case WhitePawn:
		moves = b.whitePawnMoves(s)
	case BlackPawn:
		moves = b.blackPawnMoves(s)
	default:
		moves = []Square{}
	}
	for _, val := range moves {
		if isWhite && b.board[val] < 0 {
			targets = append(targets, val)
		} else if !isWhite && b.board[val] > 0 {
			targets = append(targets, val)
		}
	}
	return targets
}

func (b *ChessBoard) bishopTargets(s Square) []Square {
	var isWhite bool
	switch b.board[s] > 0 {
	case true:
		isWhite = true
	case false:
		isWhite = false
	}

	var moves []Square
	var targets []Square

	p := b.board[s]
	switch p {
	case WhiteBishop:
		moves = b.bishopMoves(s)
	case BlackBishop:
		moves = b.bishopMoves(s)
	default:
		moves = []Square{}
	}
	for _, val := range moves {
		if isWhite && b.board[val] < 0 {
			targets = append(targets, val)
		} else if !isWhite && b.board[val] > 0 {
			targets = append(targets, val)
		}
	}
	return targets
}

func (b *ChessBoard) knightTargets(s Square) []Square {
	var isWhite bool
	switch b.board[s] > 0 {
	case true:
		isWhite = true
	case false:
		isWhite = false
	}

	var moves []Square
	var targets []Square

	p := b.board[s]
	switch p {
	case WhiteKnight:
		moves = b.knightMoves(s)
	case BlackKnight:
		moves = b.knightMoves(s)
	default:
		moves = []Square{}
	}
	for _, val := range moves {
		if isWhite && b.board[val] < 0 {
			targets = append(targets, val)
		} else if !isWhite && b.board[val] > 0 {
			targets = append(targets, val)
		}
	}
	return targets
}

func (b *ChessBoard) rookTargets(s Square) []Square {
	var isWhite bool
	switch b.board[s] > 0 {
	case true:
		isWhite = true
	case false:
		isWhite = false
	}

	var moves []Square
	var targets []Square

	p := b.board[s]
	switch p {
	case WhiteRook:
		moves = b.rookMoves(s)
	case BlackRook:
		moves = b.rookMoves(s)
	default:
		moves = []Square{}
	}
	for _, val := range moves {
		if isWhite && b.board[val] < 0 {
			targets = append(targets, val)
		} else if !isWhite && b.board[val] > 0 {
			targets = append(targets, val)
		}
	}
	return targets
}

func (b *ChessBoard) queenTargets(s Square) []Square {
	var isWhite bool
	switch b.board[s] > 0 {
	case true:
		isWhite = true
	case false:
		isWhite = false
	}

	var moves []Square
	var targets []Square

	p := b.board[s]
	switch p {
	case WhiteQueen:
		moves = b.queenMoves(s)
	case BlackQueen:
		moves = b.queenMoves(s)
	default:
		moves = []Square{}
	}
	for _, val := range moves {
		if isWhite && b.board[val] < 0 {
			targets = append(targets, val)
		} else if !isWhite && b.board[val] > 0 {
			targets = append(targets, val)
		}
	}
	return targets
}

func (b *ChessBoard) kingTargets(s Square) []Square {
	var isWhite bool
	switch b.board[s] > 0 {
	case true:
		isWhite = true
	case false:
		isWhite = false
	}

	var moves []Square
	var targets []Square

	p := b.board[s]
	switch p {
	case WhiteKing:
		moves = b.kingMoves(s)
	case BlackKing:
		moves = b.kingMoves(s)
	default:
		moves = []Square{}
	}
	for _, val := range moves {
		if isWhite && b.board[val] < 0 {
			targets = append(targets, val)
		} else if !isWhite && b.board[val] > 0 {
			targets = append(targets, val)
		}
	}
	return targets
}
