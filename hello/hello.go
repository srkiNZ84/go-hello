package main

import (
	"fmt"
	"log"

	"github.com/srkiNZ84/go-greetings"
)

func main() {
	// Set properties of the predfined Logger
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names
	names := []string{"Gladys", "Samantha", "Darrin"}
	// Get a greeting message and print it.
	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
