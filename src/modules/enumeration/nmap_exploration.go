package enumeration

import (
	"ADPwn/core/model"
	"ADPwn/modules"
)

type NmapExplorer struct {
	Name         string
	Description  string
	Version      string
	Author       string
	Dependencies []string
	Modes        []string
}

func (n *NmapExplorer) DependsOn() int {
	//TODO implement me
	panic("implement me")
}

func (n *NmapExplorer) GetName() string {
	return n.Name
}

func (n *NmapExplorer) GetDescription() string {
	return n.Description
}

func (n *NmapExplorer) GetVersion() string {
	return n.Version
}

func (n *NmapExplorer) GetAuthor() string {
	return n.Author
}

func (n *NmapExplorer) Execute(project model.Project, options []string) error {

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

	modules.GlobalRegistry.RegisterEnumerationModule(module)
}
