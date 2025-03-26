package techpalace

import (
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	customer = strings.ToUpper(customer) // Convertimos a mayúsculas
	return "Welcome to the Tech Palace, " + customer
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	border := strings.Repeat("*", numStarsPerLine) // Creamos la línea de bordes
	return border + "\n" + welcomeMsg + "\n" + border
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	oldMsg = strings.ReplaceAll(oldMsg, "*", "") // Eliminamos todos los asteriscos
	return strings.TrimSpace(oldMsg)             // Eliminamos espacios adicionales
}
