package attacks

import (
	plugin "ADPwn-core/internal/module"
	"ADPwn-core/internal/sse"
	"ADPwn-core/pkg/model/adpwn"
)

type PrinterNightmare struct {
	ConfigKey string
}

func (n *PrinterNightmare) GetConfigKey() string {
	//TODO implement me
	return n.ConfigKey
}

func (n *PrinterNightmare) DependsOn() int {
	//TODO implement me
	panic("implement me")
}

func (n *PrinterNightmare) ExecuteModule(params *adpwn.Parameter, logger *sse.SSELogger) error {
	return nil
}

// INIT
func init() {
	module := &PrinterNightmare{
		ConfigKey: "PrinterNightmare",
	}
	plugin.RegisterPlugin(module)
}
