package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}

//func handlerReadiness(w http.ResponseWriter, r *http.Request)- important function signature
//used to define HTTP handler in go
