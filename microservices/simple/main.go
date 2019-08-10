package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
)

type helloResponse struct {
	Message string `json:"message"`
}

type helloRequest struct {
	Name string `json:"name"`
}

func main () {
	port := 8080

	http.Handle("/hello", newValidationHandler(newHelloWorldHandler()))

	log.Printf("Server starting on port %v\n", 8080)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}



type helloWorldHandler struct{}

func newHelloWorldHandler() http.Handler {
	return helloWorldHandler{}
}

func (h helloWorldHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// var request helloRequest
	response := helloResponse{Message: "Hello "}

	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}


// seperate request validation
type validationHandler struct {
	next http.Handler
}

func newValidationHandler(next http.Handler) http.Handler {
	return validationHandler{next: next}
}

func (h validationHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var request helloRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(rw, "Bad request", http.StatusBadRequest)
		return
	}

	h.next.ServeHTTP(rw, r)
}


// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	var request helloRequest

// 	decoder := json.NewDecoder(r.Body)

// 	err := decoder.Decode(&request)
// 	if err != nil {
// 		http.Error(w, "Bad request", http.StatusBadRequest)
// 		return
// 	}

// 	response := helloResponse{Message: "Hello " + request.Name}

// 	encoder := json.NewEncoder(w)
// 	encoder.Encode(&response)
// }