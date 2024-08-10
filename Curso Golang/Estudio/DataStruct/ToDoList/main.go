package main

import "fmt"

func main() {
	var array [10]int
	array[0] = 1
	fmt.Println(array)
	var matrix = [2][2]int{{1, 1}, {2, 2}}
	fmt.Println(matrix)
	diasSemana := []string{"domingo", "lunes", "martes", "miercoles", "jueves", "viernes", "sabado"}
	fmt.Println(diasSemana)
	diaRebanada := diasSemana[0:4]
	fmt.Println(diaRebanada)

	diaRebanada = append(diaRebanada, "festivo", "dia patrio", "dia extra", "ampliacion array")
	fmt.Println(diaRebanada)
	diaRebanada = append(diaRebanada[:2], diaRebanada[3:]...)
	fmt.Println(diaRebanada)
	fmt.Println(len(diaRebanada))
	fmt.Println(cap(diaRebanada))
}
