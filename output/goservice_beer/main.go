package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gorilla/handlers"

	"beer/controller"
	"fmt"
)

func main() {
	// Get the "PORT" env variable
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Print("$PORT must be set\n")
		port="8080"

	}
	//<<MAIN_PLACEHOLDER>>
	fmt.Println("API End-points::")
	router := controller.NewRouter() // create routes

	// These two lines are important if you're designing a front-end to utilise this API methods
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	fmt.Println("Listening on ::"+port)

	// Launch server with CORS validations
	log.Fatal(http.ListenAndServe(":" + port, handlers.CORS(allowedOrigins, allowedMethods)(router)))
}