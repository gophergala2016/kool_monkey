package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting agent!")

	// XXX
	agent_id := 42

	// Set up Channels
	jobsChannel := make(chan string)

	fmt.Print("Initializing Jobs Orchestrator... ")
	go jobs_orchestrator(jobsChannel)
	fmt.Println("Done!")

	fmt.Print("Initializing Jobs Poller... ")
	jobs_poller(agent_id, jobsChannel)
}
