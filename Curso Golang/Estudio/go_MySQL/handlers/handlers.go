package handlers

import (
	"database/sql"
	"fmt"
	"go_MySQL/models"
	"log"
)

func GetContacts(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM contacts")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	fmt.Println("\n LISTA DE CONTACTOS:")
	fmt.Println("----------------------")
	for rows.Next() {
		//Instance contact model
		contact := models.Contact{}
		//Scan row into contact model
		err := rows.Scan(&contact.ID, &contact.Name, &contact.Email, &contact.Phone)
		if err != nil {
			log.Fatal(err)
		}
		//Print contact
		fmt.Printf("ID: %d, Nombre: %s, Email: %s, Teléfono: %s\n", contact.ID, contact.Name, contact.Email, contact.Phone)
		fmt.Println("----------------------")
	}
}
func GetContactByID(db *sql.DB, contactID int) {
	row := db.QueryRow("SELECT * from contacts where id = ?", contactID)

	contact := models.Contact{}
	err := row.Scan(&contact.ID, &contact.Name, &contact.Email, &contact.Phone)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID: %d, Nombre: %s, Email: %s, Teléfono: %s\n", contact.ID, contact.Name, contact.Email, contact.Phone)
	fmt.Println("----------------------")
}
func CreateContact(db *sql.DB, contact models.Contact) {
	query := "insert into contacts (name, email, phone) values (?, ?, ?)"

	db.Exec(query, contact.Name, contact.Email, contact.Phone)
	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Contacto creado correctamente")
}
func DeleteContact(db *sql.DB, contactID int) {
	query := "delete from contacts where id = ?"
	_, err := db.Exec(query, contactID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Contacto eliminado correctamente")
}
func UpdateContact(db *sql.DB, contact models.Contact, contactID int) {
	query := "update contacts set name = ?, email = ?, phone = ? where id = ?"
	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone, contactID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Contacto actualizado correctamente")
}
