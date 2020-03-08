package chess

type piecePosition struct {
	piece    Piece
	position Square
}

type Move struct {
	piece          Piece
	fromSquare     Square
	toSquare       Square
	piecePositions []piecePosition // Resulting pieces in each square
	moveType       MovementType
	reverseMove    *Move
}

func validMoves(fromSquare Square, board [64]Piece, ctx context) []Move {
	p := board[fromSquare]
	var moves []Move
	var player Player

	switch {
	case p > 0:
		player = White
	case p < 0:
		player = Black
	}

	switch p {
	case WhitePawn, BlackPawn:
		moves = pawnMoves(fromSquare, board, ctx.enPassantSquare)
	case WhiteBishop, BlackBishop:
		moves = bishopMoves(fromSquare, board)
	case WhiteKnight, BlackKnight:
		moves = knightMoves(fromSquare, board)
	case WhiteRook, BlackRook:
		moves = rookMoves(fromSquare, board)
	case WhiteQueen, BlackQueen:
		moves = queenMoves(fromSquare, board)
	case WhiteKing, BlackKing:
		moves = kingMoves(fromSquare, board)
		moves = append(moves, castleMoves(fromSquare, board, ctx)...)
	}
	moves = cleanMovesInCheck(moves, board, player)
	return moves
}

//remove moves that result in player being in check
func cleanMovesInCheck(m []Move, b [64]Piece, p Player) []Move {
	var cleanMoves []Move
	for _, move := range m {
		b = makeMove(move, b)
		ks := getKingSquare(p, b)
		if !inCheck(ks, b) {
			cleanMoves = append(cleanMoves, move)
		}
		b = makeMove(*move.reverseMove, b)
	}
	return cleanMoves
}

func verticalTop(s Square, b [64]Piece) []Move {
	var isWhite bool
	piece := b[s]

	switch piece > 0 {
	case true:
		isWhite = true
	case false:
		isWhite = false
	}
	col := s.col()
	row := s.row()
	pos := row*8 + col

	var movePos Square = 8
	var moveRow Square = 1
	var moveCol Square = 0

	startPos := pos + movePos
	startRow := row + moveRow
	startCol := col + moveCol
	return movementAlgorithm(s, startPos, startRow, startCol, movePos, moveRow, moveCol, isWhite, piece, b)

}

func upperRightDiag(s Square, b [64]Piece) []Move {
	var isWhite bool
	piece := b[s]
	switch piece > 0 {
	case true:
		isWhite = true
	case false:
		isWhite = false
	}
	col := s.col()
	row := s.row()
	pos := row*8 + col

	var movePos Square = 9
	var moveRow Square = 1
	var moveCol Square = 1

	startPos := pos + movePos
	startRow := row + moveRow
	startCol := col + moveCol
	return movementAlgorithm(s, startPos, startRow, startCol, movePos, moveRow, moveCol, isWhite, piece, b)

}

func horizontalRight(s Square, b [64]Piece) []Move {
	var isWhite bool
	piece := b[s]
	switch piece > 0 {
	case true:
		isWhite = true
	case false:
		isWhite = false
	}
	col := s.col()
	row := s.row()
	pos := row*8 + col

	var movePos Square = 1
	var moveRow Square = 0
	var moveCol Square = 1

	startPos := pos + movePos
	startRow := row + moveRow
	startCol := col + moveCol
	return movementAlgorithm(s, startPos, startRow, startCol, movePos, moveRow, moveCol, isWhite, piece, b)

}

func lowerRightDiag(s Square, b [64]Piece) []Move {
	var isWhite bool
	piece := b[s]
	switch piece > 0 {
	case true:
		isWhite = true
	case false:
		isWhite = false
	}
	col := s.col()
	row := s.row()
	pos := row*8 + col

	var movePos Square = -7
	var moveRow Square = -1
	var moveCol Square = 1

	startPos := pos + movePos
	startRow := row + moveRow
	startCol := col + moveCol
	return movementAlgorithm(s, startPos, startRow, startCol, movePos, moveRow, moveCol, isWhite, piece, b)

}

