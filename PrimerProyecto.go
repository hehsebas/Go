package main

import (
	"fmt"
	"math/rand"
)

func main() {
	jugar()
}
func jugar() {
	num := rand.Intn(100)
	var num2 int
	var intentos int
	const maxintent = 10
	for intentos < maxintent {
		intentos++
		fmt.Printf("Ingrese un número (numero de intentos restantes %d) :\n", (maxintent-intentos)+1)
		fmt.Scanln(&num2)
		if num2 == num {
			fmt.Println("Felicidades has ganado")
			jugardenuevo()
		} else if num2 < num {
			fmt.Println("El numero ingresado es menor")
		} else if num2 > num {
			fmt.Println("El numero ingresado es mayor")
		}
	}
	fmt.Println("Lo siento, el numero era: ", num)
	jugardenuevo()
}
func jugardenuevo() {
	var Eleccion string
	fmt.Println("Quieres jugar nuevamente? (s/n)")
	fmt.Scanln(&Eleccion)
	switch Eleccion {
	case "s":
		jugar()
	case "n":
		fmt.Println("Gracias por jugar")
	default:
		fmt.Println("Opcion ingresada no válida, por favor intenta nuevamente")
		jugardenuevo()
	}

}
