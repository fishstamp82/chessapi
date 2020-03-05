package chess

func targets(s Square, b [64]Piece) []Square {
	p := b[s]
	var targets []Square
	switch p {
	case WhitePawn:
		targets = pawnTargets(s, b)
	case BlackPawn:
		targets = pawnTargets(s, b)
	case WhiteBishop:
		targets = bishopTargets(s, b)
	case BlackBishop:
		targets = bishopTargets(s, b)
	case WhiteKnight:
		targets = knightTargets(s, b)
	case BlackKnight:
		targets = knightTargets(s, b)
	case WhiteRook:
		targets = rookTargets(s, b)
	case BlackRook:
		targets = rookTargets(s, b)
	case WhiteQueen:
		targets = queenTargets(s, b)
	case BlackQueen:
		targets = queenTargets(s, b)
	case WhiteKing:
		targets = kingTargets(s, b)
	case BlackKing:
		targets = kingTargets(s, b)
	}
	return targets
}

func pawnTargets(s Square, b [64]Piece) []Square {
	var isWhite bool
	switch b[s] > 0 {
	case true:
		isWhite = true
	case false:
		isWhite = false
	}

	var moves []Square
	var targets []Square

	p := b[s]
	switch p {
	case WhitePawn:
		moves = whitePawnMoves(s, b)
	case BlackPawn:
		moves = blackPawnMoves(s, b)
	default:
		moves = []Square{}
	}
	for _, val := range moves {
		if isWhite && b[val] < 0 {
			targets = append(targets, val)
		} else if !isWhite && b[val] > 0 {
			targets = append(targets, val)
		}
	}
	return targets
}

func bishopTargets(s Square, b [64]Piece) []Square {
	var isWhite bool
	switch b[s] > 0 {
	case true:
		isWhite = true
	case false:
		isWhite = false
	}

	var moves []Square
	var targets []Square

	p := b[s]
	switch p {
	case WhiteBishop:
		moves = bishopMoves(s, b)
	case BlackBishop:
		moves = bishopMoves(s, b)
	default:
		moves = []Square{}
	}
	for _, val := range moves {
		if isWhite && b[val] < 0 {
			targets = append(targets, val)
		} else if !isWhite && b[val] > 0 {
			targets = append(targets, val)
		}
	}
	return targets
}

func knightTargets(s Square, b [64]Piece) []Square {
	var isWhite bool
	switch b[s] > 0 {
	case true:
		isWhite = true
	case false:
		isWhite = false
	}

	var moves []Square
	var targets []Square

	p := b[s]
	switch p {
	case WhiteKnight:
		moves = knightMoves(s, b)
	case BlackKnight:
		moves = knightMoves(s, b)
	default:
		moves = []Square{}
	}
	for _, val := range moves {
		if isWhite && b[val] < 0 {
			targets = append(targets, val)
		} else if !isWhite && b[val] > 0 {
			targets = append(targets, val)
		}
	}
	return targets
}

func rookTargets(s Square, b [64]Piece) []Square {
	var isWhite bool
	switch b[s] > 0 {
	case true:
		isWhite = true
	case false:
		isWhite = false
	}

	var moves []Square
	var targets []Square

	p := b[s]
	switch p {
	case WhiteRook:
		moves = rookMoves(s, b)
	case BlackRook:
		moves = rookMoves(s, b)
	default:
		moves = []Square{}
	}
	for _, val := range moves {
		if isWhite && b[val] < 0 {
			targets = append(targets, val)
		} else if !isWhite && b[val] > 0 {
			targets = append(targets, val)
		}
	}
	return targets
}

func queenTargets(s Square, b [64]Piece) []Square {
	var isWhite bool
	switch b[s] > 0 {
	case true:
		isWhite = true
	case false:
		isWhite = false
	}

	var moves []Square
	var targets []Square

	p := b[s]
	switch p {
	case WhiteQueen:
		moves = queenMoves(s, b)
	case BlackQueen:
		moves = queenMoves(s, b)
	default:
		moves = []Square{}
	}
	for _, val := range moves {
		if isWhite && b[val] < 0 {
			targets = append(targets, val)
		} else if !isWhite && b[val] > 0 {
			targets = append(targets, val)
		}
	}
	return targets
}

func kingTargets(s Square, b [64]Piece) []Square {
	var isWhite bool
	switch b[s] > 0 {
	case true:
		isWhite = true
	case false:
		isWhite = false
	}

	var moves []Square
	var targets []Square

	p := b[s]
	switch p {
	case WhiteKing:
		moves = whiteKingMoves(s, b)
	case BlackKing:
		moves = whiteKingMoves(s, b)
	default:
		moves = []Square{}
	}
	for _, val := range moves {
		if isWhite && b[val] < 0 {
			targets = append(targets, val)
		} else if !isWhite && b[val] > 0 {
			targets = append(targets, val)
		}
	}
	return targets
}
