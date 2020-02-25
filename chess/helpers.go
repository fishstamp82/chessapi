package chess

func inSlice(t Square, list []Square) bool {
	for i := 0; i < len(list); i++ {
		if t == list[i] {
			return true
		}
	}
	return false
}
