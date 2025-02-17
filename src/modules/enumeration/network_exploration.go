package enumeration

import (
	"ADPwn/database/model"
	"ADPwn/modules/internal/base"
	"log"
)

type NetworkExplorer struct {
	Name         string
	Description  string
	Version      string
	Author       string
	Dependencies []string
	Modes        []string
}

// --- Interface Methods
func (n NetworkExplorer) GetName() string {
	return n.Name
}

func (n NetworkExplorer) GetDescription() string {
	return n.Description
}

func (n NetworkExplorer) GetVersion() string {
	return n.Version
}

func (n NetworkExplorer) GetAuthor() string {
	return n.Author
}

func (n NetworkExplorer) Execute(project model.Project, options []string) error {
	log.Println("Starting AD network enumeration")
	return nil
}

// INIT
func init() {
	module := NetworkExplorer{
		Name:        "NetworkExploration",
		Description: "ADPwn Module to enumerate ad network",
		Version:     "0.1",
		Author:      "Dustin Wickert",
	}

	base.GlobalRegistry.RegisterModule(module)
}
