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
	names := []string{"Arman", "Raghav", "Keshav"}
	messages, err := greetings.Hellos(names)
	if (err != nil) {
		log.Fatal(err)
	}
	// messages example = map[Arman:Hi Arman. Welcome! Keshav:Hi Keshav. Welcome! Raghav:Hi Raghav. Welcome!]
	for _, message := range messages {
		fmt.Println(message) // This prints out the value, not the key (unlike python)
	}
	
}
