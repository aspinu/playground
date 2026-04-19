package main

import (
	"fmt"
	"log"
	"os"

	tea "charm.land/bubbletea/v2"
)

func main() {
	// home, _ := os.UserHomeDir()
	// filename := filepath.Join(home, "Repositories", "playground", "kiknote", "errors")
	// file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	store := &DataStore{}
	if err := store.Init(); err != nil {
		log.Fatalf("unable to init store: %v", err)
	}

	m := NewModel(store)
	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
