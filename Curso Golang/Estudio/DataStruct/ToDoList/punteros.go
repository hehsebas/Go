package main

import "fmt"

type Person struct {
	nombre string
	edad   int
	genero string
	correo string
}

func (p *Person) hello() {
	fmt.Println("Hola, mi nombre es: ", p.nombre)
}
func main() {
	p := Person{"Sebastian", 23, "Masculino", "sebasrre@gmail.com"}
	p.hello()
}
