package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type SingleTest struct {
	TestId    int64  `json:"testId"`
	TargetURL string `json:"targetURL"`
	Frequency int64  `json:"frequency"`
}

type TestList struct {
	AgentId int64        `json:"agentId"`
	Status  string       `json:"status"`
	Jobs    []SingleTest `json:"jobs,omitempty"`
}

func jobs_poller(jobsChan chan []SingleTest) error {
	/* This should authenticate AND start polling */
	polling_interval := 30
	serverMethod := "alive"

	aliveData := make(map[string]interface{})
	aliveData["agentId"] = AgentId

	sleep_interval := time.Duration(polling_interval) * time.Second

	fmt.Println("Done!")

	/* Poll the /alive endpoint */
	for {
		b, _ := json.Marshal(aliveData)
		reader := strings.NewReader(string(b))

		request, err := http.NewRequest("POST",
			fmt.Sprintf("%s/%s", ServerURL, serverMethod),
			reader)

		if err != nil {
			return err
		}
		res, err := http.DefaultClient.Do(request)
		if err != nil {
			fmt.Printf("Could not contact server, %s", err)
		}

		dec := json.NewDecoder(res.Body)

		var testList TestList
		err = dec.Decode(&testList)
		if err != nil {
			fmt.Printf("Cannot decode JSON: %s", err)
			continue
		}

		jobsChan <- testList.Jobs

		// XXX we should use a timer with alarms
		time.Sleep(sleep_interval)
	}
}
