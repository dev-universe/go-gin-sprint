package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/dev-universe/go-gin-sprint/internal/day04"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func main() {
	names := []string{}
	if len(os.Args) > 1 {
		names = os.Args[1:]
	}

	data, err := day04.BuildGreetings(names)

	if err != nil {
		b, marshalErr := json.MarshalIndent(ErrorResponse{Error: err.Error()}, "", "  ")
		if marshalErr != nil {
			fmt.Println(`{ "error": "failed to marshal error response" }`)
			return
		}
		fmt.Println(string(b))
		return
	}

	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println(`{ "error": "failed to marshal response" }`)
		return
	}

	fmt.Println(string(b))
}
