package chess

func (b *ChessBoard) verticalTop(s Square, sq []Square) []Square {
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
	sq = b.movementAlgorithm(startPos, startRow, startCol, movePos, moveRow, moveCol, isWhite, sq)

	return sq
}

func (b *ChessBoard) upperRightDiag(s Square, sq []Square) []Square {
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
	sq = b.movementAlgorithm(startPos, startRow, startCol, movePos, moveRow, moveCol, isWhite, sq)

	return sq
}

func (b *ChessBoard) horizontalRight(s Square, sq []Square) []Square {
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
	sq = b.movementAlgorithm(startPos, startRow, startCol, movePos, moveRow, moveCol, isWhite, sq)

	return sq
}

func (b *ChessBoard) lowerRightDiag(s Square, sq []Square) []Square {
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
	sq = b.movementAlgorithm(startPos, startRow, startCol, movePos, moveRow, moveCol, isWhite, sq)
	return sq
}

func (b *ChessBoard) verticalBottom(s Square, sq []Square) []Square {
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
	sq = b.movementAlgorithm(startPos, startRow, startCol, movePos, moveRow, moveCol, isWhite, sq)

	return sq
}

func (b *ChessBoard) lowerLeftDiag(s Square, sq []Square) []Square {
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
	sq = b.movementAlgorithm(startPos, startRow, startCol, movePos, moveRow, moveCol, isWhite, sq)

	return sq
}

func (b *ChessBoard) horizontalLeft(s Square, sq []Square) []Square {
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
	sq = b.movementAlgorithm(startPos, startRow, startCol, movePos, moveRow, moveCol, isWhite, sq)

	return sq
}

func (b *ChessBoard) upperLeftDiag(s Square, sq []Square) []Square {
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
	sq = b.movementAlgorithm(startPos, startRow, startCol, movePos, moveRow, moveCol, isWhite, sq)

	return sq
}

func (b *ChessBoard) knightMoves(s Square) []Square {
	var isWhite bool
	switch b.board[s] > 0 {
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
		if isWhite && b.board[target] < 0 {
			moves = append(moves, target)
		} else if !isWhite && b.board[target] > 0 {
			moves = append(moves, target)
		} else if b.board[target] == Empty {
			moves = append(moves, target)
		}
	}
	return moves
}

func (b *ChessBoard) kingMoves(s Square) []Square {
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
		if isWhite && b.board[target] < 0 {
			moves = append(moves, target)
		} else if !isWhite && b.board[target] > 0 {
			moves = append(moves, target)
		} else if b.board[target] == Empty {
			moves = append(moves, target)
		}
	}
	return moves
}
func (b *ChessBoard) movementAlgorithm(startPos Square, startRow Square, startCol Square, movePos Square, moveRow Square, moveCol Square, isWhite bool, sq []Square) []Square {
	isBlack := !isWhite
	for i, r, c := startPos, startRow, startCol; (i.row() == r && i.col() == c) && ((i <= h8) && (i >= a1)); i, r, c = i+movePos, i.row()+moveRow, i.col()+moveCol {
		if isWhite && b.board[i] < 0 {
			sq = append(sq, i)
			break
		} else if isBlack && b.board[i] > 0 {
			sq = append(sq, i)
			break
		} else if isWhite && b.board[i] > 0 {
			break
		} else if isBlack && b.board[i] < 0 {
			break
		} else if b.board[i] == Empty {
			sq = append(sq, i)
		}
	}
	return sq
}

func (b *ChessBoard) whitePawnMoves(s Square) []Square {
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

func (b *ChessBoard) blackPawnMoves(s Square) []Square {
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

func (b *ChessBoard) bishopMoves(s Square) []Square {
	var moves []Square

	moves = b.upperRightDiag(s, moves)
	moves = b.upperLeftDiag(s, moves)
	moves = b.lowerRightDiag(s, moves)
	moves = b.lowerLeftDiag(s, moves)

	return moves
}

func (b *ChessBoard) rookMoves(s Square) []Square {
	var moves []Square

	moves = b.horizontalLeft(s, moves)
	moves = b.horizontalRight(s, moves)
	moves = b.verticalTop(s, moves)
	moves = b.verticalBottom(s, moves)

	return moves
}

func (b *ChessBoard) queenMoves(s Square) []Square {
	var moves []Square

	moves = b.upperRightDiag(s, moves)
	moves = b.upperLeftDiag(s, moves)
	moves = b.lowerRightDiag(s, moves)
	moves = b.lowerLeftDiag(s, moves)
	moves = b.horizontalLeft(s, moves)
	moves = b.horizontalRight(s, moves)
	moves = b.verticalTop(s, moves)
	moves = b.verticalBottom(s, moves)

	return moves
}
