package chess

// Purpose of blocks is to find squares that block a certain piece
// from attacking.
// This is used to be able to calculate check-mate, so that
// if a king can not move or kill a checking piece, one of the
// pieces under the player in check need to block the check in some way
// by standing in front of it

func blocks(s, t Square, b [64]Piece) []Square {
	p := b[s]
	var blocks []Square
	switch p {
	case WhitePawn:
		blocks = pawnBlocks(s, b)
	case BlackPawn:
		blocks = pawnBlocks(s, b)
	case WhiteKnight:
		blocks = knightBlocks(s, b)
	case BlackKnight:
		blocks = knightBlocks(s, b)
	case WhiteBishop:
		blocks = bishopBlocks(s, t, b)
	case BlackBishop:
		blocks = bishopBlocks(s, t, b)
	case WhiteRook:
		blocks = rookBlocks(s, t, b)
	case BlackRook:
		blocks = rookBlocks(s, t, b)
	case WhiteQueen:
		blocks = queenBlocks(s, t, b)
	case BlackQueen:
		blocks = queenBlocks(s, t, b)
	}
	return blocks
}

func pawnBlocks(s Square, b [64]Piece) []Square {

	var blocks []Square
	blocks = append(blocks, s)
	return blocks
}

func knightBlocks(s Square, b [64]Piece) []Square {

	var blocks []Square
	blocks = append(blocks, s)
	return blocks
}

func bishopBlocks(s, kingPos Square, b [64]Piece) []Square {

	var blocks []Square

	directions := []func(Square, []Square, [64]Piece) []Square{
		lowerLeftDiag,
		lowerRightDiag,
		upperLeftDiag,
		upperRightDiag,
	}

	for _, lambda := range directions {
		moves := lambda(s, []Square{}, b)
		blocks = []Square{}
		blocks = append(blocks, s)

		for _, square := range moves {
			if square == kingPos {
				return blocks
			}
			blocks = append(blocks, square)
		}
	}

	panic("no king square found, something wrong before this func")
}

func queenBlocks(s, kingPos Square, b [64]Piece) []Square {

	var blocks []Square

	directions := []func(Square, []Square, [64]Piece) []Square{
		lowerLeftDiag,
		lowerRightDiag,
		upperLeftDiag,
		upperRightDiag,
		verticalTop,
		horizontalRight,
		verticalBottom,
		horizontalLeft,
	}

	for _, lambda := range directions {
		moves := lambda(s, []Square{}, b)
		blocks = []Square{}
		blocks = append(blocks, s)

		for _, square := range moves {
			if square == kingPos {
				return blocks
			}
			blocks = append(blocks, square)
		}
	}

	panic("no king square found, something wrong before this func")
}

func rookBlocks(s, kingPos Square, b [64]Piece) []Square {

	var blocks []Square

	directions := []func(Square, []Square, [64]Piece) []Square{
		verticalTop,
		horizontalRight,
		verticalBottom,
		horizontalLeft,
	}

	for _, lambda := range directions {
		moves := lambda(s, []Square{}, b)
		blocks = []Square{}
		blocks = append(blocks, s)

		for _, square := range moves {
			if square == kingPos {
				return blocks
			}
			blocks = append(blocks, square)
		}
	}

	panic("no king square found, something wrong before this func")
}
