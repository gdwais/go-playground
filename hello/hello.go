package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message
	message, err := greetings.Hello("")
	// if an error was returned, print it to the console and exit the program
	if err != nil {
		fmt.Println(err)
		//log.Fatal(err)
	} else {
		fmt.Println(message)
	}

	// if no error as returned print the returned message to the console

	// Get a greeting message and print it.
	new_message, new_err := greetings.Hello("Gladys")

	if new_err != nil {
		log.Fatal(new_err)
	}
	fmt.Println(new_message)
}
