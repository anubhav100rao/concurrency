package select_statements

import (
	"fmt"
	"time"
)

func longTask(done chan<- bool) {
	// Simulate a long operation
	time.Sleep(2 * time.Second)
	done <- true
}

func RunTimeouts() {
	done := make(chan bool)
	go longTask(done)

	select {
	case <-done:
		fmt.Println("Task completed!")
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout: gave up waiting.")
	}
}
