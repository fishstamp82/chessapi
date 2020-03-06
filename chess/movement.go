package chess

import (
	"errors"
	"fmt"
)

type piecePosition struct {
	piece    Piece
	position Square
}

type Move struct {
	fromSquare     Square
	toSquare       Square
	piecePositions []piecePosition // Resulting pieces in each square
	moveType       MovementType
}

func validMoves(s Square, b [64]Piece, ctx context) ([]Move, error) {
	p := b[s]
	var moves []Move
	var err error

	switch p {
	case WhitePawn:
		moves, err = pawnMoves(s, b)
	case BlackPawn:
		moves, err = pawnMoves(s, b)
		//case WhiteBishop:
		//	moves = bishopMoves(s, b)
		//case BlackBishop:
		//	moves = bishopMoves(s, b)
		//case WhiteKnight:
		//	moves = knightMoves(s, b)
		//case BlackKnight:
		//	moves = knightMoves(s, b)
		//case WhiteRook:
		//	moves = rookMoves(s, b)
		//case BlackRook:
		//	moves = rookMoves(s, b)
		//case WhiteQueen:
		//	moves = queenMoves(s, b)
		//case BlackQueen:
		//	moves = queenMoves(s, b)
		//case WhiteKing:
		//	moves = whiteKingMoves(s, b)
		//	moves = append(moves, whiteKingCastleMoves(s, b, ctx)...)
		//case BlackKing:
		//	moves = blackKingMoves(s, b)
		//	moves = append(moves, blackKingCastleMoves(s, b, ctx)...)
	}
	if err != nil {
		return nil, err
	}
	return moves, nil
}

