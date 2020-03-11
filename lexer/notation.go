package lexer

import (
	"fmt"
	"strings"
	"unicode"
)

// Given a chess board and chess notation, find the two squares involved in the move
// This will be matched in the API calling the ChessBoard.Move(string) method
//func readAlgebraicNotation(player chess.Player, notation string, board [64]chess.Piece, ctx chess.Context) (chess.Move, error) {
//var capture bool
//var piece, targetPiece, promotionPiece chess.Piece
//var fromSquare, toSquare chess.Square
//var got, isPromotion, isCheck, isCapture bool
//var validMoves []chess.Move
//var firstNotation, promotionStr string
//var notationSlice []string
//if len(notation) < 2 {
//	return chess.Move{}, errors.New("not valid notation")
//}
//
//whiteRuneToPiece := map[rune]chess.Piece{
//	0:   chess.WhitePawn,
//	'B': chess.WhiteBishop,
//	'N': chess.WhiteKnight,
//	'Q': chess.WhiteQueen,
//	'R': chess.WhiteRook,
//	'K': chess.WhiteKing,
//}
//blackRuneToPiece := map[rune]chess.Piece{
//	0:   chess.BlackPawn,
//	'B': chess.BlackBishop,
//	'N': chess.BlackKnight,
//	'Q': chess.BlackQueen,
//	'R': chess.BlackRook,
//	'K': chess.BlackKing,
//}
//var runeToPiece map[rune]chess.Piece
//switch player {
//case chess.White:
//	runeToPiece = whiteRuneToPiece
//case chess.Black:
//	runeToPiece = blackRuneToPiece
//}
//
//notation, isCheck = checkNotation(notation)
//notationSlice, isCapture = isCaptureNotation(notation)
//if isCapture {
//	firstNotation = notationSlice[0]
//	toSquare = chess.stringToSquare[notationSlice[1]]
//	notation = notationSlice[1] + notationSlice[1]
//} else {
//	targetPiece = chess.Empty
//	notation = notationSlice[0]
//}
//
//firstChar := readPieceMust(notation)
//if firstChar == 0 {
//	piece = runeToPiece[firstChar]
//	validMoves = validMovesForPiece(piece, board, ctx)
//
//	promotionStr, isPromotion = isPromotionNotation(notation)
//	if isPromotion {
//		promotionPiece = board[chess.stringToSquare[promotionStr]]
//		return chess.createPawnPromotionMove(fromSquare, toSquare, piece, targetPiece, promotionPiece, chess.Promotion), nil
//	}
//
//	var lastTwo []byte
//	for i := 2; i > 0; i-- {
//		lastTwo = append(lastTwo, notation[len(notation)-i])
//	}
//
//	if toSquare, got = chess.stringToSquare[string(lastTwo)]; got != true {
//		return chess.none, chess.none, errors.New(fmt.Sprintf("no such square: %notation\n", lastTwo))
//	}
//	if strings.Contains(notation, captureSymbol) {
//		strings.Split(notation, captureSymbol)
//	} else {
//		for _, move := range validMoves {
//			if move.toSquare == toSquare {
//				fromSquare = move.fromSquare
//			}
//		}
//	}
//	return chess.Move{}, nil
//} else {
//	piece = runeToPiece[rune(notation[0])]
//	//remainder := notation[1:]
//}
//
//return chess.Move{}, nil
//}

func isPromotionNotation(s string) (string, bool) {
	var promotionSymbol = "="
	if strings.Contains(s, promotionSymbol) {
		return strings.Split(s, promotionSymbol)[1], true
	}
	return "", false
}

func isCaptureNotation(s string) ([]string, bool) {
	var captureSymbol = "x"
	if strings.Contains(s, captureSymbol) {
		return strings.Split(s, captureSymbol), true
	}
	return []string{s}, false
}

func checkNotation(n string) (string, bool) {
	var checkSymbol = "!"
	if strings.HasSuffix(n, checkSymbol) {
		return strings.TrimSuffix(n, checkSymbol), true
	}
	return n, false
}

//func validMovesForPiece(piece chess.Piece, board [64]chess.Piece, ctx chess.Context) []chess.Move {
//	var moves []chess.Move
//	var squares []chess.Square
//	for i := chess.a1; i <= chess.h8; i++ {
//		if board[i] == piece {
//			squares = append(squares, i)
//		}
//	}
//	for _, sq := range squares {
//		moves = append(moves, chess.validMovesForSquare(sq, board, ctx)...)
//	}
//	return nil
//}

func readPieceMust(s string) rune {
	var firstChar = rune(s[0])
	switch fc := firstChar; {
	case unicode.ToLower(fc) == fc:
		return 0
	}
	switch firstChar {
	case 'B', 'N', 'K', 'Q', 'R':
		return firstChar
	}
	panic(fmt.Sprintf("no such rune %c\n", firstChar))
}
