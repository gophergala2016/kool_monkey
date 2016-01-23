package main

import (
	"fmt"
)

func jobs_orchestrator(jobsChan chan string) error {
	/* Poll the /alive endpoint */
	for {
		newJob := <-jobsChan

		// Handle a queue of pending Jobs

		// Poll for results from Jobs

		// Send results back to the server
		fmt.Println("Got new job: %s", newJob)
	}
}
