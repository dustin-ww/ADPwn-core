package enumeration

import (
	"ADPwn/core/model"
	"ADPwn/core/plugin"
)

type DNSExplorer struct {
	Name         string
	Description  string
	Version      string
	Author       string
	Dependencies []string
	Modes        []string
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

	return nil
}

// INIT
func init() {
	module := &NetworkExplorer{
		Name:        "DNS Exploration",
		Description: "DNS Exploration",
		Version:     "0.1",
		Author:      "dw-sec",
	}

	plugin.RegisterEnumeration(module)
}
