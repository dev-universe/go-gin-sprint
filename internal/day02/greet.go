package day02

func GreetAll(names []string) []string {
	messages := []string{}

	for _, name := range names {
		messages = append(messages, "Hello "+name+"!")
	}

	return messages
}
