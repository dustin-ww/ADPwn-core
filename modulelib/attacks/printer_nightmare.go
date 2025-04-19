package attacks

import (
	plugin "ADPwn-core/internal/module"
	"ADPwn-core/internal/sse"
	"ADPwn-core/pkg/model/adpwn/input"
)

type PrinterNightmare struct {
	configKey string
}

func (n *PrinterNightmare) ConfigKey() string {
	//TODO implement me
	return n.configKey
}

func (n *PrinterNightmare) DependsOn() int {
	//TODO implement me
	panic("implement me")
}

func (n *PrinterNightmare) ExecuteModule(params *input.Parameter, logger *sse.SSELogger) error {
	return nil
}

// INIT
func init() {
	module := &PrinterNightmare{
		configKey: "PrinterNightmare",
	}
	plugin.RegisterPlugin(module)
}
