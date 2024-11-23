package states

import "fmt"

type MainMenuState struct{}

func (s *MainMenuState) Execute(context *Context) {
	fmt.Println("\nHauptmenü:")
	fmt.Println("1. Zum Untermenü")
	fmt.Println("2. Beenden")
	var choice int
	fmt.Print("Wählen Sie: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		context.SetState(&SubMenuState{})
	case 2:
		fmt.Println("Programm wird beendet.")
		context.SetState(nil)
	default:
		fmt.Println("Ungültige Auswahl.")
	}
}
