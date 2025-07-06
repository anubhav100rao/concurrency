package mutex

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func RunCounter() {
	var wg sync.WaitGroup
	c := &Counter{}

	// Launch 1000 goroutines each incrementing the counter
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}
	wg.Wait()

	var value = 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			value++
		}()
	}
	wg.Wait()

	fmt.Println("Final counter:", c.Value()) // Always 1000
	fmt.Println("Final counter with race condition: ", value)
}
