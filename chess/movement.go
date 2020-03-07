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
	piece          Piece
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
		//moves = append(moves, whiteKingCastleMoves(s, b, ctx)...)
	case BlackKing:
		moves = kingMoves(s, b)
		//moves = append(moves, blackKingCastleMoves(s, b, ctx)...)
	}
	if err != nil {
		return nil, err
	}
	return moves, nil
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
			moves = append(moves, makeMove(piece, fromSquare, toSquare, Capture))
		} else if !isWhite && b[toSquare] > 0 {
			moves = append(moves, makeMove(piece, fromSquare, toSquare, Capture))
		} else if b[toSquare] == Empty {
			moves = append(moves, makeMove(piece, fromSquare, toSquare, Regular))
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
			moves = append(moves, makeMove(piece, fromSquare, toSquare, Capture))
		} else if (b[toSquare] < 0) && isBlack {
			moves = append(moves, makeMove(piece, fromSquare, toSquare, Capture))
		} else if b[toSquare] == Empty {
			moves = append(moves, makeMove(piece, fromSquare, toSquare, Regular))
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
				if t.toSquare == f8 || t.toSquare == g8 {
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
				if t.toSquare == d8 || t.toSquare == c8 || t.toSquare == b8 {
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

func movementAlgorithm(fromSquare, startPos Square, startRow Square, startCol Square, movePos Square, moveRow Square, moveCol Square, isWhite bool, p Piece, b [64]Piece) []Move {
	isBlack := !isWhite
	var moves []Move
	for i, r, c := startPos, startRow, startCol; (i.row() == r && i.col() == c) && ((i <= h8) && (i >= a1)); i, r, c = i+movePos, i.row()+moveRow, i.col()+moveCol {
		if isWhite && b[i] < 0 {
			moves = append(moves, makeMove(p, fromSquare, i, Capture))
			break
		} else if isBlack && b[i] > 0 {
			moves = append(moves, makeMove(p, fromSquare, i, Capture))
			break
		} else if isWhite && b[i] > 0 {
			break
		} else if isBlack && b[i] < 0 {
			break
		} else if b[i] == Empty {
			moves = append(moves, makeMove(p, fromSquare, i, Regular))
		}
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
	}
}

func makeMove(p Piece, f, t Square, mt MovementType) Move {
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
