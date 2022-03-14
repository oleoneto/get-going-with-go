package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)


func init() {
	rand.Seed(time.Now().UnixNano())
}


func randomFormat() string {
	// A slice of message formats.
	formats := []string {
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Oi, %v! Prazer em conheçer você!",
		"Salut, %v! Bienvenue!",
	}

	return formats[rand.Intn(len(formats))]
}


func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("no name provided")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}
