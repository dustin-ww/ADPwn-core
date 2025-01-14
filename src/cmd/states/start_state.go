package states

import (
	"ADPwn/cmd/states/common"
	"fmt"
	"os"

	"github.com/rivo/tview"
)

type StartState struct {
	App *tview.Application
}

func (s *StartState) Execute(context *common.Context) {
	title := tview.NewTextView().
		SetText("ADPwn - Start Menu").
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	list := tview.NewList().
		AddItem("1. Select Project", "Select an existing project to perform action", '1', func() {
			context.SetState(&ProjectSelectState{App: s.App})
		}).
		AddItem("2. Create new project", "Create a new project", '2', func() {
			context.SetState(&ProjectCreateState{App: s.App})
		}).
		AddItem("3. Exit", "Exit ADPwn", '3', func() {
			fmt.Println("Exiting...")
			s.App.Stop()
			context.SetState(nil)
			os.Exit(0)
		})

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(title, 3, 0, false).
		AddItem(list, 0, 1, true)

	if err := s.App.SetRoot(flex, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
}
