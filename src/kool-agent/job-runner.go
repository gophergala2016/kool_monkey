package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"
)

func performSingleTest(url string) (time.Duration, string, error) {

	var timeTaken time.Duration

	t0 := time.Now()
	cmd := exec.Command("phantomjs", "scripts/phantomjs/netsniff.js", url)
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err := cmd.Run()
	t1 := time.Now()

	outputString := cmdOutput.String()

	if err != nil {
		return timeTaken, outputString, err
	}

	timeTaken = t1.Sub(t0)
	return timeTaken, outputString, nil
}

func job_runner(targetUrl string, frequency int, resultChan chan time.Duration) {
	sleep_interval := time.Duration(frequency) * time.Second

	for {
		fmt.Println("Running test against %s!", targetUrl)
		timeTaken, testResults, err := performSingleTest(targetUrl)

		fmt.Printf("%s", testResults)

		if err != nil {
			fmt.Printf("ERROR: Error with time test: %s.\n", err)
		}

		resultChan <- timeTaken
		time.Sleep(sleep_interval)
	}
}
