package chess

func (b *MailBoxBoard) squaresWithoutKing(p Player) []Square {
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
		piece = b.board[pos]
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

func (b *MailBoxBoard) kingSquare(p Player) Square {
	var king Piece
	switch p {
	case White:
		king = WhiteKing
	case Black:
		king = BlackKing
	}

	for pos := a1; pos <= h8; pos += 1 {
		if b.board[pos] == king {
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
