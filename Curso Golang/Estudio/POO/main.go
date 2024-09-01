package main

import "POO/animal"

func main() {

	// myBook := book.NewBook("Moby Dick", "Herman Melville", 300)

	// // myBook.PrintInfo()
	// //Method to modify the book title
	// //myBook.SetTitle("Moby Dick SE")
	// //fmt.Println(myBook.GetTitle())
	// //Composition
	// myTextBook := book.NewTextBook("Matematicas", "Profe", 200, "Editorial", "Avanzado")
	// // myTextBook.PrintInfo()

	// book.Print(myBook)
	// book.Print(myTextBook)
	// miPerro := animal.Perro{Nombre: "Jack"}
	// miGato := animal.Gato{Nombre: "Michi"}
	// animal.MakeNoise(&miPerro)
	// animal.MakeNoise(&miGato)}
	animales := []animal.Animal{
		&animal.Perro{Nombre: "Jack"},
		&animal.Gato{Nombre: "Michi"},
		&animal.Perro{Nombre: "Max"},
		&animal.Gato{Nombre: "Lola"},
	}
	for _, animal := range animales {
		animal.Sonido()
	}
}
