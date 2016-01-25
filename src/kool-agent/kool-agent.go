package main

import (
	"fmt"
)

var AgentId int64 = -1
var ServerURL string = "http://api.koolmonkey.xyz"

func main() {
	fmt.Println("Starting agent!")

	// XXX This needs to be replaced with the value got from the registration
	AgentId = 1

	// Set up Channels
	jobsChannel := make(chan []SingleTest)

	fmt.Print("Initializing Jobs Orchestrator... ")
	go jobs_orchestrator(jobsChannel)
	fmt.Println("Done!")

	fmt.Print("Initializing Jobs Poller... ")
	jobs_poller(jobsChannel)
}
