package main

import (
	"fmt"
	"os"

	"github.com/dev-universe/go-gin-sprint/internal/day02"
)

func main() {

	names := []string{}
	if len(os.Args) > 1 {
		names = os.Args[1:]
	}

	messages := day02.GreetAll(names)

	for _, msg := range messages {
		fmt.Println(msg)
	}
}
