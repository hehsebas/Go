package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Contact Structure
type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// Save contacts in json file
func saveContactToFile(contact []Contact) error {
	file, err := os.Create("contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(contact)
	if err != nil {
		return err
	}
	return nil
}

// Load contacts from json file
func loadContactFromFile(contact *[]Contact) error {
	file, err := os.Open("contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(contact)
	if err != nil {
		return err
	}
	return nil
}
func main() {
	//Contact Slice
	var contacts []Contact
	//Load Contact from json file
	err := loadContactFromFile(&contacts)
	if err != nil {
		fmt.Println("Error loading contact", err)
	}
	//Bufio instance
	reader := bufio.NewReader(os.Stdin)
	for {
		//Show user options
		fmt.Println("User Options")
		fmt.Println("-----------------------------------------------")
		fmt.Println("1. Create contact")
		fmt.Println("2. Show all contacts")
		fmt.Println("3. Exit")
		fmt.Println("-----------------------------------------------")

		var option int
		_, err = fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Reading option failed: ", err)
			return
		}

		switch option {
		// Create contact
		case 1:
			var c Contact
			fmt.Println("Type contact name: ")
			c.Name, _ = reader.ReadString('\n')
			fmt.Println("Type contact email: ")
			c.Email, _ = reader.ReadString('\n')
			fmt.Println("Type contact phone: ")
			c.Phone, _ = reader.ReadString('\n')
			//Add contact to slice.
			contacts = append(contacts, c)
			// Save json file
			if err := saveContactToFile(contacts); err != nil {
				fmt.Println("Error saving contact: ", err)
			}
			//Show all contacts
		case 2:
			fmt.Println("================================================================")
			for index, contact := range contacts {
				fmt.Printf("%d.\n Nombre: %s Email: %s Telefono: %s\n", index+1, contact.Name, contact.Email, contact.Phone)
			}
			fmt.Println("================================================================")
		case 3:
			return
		default:

			fmt.Println("\n\n\n")
			fmt.Println("Invalid Option\n\n\n")
			main()
		}
	}
}
