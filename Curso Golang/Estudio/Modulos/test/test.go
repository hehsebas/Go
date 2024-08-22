package main

import (
	"fmt"
	"log"

	"github.com/hehsebas/module"
)

func main() {
	log.SetPrefix("Module: ")
	log.SetFlags(0)

	message, error := module.Hello("Sebastian")
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(message)
}
