package main

import (
	"encoding/json" //provides the function to marshal the code
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		//error codes after 400 are client side errors so we donot need it
		log.Println("Responding with 5XX error:", msg)
	}
	type errResponse struct {
		Error string `json: "error"`
	}
	//in go, json tags like the one in the above struct are used to define how the code will Marshal the data
	respondWithJSON(w, code, errResponse{
		Error: msg,
	})

}

// Marshal means to convert the datatypes in JSON- formatted string. It is useful when we
// want to send data over an HTTP request and store it
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	//payload is used to Marshal the code
	//it returs the data as bytes so it can written as binary to HTTP
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}
	//Response
	w.Header().Add("Content- Type", "application/json")
	//Status code signifying the code that tells its successful
	w.WriteHeader(code)
	w.Write(dat)
}
