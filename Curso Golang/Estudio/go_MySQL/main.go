package main

import (
	"fmt"
	"go_MySQL/database"
	"go_MySQL/handlers"
	"go_MySQL/models"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	for {
		fmt.Println("1. Listar contactos")
		fmt.Println("2. Buscar contacto por ID")
		fmt.Println("3. Crear contacto")
		fmt.Println("4. Actualizar contacto")
		fmt.Println("5. Eliminar contacto")
		fmt.Println("6. Salir")
		var option int
		fmt.Scanln(&option)
		switch option {
		case 1:
			handlers.GetContacts(db)
		case 2:
			fmt.Println("Ingrese el ID del contacto a buscar:")
			var contactID int
			fmt.Scanln(&contactID)
			handlers.GetContactByID(db, contactID)
		case 3:
			fmt.Println("Ingrese el nombre del contacto:")
			var name string
			fmt.Scanln(&name)
			fmt.Println("Ingrese el email del contacto:")
			var email string
			fmt.Scanln(&email)
			fmt.Println("Ingrese el teléfono del contacto:")
			var phone string
			fmt.Scanln(&phone)
			handlers.CreateContact(db, models.Contact{Name: name, Email: email, Phone: phone})
		case 4:
			fmt.Println("Ingrese el ID del contacto a actualizar:")
			var contactID int
			fmt.Scanln(&contactID)
			fmt.Println("Ingrese el nombre del contacto:")
			var name string
			fmt.Scanln(&name)
			fmt.Println("Ingrese el email del contacto:")
			var email string
			fmt.Scanln(&email)
			fmt.Println("Ingrese el teléfono del contacto:")
			var phone string
			fmt.Scanln(&phone)
			handlers.UpdateContact(db, models.Contact{Name: name, Email: email, Phone: phone}, contactID)
		case 5:
			fmt.Println("Ingrese el ID del contacto a eliminar:")
			var contactID int
			fmt.Scanln(&contactID)
			handlers.DeleteContact(db, contactID)
		case 6:
			return
		}

	}
}
