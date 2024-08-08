package main

import "fmt"

func main() {
	nombre := prueba("Sebas")
	fmt.Println(nombre)
}

func prueba(name string) string {
	return "Hola mi nombre es " + name
}
