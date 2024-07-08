package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

// JSON REST API: all the request bodies coming in and out have JSON format
func main() {
	godotenv.Load(".env")
	//To read the port in .env file
	portString := os.Getenv("PORT")
	if portString == "" {
		//exit the program with the message in the parentheses
		log.Fatal("Port not found in environment")
	}
	router := chi.NewRouter()
	//chi is a lightweight router for our information
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	//cors are like what our requests allow and block
	//v1Router to hook it to the handler function
	v1Router := chi.NewRouter()
	//get to make sure only 'Get' request is handled
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerError)
	router.Mount("/v1", v1Router)
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	log.Printf("Server starting on port %v", portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

//go get github.com/joho/godotenv
//The above command downloads the package required to extract the port number from the .env file
