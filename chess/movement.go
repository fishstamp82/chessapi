package chess

func (b *Board) horizontalLeft(s Square, sq []Square) []Square {
	var isWhite bool
	switch b.board[s] > 0 {
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
	for i, r, c := startPos, startRow, startCol; (i.row() == r && i.col() == c) && ((i <= h8) && (i >= a1)); i, r, c = i+movePos, i.row()+moveRow, i.col()+moveCol {

		if isWhite && b.board[i] < 0 {
			sq = append(sq, i)
			break
		} else if !isWhite && b.board[i] > 0 {
			sq = append(sq, i)
			break
		} else if b.board[i] == Empty {
			sq = append(sq, i)
		}
	}
	return sq
}

func (b *Board) horizontalRight(s Square, sq []Square) []Square {
	var isWhite bool
	switch b.board[s] > 0 {
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
	for i, r, c := startPos, startRow, startCol; (i.row() == r && i.col() == c) && ((i <= h8) && (i >= a1)); i, r, c = i+movePos, i.row()+moveRow, i.col()+moveCol {

		if isWhite && b.board[i] < 0 {
			sq = append(sq, i)
			break
		} else if !isWhite && b.board[i] > 0 {
			sq = append(sq, i)
			break
		} else if b.board[i] == Empty {
			sq = append(sq, i)
		}
	}
	return sq
}

func (b *Board) verticalTop(s Square, sq []Square) []Square {
	var isWhite bool
	switch b.board[s] > 0 {
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
	for i, r, c := startPos, startRow, startCol; (i.row() == r && i.col() == c) && ((i <= h8) && (i >= a1)); i, r, c = i+movePos, i.row()+moveRow, i.col()+moveCol {

		if isWhite && b.board[i] < 0 {
			sq = append(sq, i)
			break
		} else if !isWhite && b.board[i] > 0 {
			sq = append(sq, i)
			break
		} else if b.board[i] == Empty {
			sq = append(sq, i)
		}
	}
	return sq
}

func (b *Board) verticalBottom(s Square, sq []Square) []Square {
	var isWhite bool
	switch b.board[s] > 0 {
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
	for i, r, c := startPos, startRow, startCol; (i.row() == r && i.col() == c) && ((i <= h8) && (i >= a1)); i, r, c = i+movePos, i.row()+moveRow, i.col()+moveCol {

		if isWhite && b.board[i] < 0 {
			sq = append(sq, i)
			break
		} else if !isWhite && b.board[i] > 0 {
			sq = append(sq, i)
			break
		} else if b.board[i] == Empty {
			sq = append(sq, i)
		}
	}
	return sq
}

func (b *Board) upperLeftDiag(s Square, sq []Square) []Square {
	var isWhite bool
	switch b.board[s] > 0 {
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
	for i, r, c := startPos, startRow, startCol; (i.row() == r && i.col() == c) && ((i <= h8) && (i >= a1)); i, r, c = i+movePos, i.row()+moveRow, i.col()+moveCol {

		if isWhite && b.board[i] < 0 {
			sq = append(sq, i)
			break
		} else if !isWhite && b.board[i] > 0 {
			sq = append(sq, i)
			break
		} else if b.board[i] == Empty {
			sq = append(sq, i)
		}
	}
	return sq
}

func (b *Board) upperRightDiag(s Square, sq []Square) []Square {
	var isWhite bool
	switch b.board[s] > 0 {
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
	for i, r, c := startPos, startRow, startCol; (i.row() == r && i.col() == c) && ((i <= h8) && (i >= a1)); i, r, c = i+movePos, i.row()+moveRow, i.col()+moveCol {
		if isWhite && b.board[i] < 0 {
			sq = append(sq, i)
			break
		} else if !isWhite && b.board[i] > 0 {
			sq = append(sq, i)
			break
		} else if b.board[i] == Empty {
			sq = append(sq, i)
		}
	}
	return sq
}
func (b *Board) lowerRightDiag(s Square, sq []Square) []Square {
	var isWhite bool
	switch b.board[s] > 0 {
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
	for i, r, c := startPos, startRow, startCol; (i.row() == r && i.col() == c) && ((i <= h8) && (i >= a1)); i, r, c = i+movePos, i.row()+moveRow, i.col()+moveCol {
		if isWhite && b.board[i] < 0 {
			sq = append(sq, i)
			break
		} else if !isWhite && b.board[i] > 0 {
			sq = append(sq, i)
			break
		} else if b.board[i] == Empty {
			sq = append(sq, i)
		}
	}
	return sq
}

func (b *Board) lowerLeftDiag(s Square, sq []Square) []Square {
	var isWhite bool
	switch b.board[s] > 0 {
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
	for i, r, c := startPos, startRow, startCol; (i.row() == r && i.col() == c) && ((i <= h8) && (i >= a1)); i, r, c = i+movePos, i.row()+moveRow, i.col()+moveCol {
		if isWhite && b.board[i] < 0 {
			sq = append(sq, i)
			break
		} else if !isWhite && b.board[i] > 0 {
			sq = append(sq, i)
			break
		} else if b.board[i] == Empty {
			sq = append(sq, i)
		}
	}
	return sq
}

func (b *Board) whitePawnMoves(s Square) []Square {
	var moves []Square
	var t Square
	var first, second Square
	col := s.col()
	row := s.row()
	pos := row*8 + col
	if row == 1 {
		first = pos + 8   // one square move
		second = pos + 16 // two square move
		if b.board[first] == 0 {
			moves = append(moves, first)
		}
		if b.board[second] == 0 && b.board[first] == 0 {
			moves = append(moves, second)
		}
	} else {
		first = pos + 8 // one square move
		if b.board[first] == 0 {
			moves = append(moves, first)
		}
	}

	upperRight := func(s []Square) []Square {
		t = pos + 9 // attack upper right
		if b.board[t] < 0 {
			s = append(s, t)
		}
		return s
	}

	upperLeft := func(s []Square) []Square {
		t = pos + 7 // attack upper right
		if b.board[t] < 0 {
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

func (b *Board) blackPawnMoves(s Square) []Square {
	var moves []Square
	var t Square
	var first, second Square
	col := s.col()
	row := s.row()
	pos := row*8 + col
	if row == 6 {
		first = pos - 8   // one square move
		second = pos - 16 // two square move
		if b.board[first] == 0 {
			moves = append(moves, first)
		}
		if b.board[second] == 0 && b.board[first] == 0 {
			moves = append(moves, second)
		}
	} else {
		first = pos - 8 // one square move
		if b.board[first] == 0 {
			moves = append(moves, first)
		}
	}

	lowerRight := func(s []Square) []Square {
		t = pos - 7 // attack lower right
		if b.board[t] > 0 {
			s = append(s, t)
		}
		return s
	}

	lowerLeft := func(s []Square) []Square {
		t = pos - 9 // attack lower left
		if b.board[t] > 0 {
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

func (b *Board) bishopMoves(s Square) []Square {
	var moves []Square

	moves = b.upperRightDiag(s, moves)
	moves = b.upperLeftDiag(s, moves)
	moves = b.lowerRightDiag(s, moves)
	moves = b.lowerLeftDiag(s, moves)

	return moves
}

func (b *Board) rookMoves(s Square) []Square {
	var moves []Square

	moves = b.horizontalLeft(s, moves)
	moves = b.horizontalRight(s, moves)
	moves = b.verticalTop(s, moves)
	moves = b.verticalBottom(s, moves)

	return moves
}
