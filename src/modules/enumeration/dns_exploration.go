package enumeration

import (
	"ADPwn/core/model/adpwn"
	"ADPwn/core/plugin"
)

type DNSExplorer struct {
	Dependencies []string
	Modes        []string
	ConfigKey    string
}

func (n *DNSExplorer) GetConfigKey() string {
	return n.ConfigKey
}

func (n *DNSExplorer) DependsOn() int {
	//TODO implement me
	panic("implement me")
}

func (n *DNSExplorer) GetDependencies() []string {
	return n.Dependencies
}

func (n *DNSExplorer) Execute(parameter adpwn.Parameter) error {

	return nil
}

// INIT
func init() {
	module := &DNSExplorer{
		ConfigKey: "DNSExplorer",
	}
	plugin.RegisterPlugin(module)

}