func verticalBottom(s Square, b [64]Piece) []Move {
	var isWhite bool
	piece := b[s]

	switch piece > 0 {
	case true:
		isWhite = true
	case false:
		isWhite = false
	}
	col := s.col()
	row := s.row()
	pos := row*8 + col

	var movePos Square = -8
	var moveRow Square = -1
	var moveCol Square = 0

	startPos := pos + movePos
	startRow := row + moveRow
	startCol := col + moveCol
	return movementAlgorithm(s, startPos, startRow, startCol, movePos, moveRow, moveCol, isWhite, piece, b)

}

func lowerLeftDiag(s Square, b [64]Piece) []Move {
	var isWhite bool
	piece := b[s]

	switch piece > 0 {
	case true:
		isWhite = true
	case false:
		isWhite = false
	}
	col := s.col()
	row := s.row()
	pos := row*8 + col

	var movePos Square = -9
	var moveRow Square = -1
	var moveCol Square = -1

	startPos := pos + movePos
	startRow := row + moveRow
	startCol := col + moveCol
	return movementAlgorithm(s, startPos, startRow, startCol, movePos, moveRow, moveCol, isWhite, piece, b)

}

func horizontalLeft(s Square, b [64]Piece) []Move {
	var isWhite bool
	piece := b[s]

	switch piece > 0 {
	case true:
		isWhite = true
	case false:
		isWhite = false
	}
	col := s.col()
	row := s.row()
	pos := row*8 + col

	var movePos Square = -1
	var moveRow Square = 0
	var moveCol Square = -1

	startPos := pos + movePos
	startRow := row + moveRow
	startCol := col + moveCol
	return movementAlgorithm(s, startPos, startRow, startCol, movePos, moveRow, moveCol, isWhite, piece, b)

}

func upperLeftDiag(s Square, b [64]Piece) []Move {
	var isWhite bool
	piece := b[s]

	switch piece > 0 {
	case true:
		isWhite = true
	case false:
		isWhite = false
	}
	col := s.col()
	row := s.row()
	pos := row*8 + col

	var movePos Square = 7
	var moveRow Square = 1
	var moveCol Square = -1

	startPos := pos + movePos
	startRow := row + moveRow
	startCol := col + moveCol
	return movementAlgorithm(s, startPos, startRow, startCol, movePos, moveRow, moveCol, isWhite, piece, b)

}

func knightMoves(fromSquare Square, b [64]Piece) []Move {
	var isWhite bool
	piece := b[fromSquare]
	switch piece > 0 {
	case true:
		isWhite = true
	case false:
		isWhite = false
	}
	_ = isWhite
	col := fromSquare.col()
	row := fromSquare.row()
	pos := row*8 + col

	topLeft := pos + 8 + 8 - 1
	topRight := pos + 8 + 8 + 1
	rightUp := pos + 1 + 1 + 8
	rightDown := pos + 1 + 1 - 8
	downRight := pos - 8 - 8 + 1
	downLeft := pos - 8 - 8 - 1
	leftDown := pos - 1 - 1 - 8
	leftUp := pos - 1 - 1 + 8

	topRow := row + 2
	topLeftCol := col - 1
	topRightCol := col + 1

	rightCol := col + 2
	rightUpRow := row + 1
	rightDownRow := row - 1

	downRow := row - 2
	downLeftCol := col - 1
	downRightCol := col + 1

	leftCol := col - 2
	leftUpRow := row + 1
	leftDownRow := row - 1

	combos := [8][3]Square{
		{topLeft, topRow, topLeftCol},
		{topRight, topRow, topRightCol},
		{rightUp, rightUpRow, rightCol},
		{rightDown, rightDownRow, rightCol},
		{downRight, downRow, downRightCol},
		{downLeft, downRow, downLeftCol},
		{leftDown, leftDownRow, leftCol},
		{leftUp, leftUpRow, leftCol},
	}
	var moves []Move

	var toSquare, r, c Square
	for _, val := range combos {
		toSquare = val[0]
		r = val[1]
		c = val[2]

		if toSquare.row() != r {
			continue
		}
		if toSquare.col() != c {
			continue
		}
		if toSquare < a1 || h8 < toSquare {
			continue
		}
		if isWhite && b[toSquare] < 0 {
			moves = append(moves, createMove(piece, fromSquare, toSquare, Capture))
		} else if !isWhite && b[toSquare] > 0 {
			moves = append(moves, createMove(piece, fromSquare, toSquare, Capture))
		} else if b[toSquare] == Empty {
			moves = append(moves, createMove(piece, fromSquare, toSquare, Regular))
		}
	}
	return moves
}

