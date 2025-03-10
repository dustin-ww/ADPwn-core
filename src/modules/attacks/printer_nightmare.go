package attacks

import (
	"ADPwn/core/model"
	"ADPwn/core/plugin"
)

type PrinterNightmare struct {
	Name            string
	Description     string
	Version         string
	Author          string
	Dependencies    []string
	Modes           []string
	ExecutionMetric string
}

func (n *PrinterNightmare) GetExecutionMetric() string {
	return n.ExecutionMetric
}

func (n *PrinterNightmare) DependsOn() int {
	//TODO implement me
	panic("implement me")
}

func (n *PrinterNightmare) GetName() string {
	return n.Name
}

func (n *PrinterNightmare) GetDescription() string {
	return n.Description
}

func (n *PrinterNightmare) GetVersion() string {
	return n.Version
}

func (n *PrinterNightmare) GetAuthor() string {
	return n.Author
}

func (n *PrinterNightmare) Execute(project model.Project, options []string) error {

	return nil
}

// INIT
func init() {
	module := &PrinterNightmare{
		Name:            "Printer Nightmare Attack",
		Description:     "ADPwn Module to exploit printer nightmare",
		Version:         "0.1",
		Author:          "Dustin",
		ExecutionMetric: "1h",
	}
	plugin.RegisterAttack(module)
}
