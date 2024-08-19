package main

import (
	"errors"
	"fmt"
)

func divide(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("No se puede dividir por cero.")
	}
	return dividendo / divisor, nil
}

func main() {
	resultado, err := divide(4, 2)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Resultado: ", resultado)
}
