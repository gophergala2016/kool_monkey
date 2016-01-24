package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var (
	DB *sql.DB
)

type Configuration struct {
	DbConn DbConnection
}

type DbConnection struct {
	Host string
	Port int
	Name string
	User string
}

type Result struct {
	AgentId      int64  `json:"agentId"`
	ResponseTime int64  `json:"response_time"`
	Url          string `json:"url"`
}

type AliveResult struct {
	AgentId interface{}              `json:"agentId"`
	Status  string                   `json:"status"`
	Message string                   `json:"message"`
	Jobs    []map[string]interface{} `json:"jobs"`
}

type TestSite struct {
	TestId    int    `json:"test_id"`
	TargetUrl string `json:"target_url"`
	Frequency int    `json:"frequency"`
}

type Agent struct {
	AgentId   int       `json:"agent_id"`
	Ip        string    `json:"ip"`
	LastAlive time.Time `json:"last_alive"`
}

func connectToDb(db DbConnection) error {
	var err error
	connStr := fmt.Sprintf("host=%s port=%d dbname=%s user=%s sslmode=disable", db.Host, db.Port, db.Name, db.User)
	DB, err = sql.Open("postgres", connStr)
	return err
}

func hello(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	hello := "Hello World!"
	w.WriteHeader(http.StatusOK)
	enc.Encode(&hello)
}

func result(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	message := make(map[string]string)

	dec := json.NewDecoder(r.Body)
	var resultData Result
	err := dec.Decode(&resultData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		message["message"] = "Invalid request"
		enc.Encode(&message)
		return
	}

	_, err = DB.Exec(
		"INSERT INTO result (agent_id, url, response_time) VALUES ($1, $2, $3)",
		resultData.AgentId,
		resultData.Url,
		resultData.ResponseTime,
	)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		message["message"] = "Couldn't save result"
		enc.Encode(&message)
		return
	}

	w.WriteHeader(http.StatusOK)
	message["message"] = "Correctly saved"
	enc.Encode(&message)
}

func alive(w http.ResponseWriter, r *http.Request) {
	var agentOk bool
	var err error

	// Read the JSON from the request.
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	var dat map[string]interface{}
	if err = json.Unmarshal(buf.Bytes(), &dat); err != nil {
		panic(err)
	}

	// Insert/update in the database the agent information.
	ip := strings.Split(r.RemoteAddr, ":")[0]
	id, ok := dat["agentId"]
	if ok {
		_, err = DB.Exec("UPDATE agent SET ip = $1, last_alive = now() WHERE id = $2", ip, dat["agentId"])
		agentOk = (err == nil)
	} else {
		err = DB.QueryRow("INSERT INTO agent (ip, last_alive) VALUES ($1, now()) RETURNING id", ip).Scan(&id)
		agentOk = (err == nil)
	}

	// Prepare and send the response to the agent
	var response AliveResult
	if agentOk {
		response.AgentId = id
		response.Status = "OK"
		w.WriteHeader(http.StatusOK)

		rows, errQuery := DB.Query("SELECT test.id, test.targetURL, test.frequency FROM test INNER JOIN testAgent ON test.id = testAgent.idTest WHERE testAgent.idAgent = $1", id)
		if errQuery == nil {
			var testId int
			var targetUrl string
			var frecuency int
			for i := 0; rows.Next(); i++ {
				job := make(map[string]interface{})
				rows.Scan(&testId, &targetUrl, &frecuency)

				// If Jobs is full it must grow.
				if i == cap(response.Jobs) {
					newSlice := make([]map[string]interface{}, len(response.Jobs), 2*len(response.Jobs)+1)
					copy(newSlice, response.Jobs)
					response.Jobs = newSlice
				}

				job["testId"] = testId
				job["targetURL"] = targetUrl
				job["frequency"] = frecuency

				response.Jobs = append(response.Jobs, job)
			}
			rows.Close()
		} else {
			fmt.Print(errQuery)
		}
	} else {
		fmt.Print(err)
		response.AgentId = -1
		response.Status = "KO"
		response.Message = "Couldn't update the agent"
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.Encode(&response)
}

func addSite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	response := make(map[string]interface{})

	var okUrl, okFreq bool
	var targetUrl string
	var frequency float64
	var testId int

	dec := json.NewDecoder(r.Body)
	siteTest := make(map[string]interface{})
	err := dec.Decode(&siteTest)
	if err == nil {
		targetUrl, okUrl = siteTest["targetUrl"].(string)
		frequency, okFreq = siteTest["frequency"].(float64)
	}
	if err != nil || !okUrl || !okFreq {
		w.WriteHeader(http.StatusBadRequest)
		response["status"] = "KO"
		response["message"] = "Invalid JSON"
		enc.Encode(&response)
		return
	}

	err = DB.QueryRow(
		"INSERT INTO test (targetUrl, frequency) VALUES ($1, $2) RETURNING id",
		targetUrl,
		int(frequency),
	).Scan(&testId)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		response["status"] = "KO"
		response["message"] = "Couldn't save test"
		enc.Encode(&response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response["status"] = "OK"
	response["message"] = "Correctly saved"
	response["testId"] = testId
	enc.Encode(&response)
}

func getSites(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	response := make(map[string]interface{})

	testId := 0
	testIdStr := r.FormValue("test_id")
	if testIdStr != "" {
		var err error
		testId, err = strconv.Atoi(testIdStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response["status"] = "KO"
			response["message"] = "Invalid test ID"
			enc.Encode(&response)
			return
		}
	}

	rows, err := DB.Query("SELECT id, targetUrl, frequency FROM test WHERE id = $1 OR $1 = 0", testId)
	defer rows.Close()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response["status"] = "KO"
		response["message"] = "Couldn't get tests"
		enc.Encode(&response)
		return
	}

	tests := make([]TestSite, 0)
	for rows.Next() {
		var testSite TestSite
		rows.Scan(&testSite.TestId, &testSite.TargetUrl, &testSite.Frequency)
		tests = append(tests, testSite)
	}

	w.WriteHeader(http.StatusOK)
	response["status"] = "OK"
	response["test_sites"] = tests
	enc.Encode(&response)
}

