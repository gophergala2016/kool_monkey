package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func sendResult(server, url string, timeTaken time.Duration) error {

	resultData := make(map[string]interface{})
	resultData["url"] = url
	resultData["response_time"] = int64(timeTaken / time.Microsecond) //time in microseconds
	resultData["agentId"] = 1

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

	if res.StatusCode != 200 {
		message := make(map[string]string)
		dec := json.NewDecoder(res.Body)
		err = dec.Decode(&message)
		if err != nil {
			return errors.New("Server didn't correctly save for some reason.")
		}
		err := errors.New(fmt.Sprintf("Server said: %s.", message["message"]))
		return err
	}

	return nil
}

func jobs_orchestrator(jobsChan chan TestList) error {
	// XXX
	serverURL := "http://localhost:3000"

	/* Poll the /alive endpoint */
	for {
		//newJob := <-jobsChan
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
