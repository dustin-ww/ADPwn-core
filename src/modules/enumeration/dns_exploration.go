package enumeration

import (
	"ADPwn/cmd/logger"
	"ADPwn/core/model"
	"ADPwn/modules/internal/base"
	"fmt"
)

type DNSExplorer struct {
	Name         string
	Description  string
	Version      string
	Author       string
	Dependencies []string
	Modes        []string
	Logger       *logger.ADPwnLogger
}

func (n *DNSExplorer) DependsOn() int {
	//TODO implement me
	panic("implement me")
}

func (n *DNSExplorer) GetName() string {
	return n.Name
}

func (n *DNSExplorer) GetDescription() string {
	return n.Description
}

func (n *DNSExplorer) GetVersion() string {
	return n.Version
}

func (n *DNSExplorer) GetAuthor() string {
	return n.Author
}

func (n *DNSExplorer) Execute(project model.Project, options []string) error {

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
