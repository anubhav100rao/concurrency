package channels

import (
	"fmt"
	"sync"
)

// dummy job that returns its input squared
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j)
		results <- j * j
	}
}

func RunWorkerPools() {
	const numWorkers = 3
	const numJobs = 10

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	var wg sync.WaitGroup

	// start workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// wait for workers then close results
	go func() {
		wg.Wait()
		close(results)
	}()

	// collect results
	for res := range results {
		fmt.Println("Result:", res)
	}
}
