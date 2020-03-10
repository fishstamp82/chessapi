package chess

import (
	"chessapi/chess"
	"testing"
)

func TestBoardMap(t *testing.T) {

	table := []struct {
		square  string
		unicode string
	}{
		{
			square:  "a1",
			unicode: "\u2656",
		},
	}

	for ind, row := range table {
		b := chess.NewMailBoxBoard()
		bMap := b.BoardMap()
		for key, _ := range bMap {
			if key == row.square {
				if bMap[row.square] != row.unicode {
					t.Errorf("expected: %s, got: %s for case %d\n", row.unicode, bMap[row.square], ind+1)
				}
			}

		}
	}
}
