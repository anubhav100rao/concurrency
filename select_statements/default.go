package select_statements

import "fmt"

func RunDefault() {
	ch := make(chan int, 1)
	ch <- 42

	select {
	case val := <-ch:
		fmt.Println("Got:", val)
	default:
		fmt.Println("No value ready, moving on.")
	}

	// defaults are used to unblock the channel
	select {
	case ch <- 7:
		fmt.Println("Sent 7 into channel")
	default:
		fmt.Println("Channel was full, could not send")
	}
}
