package main

/*since it is main-package, go build will create an executable with a name of directory. All go files within this directory should belong to the
same package and the directory name should be in lowercase.*/

import "fmt"

const englishHelloPrefix = "Hello, "
const kannadaHelloPrefix = "Namaskara, "

// since the function name starts with uppercase, it can be imported in other files
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	var prefix string
	switch language {
	case "English":
		prefix = englishHelloPrefix
	case "Kannada":
		prefix = kannadaHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("Yash", ""))
}
