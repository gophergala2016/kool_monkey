package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
	"os"
)

var (
	DB *sql.DB
)

type Result struct {
	AgentId      int64  `json:"agent_id"`
	ResponseTime int64  `json:"response_time"`
	Url          string `json:"url"`
}

func connectToDb() error {
	var err error
	DB, err = sql.Open("postgres", "host=127.0.0.1 port=24810 dbname=monkey user=kool_writer sslmode=disable")
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

func main() {
	fmt.Println("Starting server!")

	err := connectToDb()
	if err != nil {
		fmt.Println("Couldn't connect to DB!")
		os.Exit(1)
	}

	/* Initialize handlers */
	router := mux.NewRouter()
	router.HandleFunc("/hello", hello).Methods("GET")
	router.HandleFunc("/result", result).Methods("POST")

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":3000")
}
