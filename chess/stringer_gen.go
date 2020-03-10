// Code generated by "stringer -type=Piece,Player,Square,State,MovementType -output=stringer_gen.go"; DO NOT EDIT.

package chess

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[BlackKing - -6]
	_ = x[BlackQueen - -5]
	_ = x[BlackRook - -4]
	_ = x[BlackBishop - -3]
	_ = x[BlackKnight - -2]
	_ = x[BlackPawn - -1]
	_ = x[Empty-0]
	_ = x[WhitePawn-1]
	_ = x[WhiteKnight-2]
	_ = x[WhiteBishop-3]
	_ = x[WhiteRook-4]
	_ = x[WhiteQueen-5]
	_ = x[WhiteKing-6]
}

const _Piece_name = "BlackKingBlackQueenBlackRookBlackBishopBlackKnightBlackPawnEmptyWhitePawnWhiteKnightWhiteBishopWhiteRookWhiteQueenWhiteKing"

var _Piece_index = [...]uint8{0, 9, 19, 28, 39, 50, 59, 64, 73, 84, 95, 104, 114, 123}

func (i Piece) String() string {
	i -= -6
	if i < 0 || i >= Piece(len(_Piece_index)-1) {
		return "Piece(" + strconv.FormatInt(int64(i+-6), 10) + ")"
	}
	return _Piece_name[_Piece_index[i]:_Piece_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Noone-0]
	_ = x[White-1]
	_ = x[Black-2]
	_ = x[Both-3]
}

const _Player_name = "NooneWhiteBlackBoth"

var _Player_index = [...]uint8{0, 5, 10, 15, 19}

func (i Player) String() string {
	if i >= Player(len(_Player_index)-1) {
		return "Player(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Player_name[_Player_index[i]:_Player_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[none - -1]
	_ = x[a1-0]
	_ = x[b1-1]
	_ = x[c1-2]
	_ = x[d1-3]
	_ = x[e1-4]
	_ = x[f1-5]
	_ = x[g1-6]
	_ = x[h1-7]
	_ = x[a2-8]
	_ = x[b2-9]
	_ = x[c2-10]
	_ = x[d2-11]
	_ = x[e2-12]
	_ = x[f2-13]
	_ = x[g2-14]
	_ = x[h2-15]
	_ = x[a3-16]
	_ = x[b3-17]
	_ = x[c3-18]
	_ = x[d3-19]
	_ = x[e3-20]
	_ = x[f3-21]
	_ = x[g3-22]
	_ = x[h3-23]
	_ = x[a4-24]
	_ = x[b4-25]
	_ = x[c4-26]
	_ = x[d4-27]
	_ = x[e4-28]
	_ = x[f4-29]
	_ = x[g4-30]
	_ = x[h4-31]
	_ = x[a5-32]
	_ = x[b5-33]
	_ = x[c5-34]
	_ = x[d5-35]
	_ = x[e5-36]
	_ = x[f5-37]
	_ = x[g5-38]
	_ = x[h5-39]
	_ = x[a6-40]
	_ = x[b6-41]
	_ = x[c6-42]
	_ = x[d6-43]
	_ = x[e6-44]
	_ = x[f6-45]
	_ = x[g6-46]
	_ = x[h6-47]
	_ = x[a7-48]
	_ = x[b7-49]
	_ = x[c7-50]
	_ = x[d7-51]
	_ = x[e7-52]
	_ = x[f7-53]
	_ = x[g7-54]
	_ = x[h7-55]
	_ = x[a8-56]
	_ = x[b8-57]
	_ = x[c8-58]
	_ = x[d8-59]
	_ = x[e8-60]
	_ = x[f8-61]
	_ = x[g8-62]
	_ = x[h8-63]
}

const _Square_name = "nonea1b1c1d1e1f1g1h1a2b2c2d2e2f2g2h2a3b3c3d3e3f3g3h3a4b4c4d4e4f4g4h4a5b5c5d5e5f5g5h5a6b6c6d6e6f6g6h6a7b7c7d7e7f7g7h7a8b8c8d8e8f8g8h8"

var _Square_index = [...]uint8{0, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112, 114, 116, 118, 120, 122, 124, 126, 128, 130, 132}

func (i Square) String() string {
	i -= -1
	if i < 0 || i >= Square(len(_Square_index)-1) {
		return "Square(" + strconv.FormatInt(int64(i+-1), 10) + ")"
	}
	return _Square_name[_Square_index[i]:_Square_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Playing-1]
	_ = x[Check-2]
	_ = x[CheckMate-3]
	_ = x[Draw-4]
	_ = x[Promo-5]
}

const _State_name = "PlayingCheckCheckMateDrawPromo"

var _State_index = [...]uint8{0, 7, 12, 21, 25, 30}

func (i State) String() string {
	i -= 1
	if i >= State(len(_State_index)-1) {
		return "State(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _State_name[_State_index[i]:_State_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Regular-0]
	_ = x[Capture-1]
	_ = x[ShortCastle-2]
	_ = x[LongCastle-3]
	_ = x[Promotion-4]
	_ = x[CapturePromotion-5]
	_ = x[CaptureEnPassant-6]
	_ = x[CheckMove-7]
}

const _MovementType_name = "RegularCaptureShortCastleLongCastlePromotionCapturePromotionCaptureEnPassantCheckMove"

var _MovementType_index = [...]uint8{0, 7, 14, 25, 35, 44, 60, 76, 85}

func (i MovementType) String() string {
	if i >= MovementType(len(_MovementType_index)-1) {
		return "MovementType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MovementType_name[_MovementType_index[i]:_MovementType_index[i+1]]
}
