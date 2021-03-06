package chess

// row goes from 0 to 7
func (s Square) row() Square {
	return s / 8
}

// col goes from 0 to 7
func (s Square) col() Square {
	return s % 8
}

// rank goes from 1 to 8
func (s Square) rank() Square {
	return s/8 + 1
}

func getKingSquareMust(p Color, b [64]Piece) Square {
	var king Piece
	switch p {
	case White:
		king = WhiteKing
	case Black:
		king = BlackKing
	}

	for pos := a1; pos <= h8; pos += 1 {
		if b[pos] == king {
			return pos
		}
	}

	panic("must have a white King")
}

func inSquares(t Square, list []Square) bool {
	for i := 0; i < len(list); i++ {
		if t == list[i] {
			return true
		}
	}
	return false
}

func uniqueSquares(s []Square) []Square {
	var uniq []Square
	for i := 0; i < len(s); i++ {
		if !inSquares(s[i], uniq) {
			uniq = append(uniq, s[i])
		}
	}
	return uniq
}

func validPromotion(piece Piece, player Color) bool {
	switch player {
	case White:
		switch piece {
		case WhiteKnight, WhiteBishop, WhiteRook, WhiteQueen:
			return true
		default:
			return false
		}
	case Black:
		switch piece {
		case BlackKnight, BlackBishop, BlackRook, BlackQueen:
			return true
		default:
			return false
		}
	}

	return false
}

func squaresWithoutKing(p Color, b [64]Piece) []Square {
	var isWhite bool
	var piece Piece
	switch p {
	case White:
		isWhite = true
	case Black:
		isWhite = false
	}

	var pieces []Square
	for pos := a1; pos <= h8; pos += 1 {
		piece = b[pos]
		if piece == WhiteKing || piece == BlackKing {
			continue
		}
		if piece > 0 && isWhite {
			pieces = append(pieces, pos)
		} else if piece < 0 && !isWhite {
			pieces = append(pieces, pos)

		}

	}
	return pieces
}
