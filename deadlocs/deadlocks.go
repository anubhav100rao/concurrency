package deadlocks

import "sync"

func RunDeadLocks() {
	var mu1, mu2 sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)

	// Goroutine A
	go func() {
		defer wg.Done()
		mu1.Lock()
		defer mu1.Unlock()

		// Attempt to lock mu2 which may be held by B
		mu2.Lock()
		defer mu2.Unlock()
	}()

	// Goroutine B
	go func() {
		defer wg.Done()
		mu2.Lock()
		defer mu2.Unlock()

		// Attempt to lock mu1 which may be held by A
		mu1.Lock()
		defer mu1.Unlock()
	}()

	wg.Wait() // Deadlock: neither A nor B can acquire the second lock
}
