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

func jobs_poller(agentId int, jobsChan chan TestList) error {
	/* This should authenticate AND start polling */
	polling_interval := 30
	serverURL := "http://localhost:3000"
	serverMethod := "alive"

	aliveData := make(map[string]interface{})
	aliveData["agentId"] = agentId

	sleep_interval := time.Duration(polling_interval) * time.Second

	fmt.Println("Done!")

	/* Poll the /alive endpoint */
	for {
		b, _ := json.Marshal(aliveData)
		reader := strings.NewReader(string(b))

		request, err := http.NewRequest("POST",
			fmt.Sprintf("%s/%s", serverURL, serverMethod),
			reader)

		if err != nil {
			return err
		}
		res, err := http.DefaultClient.Do(request)
		if err != nil {
			fmt.Printf("Could not contact server, %s", err)
		}

		const jsonData = `{
				"agentId": 3,
				"status": "OK",
				"jobs": [{
					"testId": 1,
					"targetURL": "http://www.segundamano.mx",
					"frecuency": 30
				}, {
					"testId": 2,
					"targetURL": "http://m.segundamano.mx",
					"frecuency": 45
				}]
			}`

		dec := json.NewDecoder(res.Body)
		dec = json.NewDecoder(strings.NewReader(jsonData))
		var testList TestList

		err = dec.Decode(&testList)
		if err != nil {
			fmt.Printf("Cannot decode JSON: %s", err)
			continue
		}

		fmt.Printf("Test List: %s\n",
			testList)

		jobsChan <- testList

		// XXX we should use a timer with alarms
		time.Sleep(sleep_interval)
	}
}
