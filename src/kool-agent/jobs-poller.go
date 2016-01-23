package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func jobs_poller(agent_id int) error {
	/* This should authenticate AND start polling */
	polling_interval := 60
	serverURL := "http://localhost:3000"
	serverMethod := "alive"

	aliveData := make(map[string]interface{})
	aliveData["agentId"] = agent_id

	b, _ := json.Marshal(aliveData)
	reader := strings.NewReader(string(b))

	sleep_interval := time.Duration(polling_interval) * time.Second

	/* Poll the /alive endpoint */
	for {
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

		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		response := buf.String()

		fmt.Printf("The response from the server was: %s",
			response)

		// XXX we should use a timer with alarms
		time.Sleep(sleep_interval)
	}
}
