package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/dev-universe/go-gin-sprint/internal/day05"
)

type ErrorResponse struct {
	Error  string `json:"error"`
	Detail string `json:"detail,omitempty"`
}

func main() {
	names := []string{}
	if len(os.Args) > 1 {
		names = os.Args[1:]
	}

	data, err := day05.BuildGreetings(names)
	if err != nil {
		resp := ErrorResponse{}

		if errors.Is(err, day05.ErrNoValidNames) {
			resp.Error = "no valid names"
		} else if errors.Is(err, day05.ErrNameTooLong) {
			resp.Error = "name too long"
			resp.Detail = err.Error()
		} else {
			resp.Error = "internal error"
		}

		b, mErr := json.MarshalIndent(resp, "", "  ")
		if mErr != nil {
			fmt.Println(`{ "error": "failed to marshal error response" }`)
			return
		}
		fmt.Println(string(b))
		return
	}

	b, mErr := json.MarshalIndent(data, "", "  ")
	if mErr != nil {
		fmt.Println(`{ "error": "failed to marshal response" }`)
		return
	}
	fmt.Println(string(b))
}
