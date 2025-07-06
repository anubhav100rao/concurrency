package select_statements

import (
	"fmt"
	"time"
)

func producer(name string, ch chan<- string) {
	for i := 1; i <= 3; i++ {
		ch <- fmt.Sprintf("%s: %d", name, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func fanIn(chs ...<-chan string) <-chan string {
	out := make(chan string)
	for _, ch := range chs {
		go func(c <-chan string) {
			for msg := range c {
				out <- msg
			}
		}(ch)
	}
	return out
}

func RunFanIn() {
	a := make(chan string)
	b := make(chan string)

	go producer("A", a)
	go producer("B", b)

	merged := fanIn(a, b)

	for i := 0; i < 6; i++ {
		fmt.Println(<-merged)
	}

	// we should call this only when channel is closed
	// for message := range merged {
	// 	fmt.Println(message)
	// }
}