func getAgents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	response := make(map[string]interface{})

	rows, err := DB.Query("SELECT id, ip, last_alive FROM agent")
	defer rows.Close()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response["status"] = "KO"
		response["message"] = "Couldn't get agents"
		enc.Encode(&response)
		return
	}

	agents := make([]Agent, 0)
	for rows.Next() {
		var a Agent
		rows.Scan(&a.AgentId, &a.Ip, &a.LastAlive)
		agents = append(agents, a)
	}

	w.WriteHeader(http.StatusOK)
	response["status"] = "OK"
	response["agents"] = agents
	enc.Encode(&response)
}

func main() {
	koolDir, err := filepath.Abs(filepath.Dir(os.Args[0]) + "/../")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Starting static dashboard server at port 3002")
	go func() {
		panic(http.ListenAndServe(":3002", http.FileServer(http.Dir(koolDir+"/dashboard"))))
	}()

	fmt.Println("Starting static www server at port 3001")
	go func() {
		panic(http.ListenAndServe(":3001", http.FileServer(http.Dir(koolDir+"/www"))))
	}()

	fmt.Println("Starting api server at port 3000")

	//Read config
	cmd_cfg := flag.String("conf", koolDir+"/conf/kool-server.conf", "Config file")
	flag.Parse()
	file, err := os.Open(*cmd_cfg)
	if err != nil {
		fmt.Printf("Config - File Error: %s\n", err)
		os.Exit(1)
	}

	decoder := json.NewDecoder(file)
	conf := Configuration{}

	if err := decoder.Decode(&conf); err != nil {
		fmt.Printf("Config - Decoding Error: %s\n", err)
		os.Exit(1)
	}

	//Connect to DB
	err = connectToDb(conf.DbConn)
	if err != nil {
		fmt.Println("Couldn't connect to DB!")
		os.Exit(1)
	}

	/* Initialize handlers */
	router := mux.NewRouter()
	router.HandleFunc("/hello", hello).Methods("GET")
	router.HandleFunc("/result", result).Methods("POST")
	router.HandleFunc("/alive", alive).Methods("POST")
	router.HandleFunc("/sites", addSite).Methods("POST")
	router.HandleFunc("/sites", getSites).Methods("GET")
	router.HandleFunc("/agents", getAgents).Methods("GET")

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":3000")
}
