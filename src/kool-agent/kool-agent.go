package main

import (
	"fmt"
)

var AgentId int64 = -1
var ServerURL string = "http://localhost:3000"

func main() {
	fmt.Println("Starting agent!")

	// XXX
	AgentId = 1

	// Set up Channels
	jobsChannel := make(chan []SingleTest)

	fmt.Print("Initializing Jobs Orchestrator... ")
	go jobs_orchestrator(jobsChannel)
	fmt.Println("Done!")

	fmt.Print("Initializing Jobs Poller... ")
	jobs_poller(jobsChannel)
}
