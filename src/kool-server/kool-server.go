package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"net/http"
)

type Result struct {
	ResponseTime int64  `json:"response_time"`
	Url          string `json:"url"`
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

	//TODO actually store the info
	w.WriteHeader(http.StatusOK)
	message["message"] = "Correctly saved (not really)"
	enc.Encode(&message)
}

func main() {
	fmt.Println("Starting server!")

	/* Initialize handlers */
	router := mux.NewRouter()
	router.HandleFunc("/hello", hello).Methods("GET")
	router.HandleFunc("/result", result).Methods("POST")

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":3000")
}