func kingMoves(fromSquare Square, b [64]Piece) []Move {

	piece := b[fromSquare]
	var isWhite, isBlack bool
	switch piece {
	case WhiteKing:
		isWhite = true
		isBlack = false
	case BlackKing:
		isWhite = false
		isBlack = true

	}
	col := fromSquare.col()
	row := fromSquare.row()
	pos := row*8 + col

	top := pos + 8
	topRight := pos + 9
	right := pos + 1
	downRight := pos - 7
	down := pos - 8
	downLeft := pos - 9
	left := pos - 1
	topLeft := pos + 7

	topRow := row + 1
	downRow := row - 1
	leftCol := col - 1
	rightCol := col + 1
	sameRow := row
	sameCol := col

	combos := [8][3]Square{
		{topLeft, topRow, leftCol},
		{top, topRow, sameCol},
		{topRight, topRow, rightCol},
		{right, sameRow, rightCol},
		{downRight, downRow, rightCol},
		{down, downRow, sameCol},
		{downLeft, downRow, leftCol},
		{left, sameRow, leftCol},
	}
	var moves []Move

	var toSquare, r, c Square
	for _, val := range combos {
		toSquare = val[0]
		r = val[1]
		c = val[2]

		if toSquare.row() != r {
			continue
		}
		if toSquare.col() != c {
			continue
		}
		if toSquare < a1 || h8 < toSquare {
			continue
		}
		if (b[toSquare] < 0) && isWhite {
			moves = append(moves, createMove(piece, fromSquare, toSquare, Capture))
		} else if (b[toSquare] < 0) && isBlack {
			moves = append(moves, createMove(piece, fromSquare, toSquare, Capture))
		} else if b[toSquare] == Empty {
			moves = append(moves, createMove(piece, fromSquare, toSquare, Regular))
		}
	}
	return moves
}

func castleMoves(kingSquare Square, b [64]Piece, ctx context) []Move {
	var moves []Move
	var piece = b[kingSquare]

	if kingSquare != e1 && kingSquare != e8 {
		return moves
	}

	var isWhite bool
	var isBlack bool
	var canCastleRight bool
	var canCastleLeft bool
	var shortCastleSquares []Square
	var longCastleSquares []Square
	var opponent Player

	switch piece {
	case WhiteKing:
		isWhite = true
		canCastleRight = ctx.whiteCanCastleRight
		canCastleLeft = ctx.whiteCanCastleLeft
		shortCastleSquares = []Square{f1, g1}
		longCastleSquares = []Square{d1, c1, b1}
		opponent = Black
	case BlackKing:
		isWhite = false
		canCastleRight = ctx.blackCanCastleRight
		canCastleLeft = ctx.blackCanCastleLeft
		shortCastleSquares = []Square{f8, g8}
		longCastleSquares = []Square{d8, c8, b8}
		opponent = White
	}
	isBlack = !isWhite

	if canCastleRight {
		for _, p := range squaresWithoutKing(opponent, b) {
			for _, move := range validMoves(p, b, ctx) {
				if inSquares(move.toSquare, append(shortCastleSquares, kingSquare)) {
					canCastleRight = false
					break
				}
			}
		}
		if !allEmpty(shortCastleSquares, b) {
			canCastleRight = false
		}
	}
	if canCastleLeft {
		for _, p := range squaresWithoutKing(opponent, b) {
			for _, move := range validMoves(p, b, ctx) {
				if inSquares(move.toSquare, append(longCastleSquares, kingSquare)) {
					canCastleLeft = false
					break
				}
			}
		}
		if !allEmpty(longCastleSquares, b) {
			canCastleLeft = false
		}
	}

	if canCastleRight && isWhite {
		moves = append(moves, createCastleMove(piece, e1, g1, ShortCastle))
	}
	if canCastleLeft && isWhite {
		moves = append(moves, createCastleMove(piece, e1, c1, LongCastle))
	}
	if canCastleRight && isBlack {
		moves = append(moves, createCastleMove(piece, e8, g8, ShortCastle))
	}
	if canCastleLeft && isBlack {
		moves = append(moves, createCastleMove(piece, e8, c8, LongCastle))
	}
	return moves
}

