package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	PORT := os.Getenv("PORT")

	// Load the env file.
	// No longer needs to use the "export PORT=value".
	godotenv.Load(".env")

	if PORT == "" {
		log.Fatal("PORT is not found in environment.")
	}

	fmt.Println(PORT)
}
