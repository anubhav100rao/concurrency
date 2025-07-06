package select_statements

import (
	"fmt"
	"time"
)

func RunPeriodicWork() {
	ticker := time.NewTicker(500 * time.Millisecond)
	quite := make(chan struct{})

	go func() {
		time.Sleep(2 * time.Second)
		close(quite)
	}()

	for {
		select {
		case t := <-ticker.C:
			fmt.Println("Tick at", t.Format("15:04:05.000"))
		case <-quite:
			fmt.Println("Stopping ticker")
			ticker.Stop()
			return
		}
	}
}
