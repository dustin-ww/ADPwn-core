package states

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// Die States
type StartMenuState struct{}

type ProjectSelectMenuState struct{}
type ProjectCreateMenuState struct{}

// Das Context-Objekt, das den aktuellen Zustand speichert
type Context struct {
	CurrentState tea.Model
}

func (c *Context) SetState(state tea.Model) {
	c.CurrentState = state
}

func (c *Context) Execute() {
	if c.CurrentState != nil {
		// Update und View Methoden von Bubble Tea ausführen
		c.CurrentState.Update(tea.KeyMsg{Type: tea.KeyEnter}) // Test für Enter
		fmt.Println(c.CurrentState.View())                    // View anzeigen
	}
}

// Das Startmenü als Bubble Tea-Modell
type model struct {
	list    list.Model
	context *Context
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			// Auswahl treffen
			i, ok := m.list.SelectedItem().(item)
			if ok {
				fmt.Println("You selected:", i)
				// Je nach Auswahl den entsprechenden State setzen
				switch i {
				case "Select Project to load":
					m.context.SetState(&ProjectSelectMenuState{})
				case "Create new project":
					m.context.SetState(&ProjectCreateMenuState{})
				case "Exit":
					m.context.SetState(nil)
				}
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return "\n" + m.list.View()
}

// item ist das Listenelement
type item string

func (i item) FilterValue() string { return string(i) }

func main() {
	items := []list.Item{
		item("Select Project to load"),
		item("Create new project"),
		item("Exit"),
	}

	// Liste initialisieren
	l := list.New(items, list.NewDefaultDelegate(), 0, 10)
	l.Title = "ADPwn - Menu"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	// Context und Startmenü State setzen
	context := &Context{}
	context.SetState(&model{list: l, context: context})

	// Start Bubble Tea Anwendung
	if _, err := tea.NewProgram(context.CurrentState).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
