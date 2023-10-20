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

	if err != nil {
		fmt.Println(err)
	}

	PORT := os.Getenv("PORT")

	if PORT == "" {
		log.Fatal("PORT is not found in environment.")
	}

	router := chi.NewRouter()

	server := &http.Server{
		Handler: router,
		Addr:    ":" + PORT,
	}

	log.Printf("Server starting on port %v.", PORT)

	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