func allEmpty(squares []Square, board [64]Piece) bool {
	for _, sq := range squares {
		if board[sq] != Empty {
			return false
		}
	}
	return true
}

func movementAlgorithm(fromSquare, startPos Square, startRow Square, startCol Square, movePos Square, moveRow Square, moveCol Square, isWhite bool, p Piece, b [64]Piece) []Move {
	isBlack := !isWhite
	var moves []Move
	for i, r, c := startPos, startRow, startCol; (i.row() == r && i.col() == c) && ((i <= h8) && (i >= a1)); i, r, c = i+movePos, i.row()+moveRow, i.col()+moveCol {
		if isWhite && b[i] < 0 {
			moves = append(moves, createMove(p, fromSquare, i, Capture))
			break
		} else if isBlack && b[i] > 0 {
			moves = append(moves, createMove(p, fromSquare, i, Capture))
			break
		} else if isWhite && b[i] > 0 {
			break
		} else if isBlack && b[i] < 0 {
			break
		} else if b[i] == Empty {
			moves = append(moves, createMove(p, fromSquare, i, Regular))
		}
	}
	return moves
}

func pawnMoves(fromSquare Square, b [64]Piece, enPassant Square) []Move {
	var moves []Move

	var oneStepWhite Square = 8
	var twoStepWhite Square = 16

	var oneStepBlack Square = -8
	var twoStepBlack Square = -16

	var diagonalRightWhite Square = 9
	var diagonalLeftWhite Square = 7

	var diagonalRightBlack Square = -7
	var diagonalLeftBlack Square = -9

	var whiteStartRank, blackFinalRank Square = 2, 2
	var blackStartRank, whiteFinalRank Square = 7, 7

	var leftmostCol Square = 0
	var rightmostCol Square = 7

	// Used dependent on if its white or black pawns
	var pawn Piece
	var player Player
	var oneStep Square
	var twoStep Square
	var diagonalLeft Square
	var diagonalRight Square
	var startRank, finalRank Square
	switch b[fromSquare] {
	case WhitePawn:
		pawn = WhitePawn
		player = White
		oneStep = oneStepWhite
		twoStep = twoStepWhite
		startRank = whiteStartRank
		finalRank = whiteFinalRank
		diagonalLeft = diagonalLeftWhite
		diagonalRight = diagonalRightWhite
	case BlackPawn:
		pawn = BlackPawn
		player = Black
		oneStep = oneStepBlack
		twoStep = twoStepBlack
		startRank = blackStartRank
		finalRank = blackFinalRank
		diagonalLeft = diagonalLeftBlack
		diagonalRight = diagonalRightBlack
	default:
		panic("pawnMoves called without pawn square")
	}

	col := fromSquare.col()
	rank := fromSquare.rank()
	pos := fromSquare.row()*8 + col

	var one, two, oneDiagonal Square
	if rank == startRank {
		one = pos + oneStep
		two = pos + twoStep
		if b[one] == Empty {
			moves = append(moves, createPawnMove(pawn, fromSquare, one, Regular))
		}
		if b[one] == Empty && b[two] == Empty {
			moves = append(moves, createPawnMove(pawn, fromSquare, two, Regular))
		}
	} else if (rank < finalRank) && (player == White) {
		one = pos + oneStep
		if b[one] == Empty {
			moves = append(moves, createPawnMove(pawn, fromSquare, one, Regular))
		}
	} else if (rank > finalRank) && (player == Black) {
		one = pos + oneStep
		if b[one] == Empty {
			moves = append(moves, createPawnMove(pawn, fromSquare, one, Regular))
		}
	} else if rank == finalRank {
		one = pos + oneStep
		if b[one] == Empty {
			moves = append(moves, createPawnPromotionMoves(player, fromSquare, one, Empty, Promotion)...)
		}
	} else {
		panic("pawn can't be on this rank")
	}

	//kills without promotion
	if (col == leftmostCol) && (rank < finalRank) && (player == White) {
		oneDiagonal = pos + diagonalRight
		if b[oneDiagonal] < 0 {
			moves = append(moves, createPawnMove(pawn, fromSquare, oneDiagonal, Capture))
		} else if oneDiagonal == enPassant {
			moves = append(moves, createPawnEnPassantMove(pawn, fromSquare, oneDiagonal, CaptureEnPassant))
		}
	} else if (col == leftmostCol) && (rank > finalRank) && (player == Black) {
		oneDiagonal = pos + diagonalRight
		if b[oneDiagonal] > 0 {
			moves = append(moves, createPawnMove(pawn, fromSquare, oneDiagonal, Capture))
		} else if oneDiagonal == enPassant {
			moves = append(moves, createPawnEnPassantMove(pawn, fromSquare, oneDiagonal, CaptureEnPassant))
		}
	} else if (col == rightmostCol) && (rank < finalRank) && (player == White) {
		oneDiagonal = pos + diagonalLeft
		if b[oneDiagonal] < 0 {
			moves = append(moves, createPawnMove(pawn, fromSquare, oneDiagonal, Capture))
		} else if oneDiagonal == enPassant {
			moves = append(moves, createPawnEnPassantMove(pawn, fromSquare, oneDiagonal, CaptureEnPassant))
		}
	} else if (col == rightmostCol) && (rank > finalRank) && (player == Black) {
		oneDiagonal = pos + diagonalLeft
		if (player == Black) && (b[oneDiagonal] > 0) {
			moves = append(moves, createPawnMove(pawn, fromSquare, oneDiagonal, Capture))
		} else if oneDiagonal == enPassant {
			moves = append(moves, createPawnEnPassantMove(pawn, fromSquare, oneDiagonal, CaptureEnPassant))
		}
	} else if (rank < finalRank) && (player == White) {
		oneDiagonal = pos + diagonalRight
		if b[oneDiagonal] < 0 {
			moves = append(moves, createPawnMove(pawn, fromSquare, oneDiagonal, Capture))
		} else if oneDiagonal == enPassant {
			moves = append(moves, createPawnEnPassantMove(pawn, fromSquare, oneDiagonal, CaptureEnPassant))
		}
		oneDiagonal = pos + diagonalLeft
		if b[oneDiagonal] < 0 {
			moves = append(moves, createPawnMove(pawn, fromSquare, oneDiagonal, Capture))
		} else if oneDiagonal == enPassant {
			moves = append(moves, createPawnEnPassantMove(pawn, fromSquare, oneDiagonal, CaptureEnPassant))
		}
	} else if (rank > finalRank) && (player == Black) {
		oneDiagonal = pos + diagonalRight
		if b[oneDiagonal] > 0 {
			moves = append(moves, createPawnMove(pawn, fromSquare, oneDiagonal, Capture))
		} else if oneDiagonal == enPassant {
			moves = append(moves, createPawnEnPassantMove(pawn, fromSquare, oneDiagonal, CaptureEnPassant))
		}
		oneDiagonal = pos + diagonalLeft
		if b[oneDiagonal] > 0 {
			moves = append(moves, createPawnMove(pawn, fromSquare, oneDiagonal, Capture))
		} else if oneDiagonal == enPassant {
			moves = append(moves, createPawnEnPassantMove(pawn, fromSquare, oneDiagonal, CaptureEnPassant))
		}
	} else if (col == leftmostCol) && (rank == finalRank) {
		// killing plus promotion
		oneDiagonal = pos + diagonalRight
		if (player == White) && (b[oneDiagonal] < 0) {
			moves = append(moves, createPawnPromotionMoves(player, fromSquare, oneDiagonal, b[oneDiagonal], CapturePromotion)...)
		} else if (player == Black) && (b[oneDiagonal] > 0) {
			moves = append(moves, createPawnPromotionMoves(player, fromSquare, oneDiagonal, b[oneDiagonal], CapturePromotion)...)
		}
	} else if (col == rightmostCol) && (rank == finalRank) {
		oneDiagonal = pos + diagonalLeft
		if (player == White) && (b[oneDiagonal] < 0) {
			moves = append(moves, createPawnPromotionMoves(player, fromSquare, oneDiagonal, b[oneDiagonal], CapturePromotion)...)
		} else if (player == Black) && (b[oneDiagonal] > 0) {
			moves = append(moves, createPawnPromotionMoves(player, fromSquare, oneDiagonal, b[oneDiagonal], CapturePromotion)...)
		}
	} else if rank == finalRank {
		oneDiagonal = pos + diagonalRight
		if (player == White) && (b[oneDiagonal] < 0) {
			moves = append(moves, createPawnPromotionMoves(player, fromSquare, oneDiagonal, b[oneDiagonal], CapturePromotion)...)
		} else if (player == Black) && (b[oneDiagonal] > 0) {
			moves = append(moves, createPawnPromotionMoves(player, fromSquare, oneDiagonal, b[oneDiagonal], CapturePromotion)...)
		}
		oneDiagonal = pos + diagonalLeft
		if (player == White) && (b[oneDiagonal] < 0) {
			moves = append(moves, createPawnPromotionMoves(player, fromSquare, oneDiagonal, b[oneDiagonal], CapturePromotion)...)
		} else if (player == Black) && (b[oneDiagonal] > 0) {
			moves = append(moves, createPawnPromotionMoves(player, fromSquare, oneDiagonal, b[oneDiagonal], CapturePromotion)...)
		}
	}
	return moves
}

