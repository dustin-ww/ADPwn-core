package enumeration

import (
	plugin "ADPwn-core/internal/module"
	"ADPwn-core/internal/sse"
	"ADPwn-core/pkg/model/adpwn/input"
)

type DNSExplorer struct {
	dependencies []string
	Modes        []string
	configKey    string
}

func (n *DNSExplorer) ConfigKey() string {
	return n.configKey
}

func (n *DNSExplorer) DependsOn() int {
	//TODO implement me
	panic("implement me")
}

func (n *DNSExplorer) Dependencies() []string {
	return n.dependencies
}

func (n *DNSExplorer) ExecuteModule(params *input.Parameter, logger *sse.SSELogger) error {
	return nil
}

// INIT
func init() {
	module := &DNSExplorer{
		configKey: "DNSExplorer",
	}
	plugin.RegisterPlugin(module)

}
