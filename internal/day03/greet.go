package day03

type Greeting struct {
	Name    string `json:"name"`
	Message string `json:"greeting"`
}

func BuildGreetings(names []string) []Greeting {
	greetings := make([]Greeting, 0, len(names))
	for _, name := range names {
		greetings = append(greetings, Greeting{
			Name:    name,
			Message: "Hello " + name + "!",
		})
	}
	return greetings
}
