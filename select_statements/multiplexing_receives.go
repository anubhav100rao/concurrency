package select_statements

import (
	"fmt"
	"time"
)

func ping(ch chan<- string) {
	time.Sleep(500 * time.Millisecond)
	ch <- "ping"
}

func pong(ch chan<- string) {
	time.Sleep(300 * time.Millisecond)
	ch <- "pong"
}

func RunMultiplexingReceives() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go ping(ch1)
	go pong(ch2)

	// select picks whichever channel is ready
	select {
	case msg := <-ch1:
		fmt.Println("Received from ch1:", msg)
	case msg := <-ch2:
		fmt.Println("Received from ch2:", msg)
	}

	// these introduce blocking wait
	// fmt.Println(<-ch1)
	// fmt.Println(<-ch2)
}
