package main

import (
	"log"

	"charm.land/bubbles/v2/textarea"
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

var placeholder tea.Cmd

const (
	listView uint = iota
	titleView
	bodyView
)

type model struct {
	state     uint
	store     *DataStore
	notes     []Note
	note      Note
	cursor    int
	textarea  textarea.Model
	textinput textinput.Model
}

func NewModel(datastore *DataStore) model {
	notes, err := datastore.GetNotes()
	if err != nil {
		log.Fatalf("Unable to get Notes: %v", err)
	}
	return model{
		state:     listView,
		store:     datastore,
		notes:     notes,
		textarea:  textarea.New(),
		textinput: textinput.New(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		err  error
		cmd  tea.Cmd
		cmds []tea.Cmd
	)
	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)
	m.textinput, cmd = m.textinput.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		key := msg.String()
		switch m.state {
		case listView:
			switch key {
			case "q", "ctrl+c":
				return m, tea.Quit

			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down", "j":
				if m.cursor < len(m.notes)-1 {
					m.cursor++
				}
			case "n":
				m.textinput.SetValue("")
				m.textinput.Focus()
				m.note = Note{}
				m.state = titleView
			case "enter":
				m.note = m.notes[m.cursor]
				m.state = bodyView
				m.textarea.SetValue(m.note.Body)
				m.textarea.Focus()
				m.textarea.CursorEnd()
			}
		case titleView:
			switch key {
			case "enter":
				title := m.textinput.Value()
				if title != "" {
					m.note.Title = title
					m.state = bodyView
					m.textarea.SetValue("")
					m.textarea.Focus()
					m.textarea.CursorEnd()
				}
			case "esc":
				m.state = listView
			}
		case bodyView:
			switch key {
			case "ctrl+s":
				m.note.Body = m.textinput.Value()
				if err = m.store.SaveNotes(m.note); err != nil {
					return m, tea.Quit
				}
				m.notes, err = m.store.GetNotes()
				if err != nil {
					return m, tea.Quit
				}
				m.state = listView
			case "esc":
				m.state = listView
			}

		}
	}

	return m, tea.Batch(cmds...)
}
