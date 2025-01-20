package main

import (
	"fmt"
	"log"

	"rsc.io/quote"

	"example.com/greetings"
)

func main() {
	fmt.Println(quote.Go())

	log.SetPrefix("greetings:")
	log.SetFlags(0)
	names := []string{"Yash", "Jake"}
	message, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
