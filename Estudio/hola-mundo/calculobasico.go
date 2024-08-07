package main

import (
	"fmt"
	"math"
)

const (
// precision = 2
)

func main() {
	var lado1 float64
	var lado2 float64

	fmt.Printf("Bienvenido, por favor ingrese el valor de uno de los lados del triangulo \n")
	fmt.Scan(&lado1)
	fmt.Println("Ingrese el valor de lado 2")
	fmt.Scan(&lado2)

	var hipotenusa float64 = math.Sqrt(math.Pow(lado1, 2) + math.Pow(lado2, 2))
	fmt.Printf("El valor del area del triangulo es de: %.*f \n", precision, (lado1 * lado2 / 2))
	fmt.Printf("El valor de la hipotenusa del triangulo es de: %.*f \n", precision, hipotenusa)
	fmt.Printf("El valor del perimetro del triangulo es de: %.*f\n", precision, (lado1 + lado2 + hipotenusa))

}
