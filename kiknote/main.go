package main

import (
	"fmt"
	"log"
	"os"

	tea "charm.land/bubbletea/v2"
)

func main() {
	store := new(DataStore)
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
