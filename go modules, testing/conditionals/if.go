package main

import "fmt"

const (
	Engprefix  = "Hello! "
	Spanprefix = "Hola! "
	Frenprefix = "Bonjour! "
)

func determinePrefix(lang string) (prefix string) {
	switch lang {
	case "Spanish":
		prefix = Spanprefix
	case "French":
		prefix = Frenprefix
	default:
		prefix = Engprefix
	}
	return
}

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}

	return determinePrefix(lang) + name
}

func main() {
	fmt.Println(Hello("Chris", "Eng"))
}
