package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	portString := os.Getenv("PORT")

	// Load the env file.
	// No longer needs to use the "export PORT=value".
	godotenv.Load(".env")

	if portString == "" {
		log.Fatal("PORT is not found in environment.")
	}

	fmt.Println(portString)
}
