package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/kentlouisetonino/aggreflow/database/sqlc"
	"github.com/kentlouisetonino/aggreflow/src/helper"
	"github.com/kentlouisetonino/aggreflow/src/libs"
	_ "github.com/lib/pq"
)

func main() {
  // Load the env file.
  // No longer needs to use the "export PORT=value".
  errorGoDotenv := godotenv.Load()
  if errorGoDotenv != nil {
    fmt.Println(errorGoDotenv)
  }

  // Get the port environment variable.
  PORT := os.Getenv("PORT")
  if PORT == "" {
    log.Fatal(helper.Red + "[ERROR]" + helper.Reset + " PORT is not found in environment.")
  }

  // Get the DB__UR; environment variable.
  DB_URL := os.Getenv("DB_URL")
  if DB_URL == "" {
    log.Fatal(helper.Red + "[ERROR]" + helper.Reset + " DB_URL is not found in environment.")
  }

  // Connect to the database.
  postgresConnection, postgresConnectionError := sql.Open("postgres", DB_URL)
  if postgresConnectionError != nil {
    fmt.Println(postgresConnectionError)
    log.Fatal(helper.Red + "[ERROR]" + helper.Reset + " Cannot connect to the database.")
  }

  // API configuration
  apiConfig := helper.DatabaseConfig{
    DB: sqlc.New(postgresConnection),
  }

  // Initialize the router.
  router := chi.NewRouter()

  // Setup the cors to allow any request from the browser.
  router.Use(cors.Handler(cors.Options{
    AllowedOrigins:   []string{"https://*", "http://*"},
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"*"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: false,
    MaxAge:           300,
  }))

  // Nest the routerV2.
  // The full path for this is /v1t /health
  router.Mount("/api", libs.Router(&apiConfig))

  // Setup the server.
  server := &http.Server{
    Handler: router,
    Addr:    ":" + PORT,
  }

  // Color Reset: \033[0m
  // Color Blue: \033[34m
  log.Printf("\033[34m[INFO]\033[0m Server starting on port %v", PORT)

  // Listen to the server.
  serverError := server.ListenAndServe()
  if serverError != nil {
    log.Fatal(serverError)
  }
}
