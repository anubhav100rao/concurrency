package channels

import (
	"fmt"
	"time"
)

func RunRateLimiter() {
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	requests := make(chan int, 10)
	go func() {
		for i := 1; i <= 10; i++ {
			requests <- i
		}
		close(requests)
	}()

	for req := range requests {
		fmt.Println(<-ticker.C)
		fmt.Println("Sending requests", req, "at", time.Now().Format("15:04:05.000"), "\n")
	}
}
