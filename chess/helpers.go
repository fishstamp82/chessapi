package chess

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
