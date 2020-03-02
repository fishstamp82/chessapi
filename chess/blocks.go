package chess

func (b *MailBoxBoard) blocks(s, t Square) []Square {
	p := b.board[s]
	var blocks []Square
	switch p {
	case WhitePawn:
		blocks = b.pawnBlocks(s)
	case BlackPawn:
		blocks = b.pawnBlocks(s)
	case WhiteBishop:
		blocks = b.bishopBlocks(s, t)
	case BlackBishop:
		blocks = b.bishopBlocks(s, t)
	case WhiteKnight:
		blocks = b.knightBlocks(s)
	case BlackKnight:
		blocks = b.knightBlocks(s)
	case WhiteRook:
		blocks = b.rookBlocks(s, t)
	case BlackRook:
		blocks = b.rookBlocks(s, t)
	case WhiteQueen:
		blocks = b.queenBlocks(s, t)
	case BlackQueen:
		blocks = b.queenBlocks(s, t)
	}
	return blocks
}

func (b *MailBoxBoard) pawnBlocks(s Square) []Square {

	var blocks []Square
	blocks = append(blocks, s)
	return blocks
}

func (b *MailBoxBoard) knightBlocks(s Square) []Square {

	var blocks []Square
	blocks = append(blocks, s)
	return blocks
}

func (b *MailBoxBoard) bishopBlocks(s, kingPos Square) []Square {

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

func (b *MailBoxBoard) queenBlocks(s, kingPos Square) []Square {

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

func (b *MailBoxBoard) rookBlocks(s, kingPos Square) []Square {

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
