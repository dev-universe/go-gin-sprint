package main

import (
	"fmt"
	"os"

	"github.com/dev-universe/go-gin-sprint/internal/day01"
)

func main() {

	args := os.Args
	name := ""

	if len(args) > 1 {
		name = args[1]
	}

	fmt.Println(day01.Greet(name))
}
