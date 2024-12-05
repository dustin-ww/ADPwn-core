package states

import "fmt"

type StartMenuState struct{}

func (s *StartMenuState) Execute(context *Context) {
	fmt.Println("\nADPwn - Menu:")
	fmt.Println("1. Select Project to load")
	fmt.Println("2. Create new project")
	fmt.Println("3. Exit")
	var choice int
	fmt.Print("Please choose options: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		context.SetState(&ProjectSelectMenuState{})
	case 2:
		context.SetState(&ProjectCreateMenuState{})
	case 3:
		fmt.Println("Exit...")
		context.SetState(nil)
	default:
		fmt.Println("Invalid option.")
	}
}
