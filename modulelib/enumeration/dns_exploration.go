package enumeration

import (
	plugin "ADPwn-core/internal/module"
	"ADPwn-core/internal/sse"
	"ADPwn-core/pkg/model/adpwn"
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

func (n *DNSExplorer) ExecuteModule(params *adpwn.Parameter, logger *sse.SSELogger) error {
	return nil
}

// INIT
func init() {
	module := &DNSExplorer{
		ConfigKey: "DNSExplorer",
	}
	plugin.RegisterPlugin(module)

}
