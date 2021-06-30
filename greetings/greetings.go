package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	//if no name was given, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}
	// return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// init sets tintial values for variables used in the function
func init() {
	rand.Seed(time.Now().UnixNano())
}

//randomForamt returns one of a set of greeting messages.  The returned message is selected at random.
func randomFormat() string {
	// a slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you. %v",
		"Hail, %v and well met!",
	}

	// return a randomly selected message format by specifying a random index for the slice of formats
	return formats[rand.Intn(len(formats))]
}