func createPawnMove(p Piece, f, t Square, mt MovementType) Move {
	return Move{
		piece:      p,
		fromSquare: f,
		toSquare:   t,
		piecePositions: []piecePosition{
			{
				piece:    p,
				position: t,
			},
			{
				piece:    Empty,
				position: f,
			},
		},
		moveType: mt,
		reverseMove: &Move{
			piecePositions: []piecePosition{
				{
					piece:    p,
					position: f,
				},
				{
					piece:    Empty,
					position: t,
				},
			},
		},
	}
}

func createPawnEnPassantMove(p Piece, f, t Square, mt MovementType) Move {
	var killPosition Square
	var killPawn Piece
	switch p {
	case WhitePawn:
		killPosition = t - 8
		killPawn = BlackPawn
	case BlackPawn:
		killPosition = t + 8
		killPawn = WhitePawn
	}
	m := Move{
		piece:      p,
		fromSquare: f,
		toSquare:   t,
		piecePositions: []piecePosition{
			{
				piece:    p,
				position: t,
			},
			{
				piece:    Empty,
				position: f,
			},
			{
				piece:    Empty,
				position: killPosition,
			},
		},
		moveType: mt,
		reverseMove: &Move{
			piecePositions: []piecePosition{
				{
					piece:    p,
					position: f,
				},
				{
					piece:    Empty,
					position: t,
				},
				{
					piece:    killPawn,
					position: killPosition,
				},
			},
		},
	}
	return m
}

