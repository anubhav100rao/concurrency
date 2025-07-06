package select_statements

import "fmt"

func RunChannelCloseDetection() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
		close(ch)
	}()

	// not recommended
	// for {
	// 	select {
	// 	case v, ok := <-ch:
	// 		if !ok {
	// 			fmt.Println("Channel closed; exiting")
	// 			return
	// 		}
	// 		fmt.Println("Received:", v)
	// 	}
	// }

	for v := range ch {
		fmt.Println("Received: ", v)
	}
}
