package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type AgentConf struct {
	ServerURL string
	AgentId   int64
}

var Conf AgentConf = AgentConf{}

func main() {
	fmt.Println("Starting agent!")
	topDir, err := filepath.Abs(filepath.Dir(os.Args[0]) + "/../")

	cmd_cfg := flag.String("conf", topDir+"/conf/kool-agent.conf", "Config file")
	flag.Parse()
	file, err := os.Open(*cmd_cfg)
	if err != nil {
		fmt.Printf("Config - File Error: %s\n", err)
		os.Exit(1)
	}

	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&Conf); err != nil {
		fmt.Printf("Config - Decoding Error: %s\n", err)
		os.Exit(1)
	}

	// Set up Channels
	jobsChannel := make(chan []SingleTest)

	fmt.Print("Initializing Jobs Orchestrator... ")
	go jobs_orchestrator(jobsChannel)
	fmt.Println("Done!")

	fmt.Print("Initializing Jobs Poller... ")
	jobs_poller(jobsChannel)
}
