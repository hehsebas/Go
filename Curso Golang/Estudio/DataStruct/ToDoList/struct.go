package main

import "fmt"

type Person struct {
	nombre string
	edad   int
	genero string
	correo string
}

func main() {
	// var p Person
	// p.nombre = "Sebastian"
	// p.edad = 23
	// p.genero = "Masculino"
	// p.correo = "sebasrre@gmail.com"

	// fmt.Printf("Nombre: %s, Edad: %d, Genero: %s, Correo: %s\n", p.nombre, p.edad, p.genero, p.correo)

	p := Person{"Johan", 23, "Masculino", "sebasrre@gmail.com"}
	fmt.Println("Datos persona: \n", p)
	p2 := Person{"Sharon", 22, "Femenino", "sharon@gmail.com."}
	fmt.Println("Datos persona 2: \n", p2)

}
