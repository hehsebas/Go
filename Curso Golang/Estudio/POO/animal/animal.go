package animal

import "fmt"

//Interface
type Animal interface {
	sonido()
}

//Dog struct and methods
type perro struct {
	Nombre string
}

func (p *perro) sonido() {
	fmt.Println(p.Nombre + "*Ladra*")
}

//Cat struct and methods
type gato struct {
	Nombre string
}

func (g *gato) sonido() {
	fmt.Println(g.Nombre + "*Miau*")
}

//Func to make noise with struct
func MakeNoise(animal Animal) {
	animal.sonido()
}