func createMove(p Piece, f, t Square, mt MovementType) Move {
	return Move{
		piece:      p,
		fromSquare: f,
		toSquare:   t,
		piecePositions: []piecePosition{
			{
				piece:    p,
				position: t,
			},
			{
				piece:    Empty,
				position: f,
			},
		},
		moveType: mt,
		reverseMove: &Move{
			piecePositions: []piecePosition{
				{
					piece:    p,
					position: f,
				},
				{
					piece:    Empty,
					position: t,
				},
			},
		},
	}
}

func createCastleMove(p Piece, f, t Square, mt MovementType) Move {
	move := Move{
		piece:      p,
		fromSquare: f,
		toSquare:   t,
		piecePositions: []piecePosition{
			{
				piece:    p,
				position: t,
			},
			{
				piece:    Empty,
				position: f,
			},
		},
		reverseMove: &Move{
			piecePositions: []piecePosition{
				{
					piece:    p,
					position: f,
				},
				{
					piece:    Empty,
					position: t,
				},
			},
		},
		moveType: mt,
	}

	var isWhite, isBlack, short, long bool
	var rook Piece
	switch p {
	case WhiteKing:
		isWhite = true
		rook = WhiteRook
	case BlackKing:
		isWhite = false
		rook = BlackRook
	}
	switch mt {
	case ShortCastle:
		short = true
	case LongCastle:
		short = false
	}
	isBlack = !isWhite
	long = !short
	if isWhite && short {
		move.piecePositions = append(move.piecePositions,
			piecePosition{
				piece:    rook,
				position: f1,
			},
			piecePosition{
				piece:    Empty,
				position: h1,
			},
		)
		move.reverseMove.piecePositions = append(move.reverseMove.piecePositions,
			piecePosition{
				piece:    rook,
				position: h1,
			},
			piecePosition{
				piece:    Empty,
				position: g1,
			},
		)
	} else if isWhite && long {
		move.piecePositions = append(move.piecePositions,
			piecePosition{
				piece:    rook,
				position: d1,
			}, piecePosition{
				piece:    Empty,
				position: a1,
			},
		)
		move.reverseMove.piecePositions = append(move.reverseMove.piecePositions,
			piecePosition{
				piece:    rook,
				position: a1,
			},
			piecePosition{
				piece:    Empty,
				position: d1,
			},
		)
	} else if isBlack && short {
		move.piecePositions = append(move.piecePositions,
			piecePosition{
				piece:    rook,
				position: f8,
			}, piecePosition{
				piece:    Empty,
				position: h8,
			},
		)
		move.reverseMove.piecePositions = append(move.reverseMove.piecePositions,
			piecePosition{
				piece:    rook,
				position: h8,
			},
			piecePosition{
				piece:    Empty,
				position: f8,
			},
		)
	} else if isBlack && long {
		move.piecePositions = append(move.piecePositions,
			piecePosition{
				piece:    rook,
				position: d8,
			}, piecePosition{
				piece:    Empty,
				position: a8,
			},
		)
		move.reverseMove.piecePositions = append(move.reverseMove.piecePositions,
			piecePosition{
				piece:    rook,
				position: a8,
			},
			piecePosition{
				piece:    Empty,
				position: d8,
			},
		)
	}
	return move
}

