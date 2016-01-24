package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

func performSingleTest(job *Job) (time.Duration, string, error) {

	var timeTaken time.Duration
	fmt.Printf("[Test %d] Calling PhantomJS with target URL: %s\n",
		job.TestId, job.TargetURL)

	t0 := time.Now()
	cmd := exec.Command("phantomjs", "scripts/phantomjs/netsniff.js", job.TargetURL)
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err := cmd.Run()
	outputString := cmdOutput.String()
	t1 := time.Now()

	timeTaken = t1.Sub(t0)

	if err != nil {
		return timeTaken, outputString, err
	}

	return timeTaken, outputString, nil
}

func uploadResults(job *Job, testResults string) error {
	// Do something
	resultData := make(map[string]interface{})
	resultData["url"] = job.TargetURL
	resultData["testResults"] = testResults
	resultData["agentId"] = AgentId

	b, _ := json.Marshal(resultData)
	reader := strings.NewReader(string(b))
	request, err := http.NewRequest("POST", fmt.Sprintf("%s/result", ServerURL), reader)
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

func job_runner(job *Job) {
	sleepInterval := time.Duration(job.Frequency) * time.Second
	fmt.Printf("[Test %d] Sleep interval will be %v.\n",
		job.TestId, sleepInterval)

	for {
		select {
		case <-job.CtrlChan:
			fmt.Printf("[Test %d] Received kill message, shutting down",
				job.TestId)
			return
		default:
			fmt.Printf("[Test %d] Targeting %s\n", job.TestId, job.TargetURL)
			duration, testResults, err := performSingleTest(job)

			if err != nil {
				fmt.Printf("[Test %d] ERROR: Error with time test: %s %s.\n",
					job.TestId, testResults, err)
			} else {
				fmt.Printf("[Test %d] Test completed after %v.\n",
					job.TestId, duration)
				uploadResults(job, testResults)
			}
			time.Sleep(sleepInterval)
		}
	}
}
