package day04

import (
	"errors"
	"strings"
)

var ErrNoValidNames = errors.New("no valid names")

type Greeting struct {
	Name    string `json:"name"`
	Message string `json:"greeting"`
}

func BuildGreetings(names []string) ([]Greeting, error) {
	greetings := make([]Greeting, 0, len(names))

	for _, name := range names {
		n := strings.TrimSpace(name)
		if n == "" {
			continue
		}
		greetings = append(greetings, Greeting{
			Name:    n,
			Message: "Hello " + n + "!",
		})
	}

	if len(greetings) == 0 {
		return nil, ErrNoValidNames
	}

	return greetings, nil
}
