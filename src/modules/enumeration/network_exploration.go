package enumeration

import (
	"ADPwn/cmd/logger"
	"ADPwn/database/model"
	"ADPwn/modules/internal/base"
	"fmt"
)

type NetworkExplorer struct {
	Name         string
	Description  string
	Version      string
	Author       string
	Dependencies []string
	Modes        []string
	Logger       *logger.ADPwnLogger
}

func (n *NetworkExplorer) DependsOn() int {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkExplorer) GetName() string {
	return n.Name
}

func (n *NetworkExplorer) GetDescription() string {
	return n.Description
}

func (n *NetworkExplorer) GetVersion() string {
	return n.Version
}

func (n *NetworkExplorer) GetAuthor() string {
	return n.Author
}

func (n *NetworkExplorer) Execute(project model.Project, options []string) error {
	if n.Logger == nil {
		n.Logger = logger.NewADPwnLogger()
	}

	n.Logger.Log("[*] Starting AD network enumeration")
	n.Logger.Log(fmt.Sprintf("[*] Scanning project: %s", project.Name))
	n.Logger.Log(fmt.Sprintf("[*] Options: %v", options))

	return nil
}

// INIT
func init() {
	module := &NetworkExplorer{
		Name:        "NetworkExploration",
		Description: "ADPwn Module to enumerate ad network",
		Version:     "0.1",
		Author:      "Dustin Wickert",
	}

	base.GlobalRegistry.RegisterEnumerationModule(module)
}
