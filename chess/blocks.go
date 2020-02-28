package chess

func (b *ChessBoard) pawnBlocks(s Square) []Square {

	var blocks []Square
	blocks = append(blocks, s)
	return blocks
}

func (b *ChessBoard) knightBlocks(s Square) []Square {

	var blocks []Square
	blocks = append(blocks, s)
	return blocks
}

func (b *ChessBoard) bishopBlocks(s, kingPos Square) []Square {

	var blocks []Square

	directions := []func(Square, []Square) []Square{
		b.lowerLeftDiag,
		b.lowerRightDiag,
		b.upperLeftDiag,
		b.upperRightDiag,
	}

	for _, lambda := range directions {
		moves := lambda(s, []Square{})
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

func (b *ChessBoard) queenBlocks(s, kingPos Square) []Square {

	var blocks []Square

	directions := []func(Square, []Square) []Square{
		b.lowerLeftDiag,
		b.lowerRightDiag,
		b.upperLeftDiag,
		b.upperRightDiag,
		b.verticalTop,
		b.horizontalRight,
		b.verticalBottom,
		b.horizontalLeft,
	}

	for _, lambda := range directions {
		moves := lambda(s, []Square{})
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

func (b *ChessBoard) rookBlocks(s, kingPos Square) []Square {

	var blocks []Square

	directions := []func(Square, []Square) []Square{
		b.verticalTop,
		b.horizontalRight,
		b.verticalBottom,
		b.horizontalLeft,
	}

	for _, lambda := range directions {
		moves := lambda(s, []Square{})
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