func verticalTop(s Square, sq []Square, b [64]Piece) []Square {
	var isWhite bool
	switch b[s] > 0 {
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
	sq = movementAlgorithm(startPos, startRow, startCol, movePos, moveRow, moveCol, isWhite, sq, b)

	return sq
}

func upperRightDiag(s Square, sq []Square, b [64]Piece) []Square {
	var isWhite bool
	switch b[s] > 0 {
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
	sq = movementAlgorithm(startPos, startRow, startCol, movePos, moveRow, moveCol, isWhite, sq, b)

	return sq
}

func horizontalRight(s Square, sq []Square, b [64]Piece) []Square {
	var isWhite bool
	switch b[s] > 0 {
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
	sq = movementAlgorithm(startPos, startRow, startCol, movePos, moveRow, moveCol, isWhite, sq, b)

	return sq
}

func lowerRightDiag(s Square, sq []Square, b [64]Piece) []Square {
	var isWhite bool
	switch b[s] > 0 {
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
	sq = movementAlgorithm(startPos, startRow, startCol, movePos, moveRow, moveCol, isWhite, sq, b)
	return sq
}

func verticalBottom(s Square, sq []Square, b [64]Piece) []Square {
	var isWhite bool
	switch b[s] > 0 {
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
	sq = movementAlgorithm(startPos, startRow, startCol, movePos, moveRow, moveCol, isWhite, sq, b)

	return sq
}

func lowerLeftDiag(s Square, sq []Square, b [64]Piece) []Square {
	var isWhite bool
	switch b[s] > 0 {
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
	sq = movementAlgorithm(startPos, startRow, startCol, movePos, moveRow, moveCol, isWhite, sq, b)

	return sq
}

func horizontalLeft(s Square, sq []Square, b [64]Piece) []Square {
	var isWhite bool
	switch b[s] > 0 {
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
	sq = movementAlgorithm(startPos, startRow, startCol, movePos, moveRow, moveCol, isWhite, sq, b)

	return sq
}

func upperLeftDiag(s Square, sq []Square, b [64]Piece) []Square {
	var isWhite bool
	switch b[s] > 0 {
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
	sq = movementAlgorithm(startPos, startRow, startCol, movePos, moveRow, moveCol, isWhite, sq, b)

	return sq
}

func knightMoves(s Square, b [64]Piece) []Square {
	var isWhite bool
	switch b[s] > 0 {
	case true:
		isWhite = true
	case false:
		isWhite = false
	}
	_ = isWhite
	col := s.col()
	row := s.row()
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
	var moves []Square

	var target, r, c Square
	for _, val := range combos {
		target = val[0]
		r = val[1]
		c = val[2]

		if target.row() != r {
			continue
		}
		if target.col() != c {
			continue
		}
		if target < a1 || h8 < target {
			continue
		}
		if isWhite && b[target] < 0 {
			moves = append(moves, target)
		} else if !isWhite && b[target] > 0 {
			moves = append(moves, target)
		} else if b[target] == Empty {
			moves = append(moves, target)
		}
	}
	return moves
}

func whiteKingMoves(s Square, b [64]Piece) []Square {

	col := s.col()
	row := s.row()
	pos := row*8 + col

	topLeft := pos + 7
	top := pos + 8
	topRight := pos + 9
	right := pos + 1
	downRight := pos - 7
	down := pos - 8
	downLeft := pos - 9
	left := pos - 1

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
	var moves []Square

	var target, r, c Square
	for _, val := range combos {
		target = val[0]
		r = val[1]
		c = val[2]

		if target.row() != r {
			continue
		}
		if target.col() != c {
			continue
		}
		if target < a1 || h8 < target {
			continue
		}
		if b[target] < 0 {
			moves = append(moves, target)
		} else if b[target] == Empty {
			moves = append(moves, target)
		}
	}
	return moves
}

func blackKingMoves(s Square, b [64]Piece) []Square {
	col := s.col()
	row := s.row()
	pos := row*8 + col

	topLeft := pos + 7
	top := pos + 8
	topRight := pos + 9
	right := pos + 1
	downRight := pos - 7
	down := pos - 8
	downLeft := pos - 9
	left := pos - 1

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
	var moves []Square

	var target, r, c Square
	for _, val := range combos {
		target = val[0]
		r = val[1]
		c = val[2]

		if target.row() != r {
			continue
		}
		if target.col() != c {
			continue
		}
		if target < a1 || h8 < target {
			continue
		}

		if b[target] > 0 {
			moves = append(moves, target)
		} else if b[target] == Empty {
			moves = append(moves, target)
		}
	}

	return moves
}

//func whiteKingCastleMoves(s Square, b [64]Piece, ctx context) []Square {
//	var moves []Square
//
//	if s != e1 {
//		return moves
//	}
//
//	canCastleRight := ctx.whiteCanCastleRight
//	canCastleLeft := ctx.whiteCanCastleLeft
//	if canCastleRight {
//		if (b[f1] != Empty) || (b[g1] != Empty) {
//			canCastleRight = false
//		}
//		for _, p := range squaresWithoutKing(Black, b) {
//			for _, t := range validMoves(p, b, ctx) {
//				if t == f1 || t == g1 {
//					canCastleRight = false
//					break
//				}
//			}
//		}
//	}
//	if canCastleLeft {
//		for _, p := range squaresWithoutKing(Black, b) {
//			for _, t := range targets(p, b) {
//				if t == d1 || t == c1 || t == b1 {
//					canCastleLeft = false
//				}
//			}
//		}
//		if (b[b1] != Empty) || (b[c1] != Empty) || (b[d1] != Empty) {
//			canCastleLeft = false
//		}
//	}
//
//	if canCastleRight {
//		moves = append(moves, g1)
//	}
//	if canCastleLeft {
//		moves = append(moves, c1)
//	}
//	return moves
//}

func blackKingCastleMoves(s Square, b [64]Piece, ctx context) []Square {

	var moves []Square
	if s != e8 {
		return moves
	}
	canCastleRight := ctx.blackCanCastleRight
	canCastleLeft := ctx.blackCanCastleLeft
	if canCastleRight {
		for _, p := range squaresWithoutKing(White, b) {
			for _, t := range targets(p, b) {
				if t == f8 || t == g8 {
					canCastleRight = false
				}
			}
		}
		if (b[f8] != Empty) || (b[g8] != Empty) {
			canCastleRight = false
		}
	}
	if canCastleLeft {
		for _, p := range squaresWithoutKing(White, b) {
			for _, t := range targets(p, b) {
				if t == d8 || t == c8 || t == b8 {
					canCastleLeft = false
				}
			}
		}
		if (b[b8] != Empty) || (b[c8] != Empty) || (b[d8] != Empty) {
			canCastleLeft = false
		}
	}

	if canCastleRight {
		moves = append(moves, g8)
	}
	if canCastleLeft {
		moves = append(moves, c8)
	}
	return moves
}

func movementAlgorithm(startPos Square, startRow Square, startCol Square, movePos Square, moveRow Square, moveCol Square, isWhite bool, sq []Square, b [64]Piece) []Square {
	isBlack := !isWhite
	for i, r, c := startPos, startRow, startCol; (i.row() == r && i.col() == c) && ((i <= h8) && (i >= a1)); i, r, c = i+movePos, i.row()+moveRow, i.col()+moveCol {
		if isWhite && b[i] < 0 {
			sq = append(sq, i)
			break
		} else if isBlack && b[i] > 0 {
			sq = append(sq, i)
			break
		} else if isWhite && b[i] > 0 {
			break
		} else if isBlack && b[i] < 0 {
			break
		} else if b[i] == Empty {
			sq = append(sq, i)
		}
	}
	return sq
}

func whitePawnMoves(s Square, b [64]Piece) []Square {
	var moves []Square
	var t Square
	var first, second Square
	col := s.col()
	row := s.row()
	pos := row*8 + col
	if row == 1 {
		first = pos + 8   // one square move
		second = pos + 16 // two square move
		if b[first] == 0 {
			moves = append(moves, first)
		}
		if b[second] == 0 && b[first] == 0 {
			moves = append(moves, second)
		}
	} else {
		first = pos + 8 // one square move
		if b[first] == 0 {
			moves = append(moves, first)
		}
	}

	upperRight := func(s []Square) []Square {
		t = pos + 9 // attack upper right
		if b[t] < 0 {
			s = append(s, t)
		}
		return s
	}

	upperLeft := func(s []Square) []Square {
		t = pos + 7 // attack upper right
		if b[t] < 0 {
			s = append(s, t)
		}
		return s
	}
	if col == 0 {
		moves = upperRight(moves)
	} else if col == 7 {
		moves = upperLeft(moves)
	} else {
		moves = upperRight(moves)
		moves = upperLeft(moves)
	}
	return moves
}

func pawnMoves(fromSquare Square, b [64]Piece) ([]Move, error) {
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
	_ = player
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
		return nil, errors.New(fmt.Sprintf("illegal method on square: %s, piece: %s\n", fromSquare, b[fromSquare]))
	}

	col := fromSquare.col()
	rank := fromSquare.rank()
	pos := fromSquare.row()*8 + col

	var one, two, oneDiagonal Square
	if rank == startRank {
		one = pos + oneStep
		two = pos + twoStep
		if b[one] == Empty {
			moves = append(moves, makePawnMoves(pawn, fromSquare, one, Regular))
		}
		if b[one] == Empty && b[two] == Empty {
			moves = append(moves, makePawnMoves(pawn, fromSquare, two, Regular))
		}
	} else if (rank < finalRank) && (player == White) {
		one = pos + oneStep
		if b[one] == Empty {
			moves = append(moves, makePawnMoves(pawn, fromSquare, one, Regular))
		}
	} else if (rank > finalRank) && (player == Black) {
		one = pos + oneStep
		if b[one] == Empty {
			moves = append(moves, makePawnMoves(pawn, fromSquare, one, Regular))
		}
	} else if rank == finalRank {
		one = pos + oneStep
		if b[one] == Empty {
			moves = append(moves, makePawnPromotionMoves(player, fromSquare, one, Promotion)...)
		}
	} else {
		return nil, errors.New("pawn can't be on this rank")
	}

	//kills without promotion
	if (col == leftmostCol) && (rank < finalRank) && (player == White) {
		oneDiagonal = pos + diagonalRight
		if b[oneDiagonal] < 0 {
			moves = append(moves, makePawnMoves(pawn, fromSquare, oneDiagonal, Capture))
		}
	} else if (col == leftmostCol) && (rank > finalRank) && (player == Black) {
		if b[oneDiagonal] > 0 {
			moves = append(moves, makePawnMoves(pawn, fromSquare, oneDiagonal, Capture))
		}
	} else if (col == rightmostCol) && (rank < finalRank) && (player == White) {
		oneDiagonal = pos + diagonalLeft
		if b[oneDiagonal] < 0 {
			moves = append(moves, makePawnMoves(pawn, fromSquare, oneDiagonal, Capture))
		}
	} else if (col == rightmostCol) && (rank > finalRank) && (player == Black) {
		if (player == Black) && (b[oneDiagonal] > 0) {
			moves = append(moves, makePawnMoves(pawn, fromSquare, oneDiagonal, Capture))
		}
	} else if (rank < finalRank) && (player == White) {
		oneDiagonal = pos + diagonalRight
		if b[oneDiagonal] < 0 {
			moves = append(moves, makePawnMoves(pawn, fromSquare, oneDiagonal, Capture))
		}
		oneDiagonal = pos + diagonalLeft
		if b[oneDiagonal] < 0 {
			moves = append(moves, makePawnMoves(pawn, fromSquare, oneDiagonal, Capture))
		}
	} else if (rank > finalRank) && (player == Black) {
		oneDiagonal = pos + diagonalRight
		if b[oneDiagonal] > 0 {
			moves = append(moves, makePawnMoves(pawn, fromSquare, oneDiagonal, Capture))
		}
		oneDiagonal = pos + diagonalLeft
		if b[oneDiagonal] > 0 {
			moves = append(moves, makePawnMoves(pawn, fromSquare, oneDiagonal, Capture))
		}
	} else if (col == leftmostCol) && (rank == finalRank) {
		// killing plus promotion
		oneDiagonal = pos + diagonalRight
		if (player == White) && (b[oneDiagonal] < 0) {
			moves = append(moves, makePawnPromotionMoves(player, fromSquare, oneDiagonal, CapturePromotion)...)
		} else if (player == Black) && (b[oneDiagonal] > 0) {
			moves = append(moves, makePawnPromotionMoves(player, fromSquare, oneDiagonal, CapturePromotion)...)
		}
	} else if (col == rightmostCol) && (rank == finalRank) {
		oneDiagonal = pos + diagonalLeft
		if (player == White) && (b[oneDiagonal] < 0) {
			moves = append(moves, makePawnPromotionMoves(player, fromSquare, oneDiagonal, CapturePromotion)...)
		} else if (player == Black) && (b[oneDiagonal] > 0) {
			moves = append(moves, makePawnPromotionMoves(player, fromSquare, oneDiagonal, CapturePromotion)...)
		}
	} else if rank == finalRank {
		oneDiagonal = pos + diagonalRight
		if (player == White) && (b[oneDiagonal] < 0) {
			moves = append(moves, makePawnPromotionMoves(player, fromSquare, oneDiagonal, CapturePromotion)...)
		} else if (player == Black) && (b[oneDiagonal] > 0) {
			moves = append(moves, makePawnPromotionMoves(player, fromSquare, oneDiagonal, CapturePromotion)...)
		}
		oneDiagonal = pos + diagonalLeft
		if (player == White) && (b[oneDiagonal] < 0) {
			moves = append(moves, makePawnPromotionMoves(player, fromSquare, oneDiagonal, CapturePromotion)...)
		} else if (player == Black) && (b[oneDiagonal] > 0) {
			moves = append(moves, makePawnPromotionMoves(player, fromSquare, oneDiagonal, CapturePromotion)...)
		}
	}
	return moves, nil
}

func makePawnMoves(p Piece, f, t Square, mt MovementType) Move {
	return Move{
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
	}
}

func makeMoves(p Piece, f, t Square, mt MovementType) Move {
	return Move{
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
	}
}

func makePawnPromotionMoves(p Player, f, t Square, mt MovementType) []Move {
	var bishop, knight, rook, queen Piece
	switch p {
	case White:
		bishop = WhiteBishop
		knight = WhiteKnight
		rook = WhiteRook
		queen = WhiteQueen
	case Black:
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
			moveType: mt,
		})
	}
	return moves
}

func blackPawnMoves(s Square, b [64]Piece) []Square {
	var moves []Square
	var t Square
	var first, second Square
	col := s.col()
	row := s.row()
	pos := row*8 + col
	if row == 6 {
		first = pos - 8   // one square move
		second = pos - 16 // two square move
		if b[first] == 0 {
			moves = append(moves, first)
		}
		if b[second] == 0 && b[first] == 0 {
			moves = append(moves, second)
		}
	} else {
		first = pos - 8 // one square move
		if b[first] == 0 {
			moves = append(moves, first)
		}
	}

	lowerRight := func(s []Square) []Square {
		t = pos - 7 // attack lower right
		if b[t] > 0 {
			s = append(s, t)
		}
		return s
	}

	lowerLeft := func(s []Square) []Square {
		t = pos - 9 // attack lower left
		if b[t] > 0 {
			s = append(s, t)
		}
		return s
	}
	if col == 0 {
		moves = lowerRight(moves)
	} else if col == 7 {
		moves = lowerLeft(moves)
	} else {
		moves = lowerRight(moves)
		moves = lowerLeft(moves)
	}
	return moves
}

func bishopMoves(s Square, b [64]Piece) []Square {
	var moves []Square

	moves = upperRightDiag(s, moves, b)
	moves = upperLeftDiag(s, moves, b)
	moves = lowerRightDiag(s, moves, b)
	moves = lowerLeftDiag(s, moves, b)

	return moves
}

func rookMoves(s Square, b [64]Piece) []Square {
	var moves []Square

	moves = horizontalLeft(s, moves, b)
	moves = horizontalRight(s, moves, b)
	moves = verticalTop(s, moves, b)
	moves = verticalBottom(s, moves, b)

	return moves
}

func queenMoves(s Square, b [64]Piece) []Square {
	var moves []Square

	moves = upperRightDiag(s, moves, b)
	moves = upperLeftDiag(s, moves, b)
	moves = lowerRightDiag(s, moves, b)
	moves = lowerLeftDiag(s, moves, b)
	moves = horizontalLeft(s, moves, b)
	moves = horizontalRight(s, moves, b)
	moves = verticalTop(s, moves, b)
	moves = verticalBottom(s, moves, b)

	return moves
}
