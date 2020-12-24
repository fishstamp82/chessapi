package chess

import "time"

func timeNow() int64 {
	return time.Now().UTC().UnixNano()
}
