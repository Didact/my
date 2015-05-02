package my

import (
	"time"
)

func Wait(b *bool) {
	for !(*b) {
		time.Sleep(10 * time.Millisecond)
	}
}
