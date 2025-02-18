package main

import (
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("Greetings: ")
	log.SetFlags(0)

	names := []string{
		"Sebastian",
		"Sharon",
		"Sandra",
	}
	messages, err := Go.Hello(names)
	if err != nil {
		log.Fatal(err)
	}
	// message, err := greetings.Hello("Sebastian")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println(messages)
}
