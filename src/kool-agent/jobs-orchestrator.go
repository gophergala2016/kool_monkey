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

var jobsList map[int64]*Job = make(map[int64]*Job)

func jobs_orchestrator(jobsChan chan []SingleTest) error {
	/* Poll the /alive endpoint */
	for {
		newJobsList := <-jobsChan
		fmt.Printf("[Orchestrator] Parsing new Jobs list\n")
		var activeJobs map[int64]int64 = make(map[int64]int64)

		// Handle a queue of pending Jobs
		for _, job := range newJobsList {
			if _, ok := jobsList[job.TestId]; ok {
				// Existing Job, check if we need to update details
				fmt.Printf("[Orchestrator] Job ID %d already set up, checking for changes\n",
					job.TestId)
				oldJob := jobsList[job.TestId]

				if job.Frequency != oldJob.Frequency {
					jobsList[job.TestId].Frequency = job.Frequency
				}
				if job.TargetURL != oldJob.TargetURL {
					jobsList[job.TestId].TargetURL = job.TargetURL
				}
			} else {
				fmt.Printf("[Orchestrator] Got a new Job with ID %d creating a new process\n",
					job.TestId)
				// New Job, add it to the jobs list
				var newJob Job
				newJob.TestId = job.TestId
				newJob.TargetURL = job.TargetURL
				newJob.Frequency = job.Frequency

				newJob.CtrlChan = make(chan string)

				jobsList[job.TestId] = &newJob
				go job_runner(&newJob)
			}
			activeJobs[job.TestId] = 1
		}

		// If not in the list, kill the relevant goroutine
		for index, _ := range jobsList {
			if _, ok := activeJobs[index]; !ok {
				fmt.Printf("[Orchestrator] Job ID %d has been disabled, shutting it down\n",
					index)
				jobsList[index].CtrlChan <- "die"
				delete(jobsList, index)
			}
		}
	}
}
