package main

import (
	"github.com/anubhav100rao/concurrency/select_statements"
)

func main() {
	// select_statements.RunMultiplexingReceives()
	// select_statements.RunTimeouts()
	// select_statements.RunDefault()
	// select_statements.RunChannelCloseDetection()
	select_statements.RunFanIn()
}