func createPawnPromotionMoves(p Player, f, t Square, targetPiece Piece, mt MovementType) []Move {
	var pawn, bishop, knight, rook, queen, target Piece
	target = targetPiece
	switch p {
	case White:
		bishop = WhiteBishop
		knight = WhiteKnight
		rook = WhiteRook
		queen = WhiteQueen
		pawn = WhitePawn
	case Black:
		pawn = BlackPawn
		bishop = BlackBishop
		knight = BlackKnight
		rook = BlackRook
		queen = BlackQueen
	}
	var moves []Move
	for _, piece := range []Piece{bishop, knight, rook, queen} {
		moves = append(moves, Move{
			fromSquare: f,
			toSquare:   t,
			piecePositions: []piecePosition{
				{
					piece:    piece,
					position: t,
				},
				{
					piece:    Empty,
					position: f,
				},
			},
			reverseMove: &Move{
				piecePositions: []piecePosition{
					{
						piece:    pawn,
						position: f,
					},
					{
						piece:    target,
						position: t,
					},
				},
			},
			moveType: mt,
		})
	}
	return moves
}

func bishopMoves(s Square, b [64]Piece) []Move {
	var moves []Move

	moves = append(moves, upperRightDiag(s, b)...)
	moves = append(moves, upperLeftDiag(s, b)...)
	moves = append(moves, lowerRightDiag(s, b)...)
	moves = append(moves, lowerLeftDiag(s, b)...)

	return moves
}

func rookMoves(s Square, b [64]Piece) []Move {
	var moves []Move

	moves = append(moves, horizontalLeft(s, b)...)
	moves = append(moves, horizontalRight(s, b)...)
	moves = append(moves, verticalTop(s, b)...)
	moves = append(moves, verticalBottom(s, b)...)

	return moves
}

func queenMoves(s Square, b [64]Piece) []Move {
	var moves []Move

	moves = append(moves, upperRightDiag(s, b)...)
	moves = append(moves, upperLeftDiag(s, b)...)
	moves = append(moves, lowerRightDiag(s, b)...)
	moves = append(moves, lowerLeftDiag(s, b)...)
	moves = append(moves, horizontalLeft(s, b)...)
	moves = append(moves, horizontalRight(s, b)...)
	moves = append(moves, verticalTop(s, b)...)
	moves = append(moves, verticalBottom(s, b)...)

	return moves
}
