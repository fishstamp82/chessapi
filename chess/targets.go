package chess

func (b *Board) pawnTargets(s Square) []Square {
	var isWhite bool
	switch b.board[s] > 0 {
	case true:
		isWhite = true
	case false:
		isWhite = false
	}

	var moves []Square
	targets := []Square{}

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

func (b *Board) bishopTargets(s Square) []Square {
	var isWhite bool
	switch b.board[s] > 0 {
	case true:
		isWhite = true
	case false:
		isWhite = false
	}

	var moves []Square
	targets := []Square{}

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

func (b *Board) knightTargets(s Square) []Square {
	var isWhite bool
	switch b.board[s] > 0 {
	case true:
		isWhite = true
	case false:
		isWhite = false
	}

	var moves []Square
	targets := []Square{}

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

func (b *Board) rookTargets(s Square) []Square {
	var isWhite bool
	switch b.board[s] > 0 {
	case true:
		isWhite = true
	case false:
		isWhite = false
	}

	var moves []Square
	targets := []Square{}

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

func (b *Board) queenTargets(s Square) []Square {
	var isWhite bool
	switch b.board[s] > 0 {
	case true:
		isWhite = true
	case false:
		isWhite = false
	}

	var moves []Square
	targets := []Square{}

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
