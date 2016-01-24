package main

import (
	"fmt"
)

type Job struct {
	TestId    int64
	TargetURL string
	Frequency int64
	CtrlChan  chan string
}

var jobsList map[int64]*Job

func jobs_orchestrator(jobsChan chan []SingleTest) error {
	// XXX
	//serverURL := "http://localhost:3000"

	/* Poll the /alive endpoint */
	for {
		newJobsList := <-jobsChan
		fmt.Printf("[Orchestrator] Parsing new Jobs list\n")
		var activeJobs map[int64]int64 = make(map[int64]int64)

		// Handle a queue of pending Jobs
		for _, job := range newJobsList {
			if _, ok := jobsList[job.TestId]; ok {
				// Existing Job, check if we need to update details
				oldJob := jobsList[job.TestId]

				if job.Frequency != oldJob.Frequency {
					jobsList[job.TestId].Frequency = job.Frequency
				}
				if job.TargetURL != oldJob.TargetURL {
					jobsList[job.TestId].TargetURL = job.TargetURL
				}
			} else {
				// New Job, add it to the jobs list
				var newJob Job
				newJob.TestId = job.TestId
				newJob.TargetURL = job.TargetURL
				newJob.Frequency = job.Frequency

				newJob.CtrlChan = make(chan string)

				go job_runner(&newJob)
			}
			activeJobs[job.TestId] = 1
		}

		// If not in the list, kill the relevant goroutine
		for index, _ := range jobsList {
			if _, ok := activeJobs[index]; ok {
			} else {
				// Kill the goroutine, bye
			}
		}

		// Poll the dataChannels
		//duration := <-jobChan

		// Send results back to the server
		//fmt.Printf("Time taken was: %v.\n", duration)
		//err := sendResult(serverURL, targetURL, duration)

		//if err != nil {
		//	fmt.Printf("ERROR: Error sending result to server: %s.\n", err)
		//}
	}
}
