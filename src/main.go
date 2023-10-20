package main

import "fmt"
import "log"
import "os"
import "github.com/joho/godotenv"

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

	fmt.Println(PORT)
}

