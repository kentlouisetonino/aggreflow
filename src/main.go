package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	// Load the env file.
	// No longer needs to use the "export PORT=value".
	err := godotenv.Load(".env")

	// Check if their is an error.
	if err != nil {
		fmt.Println(err)
	}

	// Get the port environment variable.
	PORT := os.Getenv("PORT")

	// Check again to make sure if it exist or not.
	if PORT == "" {
		log.Fatal("PORT is not found in environment.")
	}

	// Initialize the router.
	router := chi.NewRouter()

	// Setup the server.
	server := &http.Server{
		Handler: router,
		Addr:    ":" + PORT,
	}

	log.Printf("Server starting on port %v.", PORT)

	// Listen to the server.
	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
