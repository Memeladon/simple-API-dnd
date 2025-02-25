package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello World")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("$PORT must be set")
	}

	fmt.Println("Port:", portString)
}
