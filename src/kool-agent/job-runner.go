package main

import (
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

func job_runner(targetUrl string, frequency int, resultChan chan time.Duration) {
	sleep_interval := time.Duration(frequency) * time.Second

	for {
		fmt.Println("Running test against %s!", targetUrl)
		timeTaken, err := urlTimeTest(targetUrl)

		if err != nil {
			fmt.Printf("ERROR: Error with time test: %s.\n", err)
		}

		resultChan <- timeTaken
		time.Sleep(sleep_interval)
	}
}
