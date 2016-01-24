package main

import (
	"strings"
	"log"
	"bytes"
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
	AgentId		 int64	`json:"agent_id"`
	ResponseTime int64	`json:"response_time"`
	Url			 string `json:"url"`
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

func alive(w http.ResponseWriter, r *http.Request) {
	// Read the JSON from the request.
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	var dat map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &dat); err != nil {
		panic(err)
	}

	// Connect to the database
	db, err := sql.Open("postgres", "user=kool_writer dbname=monkey port=20010 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// Insert/update in the database the agent information.
	ip := strings.Split(r.RemoteAddr,":")[0]
	response := make(map[string]interface{});
	_, ok := dat["id"]
	if ok {
		_,err := db.Exec("UPDATE agent SET ip = $1, last_alive = now() WHERE id = $2", ip, dat["id"])
		if err != nil {
			fmt.Print(err)
			response["id"] = -1;
			response["status"] = "KO";
		} else {
			response["id"] = dat["id"];
			response["status"] = "OK";
		}
	} else {
		var id int
		err := db.QueryRow("INSERT INTO agent (ip, last_alive) VALUES ($1, now()) RETURNING id", ip).Scan(&id)
		if err != nil {
			fmt.Print(err)
			response["id"] = -1;
			response["status"] = "KO";
		} else {
			response["id"] = id
			response["status"] = "OK";
		}
	}

	// Sent the response to the agent
	enc := json.NewEncoder(w)
	w.WriteHeader(http.StatusOK)
	enc.Encode(&response)
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
	router.HandleFunc("/alive", alive).Methods("POST")

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":3000")
}
