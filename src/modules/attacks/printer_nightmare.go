package attacks

import (
	"ADPwn/core/model"
	"ADPwn/core/plugin"
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

func (n *PrinterNightmare) Execute(project model.Project, options []string) error {

	return nil
}

// INIT
func init() {
	module := &PrinterNightmare{
		ConfigKey: "PrinterNightmare",
	}
	plugin.RegisterPlugin(module)
}
