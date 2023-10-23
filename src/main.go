package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/kentlouisetonino/aggreflow/src/helper"
)

func main() {
  // Load the env file.
  // No longer needs to use the "export PORT=value".
  err := godotenv.Load()

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

  // Setup the cors to allow any request from the browser.
  router.Use(cors.Handler(cors.Options{
    AllowedOrigins: []string{"https://*", "http://*"},
    AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders: []string{"*"},
    ExposedHeaders: []string{"Link"},
    AllowCredentials: false,
    MaxAge: 300,
  }))

  // For handling the passed json value.
  routerV1 := chi.NewRouter()
  routerV1.HandleFunc("/health", helper.HandleHealth)
  routerV1.Get("/error", helper.HandleError)

  // Nest the routerV2.
  // The full path for this is /v1t /health
  router.Mount("/v1", routerV1)

  // Setup the server.
  server := &http.Server{
    Handler: router,
    Addr: ":" + PORT,
  }

  log.Printf("Server starting on port %v.", PORT)

  // Listen to the server.
  err = server.ListenAndServe()

  if err != nil {
    log.Fatal(err)
  }
}

