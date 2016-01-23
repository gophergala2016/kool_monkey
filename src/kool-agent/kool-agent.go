package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Starting agent!")

	serverURL := "http://localhost:3000"

	reader := strings.NewReader("")
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/hello", serverURL), reader)
	if err != nil {
		fmt.Printf("ERROR: Could not create request.\n")
	} else {

		res, err := http.DefaultClient.Do(request)
		if err != nil {
			fmt.Printf("ERROR: Could not send request to %s/hello.\n", serverURL)
		} else {

			buf := new(bytes.Buffer)
			buf.ReadFrom(res.Body)
			response := buf.String()

			fmt.Printf("The response was: %s\n", response)
		}
	}
}
