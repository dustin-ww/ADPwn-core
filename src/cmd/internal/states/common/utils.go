package common

import "github.com/rivo/tview"

func ShowSuccessAlert(app *tview.Application, message string, onClose func()) {
	buildAlert(app, message, false, onClose)
}

func ShowErrorAlert(app *tview.Application, message string, onClose func()) {
	buildAlert(app, message, true, onClose)
}

func buildAlert(app *tview.Application, message string, isError bool, onClose func()) {
	modal := tview.NewModal().
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if onClose != nil {
				onClose()
			}
		})

	if isError {
		modal.SetText("[red]" + message + "[white]")
	} else {
		modal.SetText("[green]" + message + "[white]")
	}

	app.SetRoot(modal, true).SetFocus(modal)
}
