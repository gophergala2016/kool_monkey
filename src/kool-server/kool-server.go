package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
	"os"
	"strings"
)

var (
	DB *sql.DB
)

type Result struct {
	AgentId      int64  `json:"agentId"`
	ResponseTime int64  `json:"response_time"`
	Url          string `json:"url"`
}

func connectToDb() error {
	var err error
	DB, err = sql.Open("postgres", "host=127.0.0.1 port=20010 dbname=monkey user=kool_writer sslmode=disable")
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
	response := make(map[string]interface{})
	id, ok := dat["agentId"]
	if ok {
		_, err = DB.Exec("UPDATE agent SET ip = $1, last_alive = now() WHERE id = $2", ip, dat["agentId"])
		agentOk = (err == nil)
	} else {
		err = DB.QueryRow("INSERT INTO agent (ip, last_alive) VALUES ($1, now()) RETURNING id", ip).Scan(&id)
		agentOk = (err == nil)
	}

	if agentOk {
		response["agentId"] = id
		response["status"] = "OK"
		w.WriteHeader(http.StatusOK)
	} else {
		fmt.Print(err)
		response["agentId"] = -1
		response["status"] = "KO"
		response["message"] = "Couldn't update the agent"
		w.WriteHeader(http.StatusInternalServerError)
	}

	// testId, targetURL, frequency
	// Sent the response to the agent
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.Encode(&response)
}

func main() {
	fmt.Println("Starting static dashboard server at port 3002")
	go func() {
		panic(http.ListenAndServe(":3002", http.FileServer(http.Dir("/opt/kool_monkey/dashboard"))))
	}()

	fmt.Println("Starting static www server at port 3001")
	go func() {
		panic(http.ListenAndServe(":3001", http.FileServer(http.Dir("/opt/kool_monkey/www"))))
	}()

	fmt.Println("Starting api server at port 3000")

	err := connectToDb()
	if err != nil {
		fmt.Println("Couldn't connect to DB!")
		os.Exit(1)
	}

	/* Initialize handlers */
	router := mux.NewRouter()
	router.HandleFunc("/hello", hello).Methods("GET")
	router.HandleFunc("/result", result).Methods("POST")
	router.HandleFunc("/alive", alive).Methods("POST")

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":3000")
}
