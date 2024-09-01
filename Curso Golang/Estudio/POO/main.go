package main

import (
	"POO/book"
)

func main() {

	myBook := book.NewBook("Moby Dick", "Herman Melville", 300)

	// myBook.PrintInfo()
	//Method to modify the book title
	//myBook.SetTitle("Moby Dick SE")
	//fmt.Println(myBook.GetTitle())
	//Composition
	myTextBook := book.NewTextBook("Matematicas", "Profe", 200, "Editorial", "Avanzado")
	// myTextBook.PrintInfo()

	book.Print(myBook)
	book.Print(myTextBook)
}
