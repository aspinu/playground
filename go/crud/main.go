package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dixonwille/wmenu"
	//"github.com/dixonwille/wmenu/v5"
)

func main() {
	db, err := sql.Open("sqlite3", "./names.db")
	checkErr(err)

	defer db.Close()

	menu := wmenu.NewMenu("What would you like to do?")

	menu.Action(func(opts []wmenu.Opt) error { handleFunc(db, opts); return nil })

	menu.Option("Add a new Person", 0, true, nil)
	menu.Option("Find a Person", 1, false, nil)
	menu.Option("Update a Person information", 2, false, nil)
	menu.Option("Delete a person by ID", 3, false, nil)
	menu.Option("Quit Application", 4, false, nil)
	menuerr := menu.Run()

	if menuerr != nil {
		log.Fatal(menuerr)
	}
}

func handleFunc(db *sql.DB, opts []wmenu.Opt) {
	switch opts[0].Value {

	case 0:
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter a first name: ")
		firstName, _ := reader.ReadString('\n')
		if firstName != "\n" {
			firstName = strings.TrimSuffix(firstName, "\n")
		}

		fmt.Println("Enter a last name: ")
		lastName, _ := reader.ReadString('\n')
		if lastName != "\n" {
			lastName = strings.TrimSuffix(lastName, "\n")
		}

		fmt.Println("Enter an email address: ")
		email, _ := reader.ReadString('\n')

		if email != "\n" {
			email = strings.TrimSuffix(email, "\n")
		}

		fmt.Println("Enter an IP address: ")
		ipAddress, _ := reader.ReadString('\n')
		if ipAddress != "\n" {
			ipAddress = strings.TrimSuffix(ipAddress, "\n")
		}

		newPerson := person{
			first_name: firstName,
			last_name:  lastName,
			email:      email,
			ip_address: ipAddress,
		}

		addPerson(db, newPerson)

		break

	case 1:
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter a name to search for: ")
		searchString, _ := reader.ReadString('\n')

		searchString = strings.TrimSuffix(searchString, "\n")

		people := searchForPerson(db, searchString)
		fmt.Printf("Found %v results\n", len(people))

		for _, ourPerson := range people {
			fmt.Printf("\n----\nID: %v\nFirst Name: %s\nLast Name: %s\nEmail: %s\nIP Address: %s\n----\n", ourPerson.id, ourPerson.first_name, ourPerson.last_name, ourPerson.email, ourPerson.ip_address)
		}

		break

	case 2:

		reader := bufio.NewReader(os.Stdin)

		fmt.Println("Enter the person id to Enter the person id to update: ")

		updatedId, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		curnetPerson := getPersonById(db, updatedId)

		fmt.Printf("First name (Currently %s): ", curnetPerson.first_name)
		firstName, _ := reader.ReadString('\n')
		if firstName != "\n" {
			curnetPerson.first_name = strings.TrimSuffix(firstName, "\n")
		}

		fmt.Printf("Last Name (Currently %s): ", curnetPerson.last_name)
		lastName, _ := reader.ReadString('\n')
		if lastName != "\n" {
			curnetPerson.last_name = strings.TrimSuffix(lastName, "\n")
		}

		fmt.Printf("Email (Currently %s): ", curnetPerson.email)
		email, _ := reader.ReadString('\n')
		if email != "\n" {
			curnetPerson.email = strings.TrimSuffix(email, "\n")
		}

		fmt.Printf("IP Address (Currently %s): ", curnetPerson.ip_address)
		ipAddress, _ := reader.ReadString('\n')
		if ipAddress != "\n" {
			curnetPerson.ip_address = strings.TrimSuffix(ipAddress, "\n")
		}

		affected := updatePerson(db, curnetPerson)

		if affected == 1 {
			fmt.Println("One row affected")
		}
		break

	case 3:
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter the ID for the record you want to delete: ")
		searchString, _ := reader.ReadString('\n')

		idToDelete := strings.TrimSuffix(searchString, "\n")

		affected := deletePerson(db, idToDelete)

		if affected == 1 {
			fmt.Println("Deleted from database")
		}
		break

	case 4:
		fmt.Println("Goodbye")
		os.Exit(3)

	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
