package greetings

import (
	"fmt"
	"errors"
	"math/rand"
	"time"
)


func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	//message := fmt.Sprintf("Hello %v!", name)
	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}


func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	// Looping a slice array
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	// A slice is an dynamic array with no pre-defined size.
	var formats []string
	formats = []string{
		"Hi, %v. Welcome",
		"Great to see you %v",
		"Hail %v", // Note comma at the end.
	}
	formats = append(formats,"Welcome %v")

	return formats[rand.Intn(len(formats))]
}
