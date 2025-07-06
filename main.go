package main

import (
	// "github.com/anubhav100rao/concurrency/select_statements"
	"github.com/anubhav100rao/concurrency/channels"
)

func main() {
	// select_statements.RunMultiplexingReceives()
	// select_statements.RunTimeouts()
	// select_statements.RunDefault()
	// select_statements.RunChannelCloseDetection()
	// select_statements.RunFanIn()
	// select_statements.RunPeriodicWork()

	// channels.RunWorkerPools()
	// channels.RunPipleline()
	// channels.RunRateLimiter()
	channels.RunFinInOutParallelFetching()
}
