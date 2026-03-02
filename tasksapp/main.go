package main

import (
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
)

type taskFileds struct {
	taskName    string
	taskDetails string
	taskStatus  string
}

func main() {
	var (
		task         string
		confirmation bool
	)

	// Create a new form
	form := huh.NewForm(
		huh.NewGroup(
			// Text Input
			// huh.NewInput().
			// 	Title("What is your name?").
			// 	Value(&name).
			// 	Placeholder("Pizza Enthusiast"),

			// Select Menu
			huh.NewSelect[string]().
				Title("Start").
				Options(
					huh.NewOption("CreateTask", "func call Create"),
					huh.NewOption("ReadTasks", "func call Read"),
					huh.NewOption("UpdateTask", "func call Update"),
					huh.NewOption("DeleteTask", "func call Delete"),
				).
				Value(&task),

			// Confirm Toggle
			huh.NewConfirm().
				Title("Confirm").
				Value(&confirmation),
		),
	)

	// Run the form
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	// Access the results via your pointers
	fmt.Printf("A %s (Spicy: %v)\n", task, confirmation)
}
