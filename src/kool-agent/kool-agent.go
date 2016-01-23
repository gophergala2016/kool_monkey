package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func urlTimeTest(url string) (time.Duration, error) {

	var timeTaken time.Duration

	reader := strings.NewReader("")
	request, err := http.NewRequest("GET", url, reader)
	if err != nil {
		return timeTaken, err
	}

	t0 := time.Now()
	_, err = http.DefaultClient.Do(request)
	t1 := time.Now()
	if err != nil {
		return timeTaken, err
	}

	timeTaken = t1.Sub(t0)
	return timeTaken, nil
}

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

func main() {
	fmt.Println("Starting agent!")

	// XXX
	agent_id := 42

	// Set up Channels
	jobsChannel := make(chan string)

	fmt.Print("Initializing Jobs Poller... ")
	go jobs_poller(agent_id, jobsChannel)
	fmt.Println("Done!")

	fmt.Print("Initializing Jobs Orchestrator... ")
	go jobs_orchestrator(jobsChannel)
	fmt.Println("Done!")

	serverURL := "http://localhost:3000"
	interval := time.Duration(30) * time.Second
	url := "http://www.segundamano.mx"

	for {
		fmt.Println("Starting time test!")
		timeTaken, err := urlTimeTest(url)
		if err != nil {
			fmt.Printf("ERROR: Error with time test: %s.\n", err)
		} else {
			fmt.Printf("Time taken was: %v.\n", timeTaken)
			err := sendResult(serverURL, url, timeTaken)
			if err != nil {
				fmt.Printf("ERROR: Error sending result to server: %s.\n", err)
			}
		}

		time.Sleep(interval)
	}
}
