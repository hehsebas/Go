package module

import (
	"errors"
	"fmt"
	"math/rand"
)

// Module fnc hello
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty Name")
	}
	message := fmt.Sprintf(AleatoriesName(), name)
	return message, nil
}
func AleatoriesName() string {
	AleatoriesName := []string{
		"Hello, %v Welcome",
		"Hola, %v Bienvenido",
		"Salute, %v ",
		"Bonjour, %v ",
	}
	return AleatoriesName[rand.Intn(len(AleatoriesName))]
}
