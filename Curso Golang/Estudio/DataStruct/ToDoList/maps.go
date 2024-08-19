package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}
	for clave, valor := range colors {
		fmt.Printf("Clave: %s, Valor: %s \n", clave, valor)
	}
}
