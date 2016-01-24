package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func sendResult(server, url string, timeTaken time.Duration) error {

	resultData := make(map[string]interface{})
	resultData["url"] = url
	resultData["time"] = int64(timeTaken / time.Microsecond) //time in microseconds

	b, _ := json.Marshal(resultData)
	reader := strings.NewReader(string(b))
	request, err := http.NewRequest("POST", fmt.Sprintf("%s/result", server), reader)
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	response := buf.String()

	fmt.Printf("The response from the server was: %s", response)
	return nil
}

func jobs_orchestrator(jobsChan chan string) error {
	// XXX
	serverURL := "http://localhost:3000"

	/* Poll the /alive endpoint */
	for {
		newJob := <-jobsChan
		fmt.Println("Got new job: %s", newJob)

		// Handle a queue of pending Jobs

		// XXX Use real data
		targetURL := "http://www.segundamano.mx"
		jobFrequency := 120

		// XXX Use a custom struct
		jobChan := make(chan time.Duration)

		// Poll for results from Jobs
		go job_runner(targetURL, jobFrequency, jobChan)

		duration := <-jobChan

		// Send results back to the server
		fmt.Printf("Time taken was: %v.\n", duration)
		err := sendResult(serverURL, targetURL, duration)

		if err != nil {
			fmt.Printf("ERROR: Error sending result to server: %s.\n", err)
		}
	}
}
