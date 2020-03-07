package chess

func targets(s Square, b [64]Piece) []Move {
	p := b[s]
	var targets []Move
	switch p {
	//case WhitePawn:
	//	targets = pawnTargets(s, b)
	//case BlackPawn:
	//	targets = pawnTargets(s, b)
	case WhiteBishop:
		targets = bishopTargets(s, b)
	case BlackBishop:
		targets = bishopTargets(s, b)
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

//
//func pawnTargets(s Square, b [64]Piece) []Square {
//	var isWhite bool
//	switch b[s] > 0 {
//	case true:
//		isWhite = true
//	case false:
//		isWhite = false
//	}
//
//	var moves []Square
//	var targets []Square
//
//	p := b[s]
//	switch p {
//	case WhitePawn:
//		moves = whitePawnMoves(s, b)
//	case BlackPawn:
//		moves = blackPawnMoves(s, b)
//	default:
//		moves = []Square{}
//	}
//	for _, val := range moves {
//		if isWhite && b[val] < 0 {
//			targets = append(targets, val)
//		} else if !isWhite && b[val] > 0 {
//			targets = append(targets, val)
//		}
//	}
//	return targets
//}

func bishopTargets(s Square, b [64]Piece) []Move {
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
	case WhiteBishop:
		moves = bishopMoves(s, b)
	case BlackBishop:
		moves = bishopMoves(s, b)
	default:
		moves = []Move{}
	}
	for _, move := range moves {
		if isWhite && b[move.toSquare] < 0 {
			targets = append(targets, makeMove(piece, move.fromSquare, move.toSquare, Capture))
		} else if !isWhite && b[move.toSquare] > 0 {
			targets = append(targets, makeMove(piece, move.fromSquare, move.toSquare, Capture))
		}
	}
	return targets
}

func knightTargets(s Square, b [64]Piece) []Move {
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
	case WhiteKnight:
		moves = knightMoves(s, b)
	case BlackKnight:
		moves = knightMoves(s, b)
	default:
		moves = []Move{}
	}
	for _, move := range moves {
		if isWhite && b[move.toSquare] < 0 {
			targets = append(targets, makeMove(piece, move.fromSquare, move.toSquare, Capture))
		} else if !isWhite && b[move.toSquare] > 0 {
			targets = append(targets, makeMove(piece, move.fromSquare, move.toSquare, Capture))

		}
	}
	return targets
}

func generalTargets(s Square, b [64]Piece) []Move {
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
	case WhiteBishop:
		moves = bishopMoves(s, b)
	case BlackBishop:
		moves = bishopMoves(s, b)
	case WhiteKnight:
		moves = knightMoves(s, b)
	case BlackKnight:
		moves = knightMoves(s, b)
	case WhiteRook:
		moves = rookMoves(s, b)
	case BlackRook:
		moves = rookMoves(s, b)
	case WhiteQueen:
		moves = queenMoves(s, b)
	case BlackQueen:
		moves = queenMoves(s, b)
	case WhiteKing:
		moves = kingMoves(s, b)
	case BlackKing:
		moves = kingMoves(s, b)
	default:
		moves = []Move{}
	}
	for _, move := range moves {
		if isWhite && b[move.toSquare] < 0 {
			targets = append(targets, makeMove(piece, move.fromSquare, move.toSquare, Capture))
		} else if !isWhite && b[move.toSquare] > 0 {
			targets = append(targets, makeMove(piece, move.fromSquare, move.toSquare, Capture))

		}
	}
	return targets
}
