package helloworld

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"

	englishPrefix = "Hello, "
	spanishPrefix = "Hola, "
	frenchPrefix  = "Bonjour, "
)

func Hello(name, language string) string {
	return greetingPrefix(language) + normalizeName(name)
}

func normalizeName(name string) string {
	if name == "" {
		return "World"
	}
	return name
}

func greetingPrefix(language string) string {
	switch language {
	case french:
		return frenchPrefix
	case spanish:
		return spanishPrefix
	default:
		return englishPrefix
	}
}

func main() {
	fmt.Println(Hello("world", ""))
}
