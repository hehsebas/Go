package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("hola.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	defer file.Close()

	_, err = file.Write([]byte("Holaaaa"))
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}
