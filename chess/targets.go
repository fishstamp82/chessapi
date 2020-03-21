package chess

func getTargets(s Square, b [64]Piece) []Move {
	p := b[s]
	var targets []Move
	switch p {
	case WhitePawn:
		targets = pawnTargets(s, b)
	case BlackPawn:
		targets = pawnTargets(s, b)
	case WhiteBishop:
		targets = generalTargets(s, b)
	case BlackBishop:
		targets = generalTargets(s, b)
	case WhiteKnight:
		targets = knightTargets(s, b)
	case BlackKnight:
		targets = knightTargets(s, b)
	case WhiteRook:
		targets = generalTargets(s, b)
	case BlackRook:
		targets = generalTargets(s, b)
	case WhiteQueen:
		targets = generalTargets(s, b)
	case BlackQueen:
		targets = generalTargets(s, b)
	case WhiteKing:
		targets = generalTargets(s, b)
	case BlackKing:
		targets = generalTargets(s, b)
	}
	return targets
}

func pawnTargets(s Square, b [64]Piece) []Move {
	var isWhite bool
	switch b[s] > 0 {
	case true:
		isWhite = true
	case false:
		isWhite = false
	}

	var moves []Move
	var targets []Move

	piece := b[s]
	switch piece {
	case WhitePawn:
		moves = pawnMoves(s, b, none)
	case BlackPawn:
		moves = pawnMoves(s, b, none)
	default:
		moves = []Move{}
	}
	for _, move := range moves {
		if isWhite && b[move.toSquare] < 0 {
			targets = append(targets, createPawnMove(piece, move.fromSquare, move.toSquare, []MovementType{Capture, PawnMove}))
		} else if !isWhite && b[move.toSquare] > 0 {
			targets = append(targets, createPawnMove(piece, move.fromSquare, move.toSquare, []MovementType{Capture, PawnMove}))
		}
	}
	return targets
}

func knightTargets(s Square, board [64]Piece) []Move {
	var isWhite bool
	switch board[s] > 0 {
	case true:
		isWhite = true
	case false:
		isWhite = false
	}

	var moves []Move
	var targets []Move

	piece := board[s]
	switch piece {
	case WhiteKnight:
		moves = knightMoves(s, board)
	case BlackKnight:
		moves = knightMoves(s, board)
	default:
		moves = []Move{}
	}
	for _, move := range moves {
		if isWhite && board[move.toSquare] < 0 {
			targets = append(targets, createMove(board, move.fromSquare, move.toSquare, []MovementType{Capture}))
		} else if !isWhite && board[move.toSquare] > 0 {
			targets = append(targets, createMove(board, move.fromSquare, move.toSquare, []MovementType{Capture}))

		}
	}
	return targets
}

func generalTargets(s Square, board [64]Piece) []Move {
	var isWhite bool
	switch board[s] > 0 {
	case true:
		isWhite = true
	case false:
		isWhite = false
	}

	var moves []Move
	var targets []Move

	piece := board[s]
	switch piece {
	case WhiteBishop:
		moves = bishopMoves(s, board)
	case BlackBishop:
		moves = bishopMoves(s, board)
	case WhiteKnight:
		moves = knightMoves(s, board)
	case BlackKnight:
		moves = knightMoves(s, board)
	case WhiteRook:
		moves = rookMoves(s, board)
	case BlackRook:
		moves = rookMoves(s, board)
	case WhiteQueen:
		moves = queenMoves(s, board)
	case BlackQueen:
		moves = queenMoves(s, board)
	case WhiteKing:
		moves = kingMoves(s, board)
	case BlackKing:
		moves = kingMoves(s, board)
	default:
		moves = []Move{}
	}
	for _, move := range moves {
		if isWhite && board[move.toSquare] < 0 {
			targets = append(targets, createMove(board, move.fromSquare, move.toSquare, []MovementType{Capture}))
		} else if !isWhite && board[move.toSquare] > 0 {
			targets = append(targets, createMove(board, move.fromSquare, move.toSquare, []MovementType{Capture}))

		}
	}
	return targets
}
