package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/dev-universe/go-gin-sprint/internal/day03"
)

func main() {
	// Step 1: read args (excluding program name)
	names := []string{}
	if len(os.Args) > 1 {
		names = os.Args[1:]
	}

	// Step 2: build greeting data
	data := day03.BuildGreetings(names)

	// Step 3: convert to pretty JSON
	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}

	// Step 4: print JSON
	fmt.Println(string(jsonBytes))
}
