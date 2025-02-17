package states

import (
	"ADPwn/cmd/internal/states/common"
	"ADPwn/database/model"
	"ADPwn/database/service"
	db_context "context"
	"github.com/rivo/tview"
	"log"
	"time"
)

type AddUserState struct {
	App     *tview.Application
	Project model.Project

	username string
	password string
	ntlmHash string
	isAdmin  bool
}

func (s *AddUserState) Execute(context *common.Context) {
	//inputField := tview.NewInputField().
	//	SetLabel("Add User").
	//	SetFieldWidth(20)
	form := tview.NewForm().
		AddInputField("Username", "", 20, nil, func(text string) {
			s.username = text
		}).
		AddInputField("Password", "", 20, nil, func(text string) {
			s.password = text
		}).
		AddInputField("NTLM Hash", "", 40, nil, func(text string) {
			s.ntlmHash = text
		}).
		AddCheckbox("Is Domainadmin", false, func(checked bool) {
			s.isAdmin = checked
		}).
		AddButton("Save", func() {
			s.addUser()
			common.ShowSuccessAlert(s.App, "User added successfully!", func() {
				context.SetState(&MainState{App: s.App, Project: s.Project})
			})

		}).
		AddButton("Back", func() {
			context.SetState(&MainState{App: s.App, Project: s.Project})
		})
	form.SetBorder(true).SetTitle("Add User").SetTitleAlign(tview.AlignLeft)

	s.App.SetRoot(form, true).SetFocus(form)
}

func (s *AddUserState) addUser() {
	ctx, cancel := db_context.WithTimeout(db_context.Background(), 5*time.Second)

	defer cancel()

	projectService, _ := service.NewProjectService()
	project, err := projectService.SaveUser(ctx, s.Project, s.username, s.password, s.ntlmHash, s.isAdmin)
	s.Project = project

	if err != nil {
		log.Fatal("error while saving new user to project: ", err)
	}
}
