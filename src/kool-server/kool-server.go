package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {

	enc := json.NewEncoder(w)
	hello := "Hello World!"
	w.WriteHeader(http.StatusOK)
	enc.Encode(&hello)
}

func result(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)

	dec := json.NewDecoder(r.Body)
	var resultData map[string]interface{}
	err := dec.Decode(&resultData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorMessage := make(map[string]string)
		errorMessage["message"] = "Invalid request"
		enc.Encode(&errorMessage)
		return
	}

	//TODO actually store the info
	w.WriteHeader(http.StatusOK)
	message := make(map[string]string)
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
