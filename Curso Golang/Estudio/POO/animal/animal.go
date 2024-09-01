package animal

import "fmt"

//Interface
type Animal interface {
	Sonido()
}

//Dog struct and methods
type Perro struct {
	Nombre string
}

func (p *Perro) Sonido() {
	fmt.Println(p.Nombre + "*Ladra*")
}

//Cat struct and methods
type Gato struct {
	Nombre string
}

func (g *Gato) Sonido() {
	fmt.Println(g.Nombre + "*Miau*")
}

//Func to make noise with struct
func MakeNoise(animal Animal) {
	animal.Sonido()
}
