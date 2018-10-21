package main

import (
	"fmt"
	"time"
)

func main() {
	spinner(100 * time.Millisecond)
}

func spinner(duration time.Duration) {
	for {
		for _, v := range `-\|/` {
			fmt.Printf("\r%c ", v)
			time.Sleep(duration)
		}
	}
}
