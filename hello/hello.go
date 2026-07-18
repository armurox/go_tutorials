package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set the prefix for every log message
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := greetings.Hello("Arman")
	if (err != nil) {
		log.Fatal(err)
	}
	fmt.Println(message)
}
