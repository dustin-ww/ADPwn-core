package states

import "fmt"

// SubMenuState repräsentiert das Untermenü
type SubMenuState struct{}

func (s *SubMenuState) Execute(context *Context) {
	fmt.Println("\nUntermenü:")
	fmt.Println("1. Zurück zum Hauptmenü")
	fmt.Println("2. Aktion ausführen")
	var choice int
	fmt.Print("Wählen Sie: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		context.SetState(&MainMenuState{})
	case 2:
		fmt.Println("Aktion ausgeführt.")
	default:
		fmt.Println("Ungültige Auswahl.")
	}
}
