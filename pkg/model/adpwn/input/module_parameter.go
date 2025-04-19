package input

import "ADPwn-core/pkg/model"

type InputValue interface {
	typeName() string
}
type Parameter struct {
	RunID    string                `json:"runId"`
	Inputs   map[string]InputValue `json:"inputs"`
	Metadata map[string]string     `json:"metadata"`
}

func (p *Parameter) GetTextInput(key string) *string {
	if iv, ok := p.Inputs[key]; ok {
		if ti, ok := iv.(TextInputValue); ok {
			return &ti.Value
		}
	}
	return nil
}

func (p *Parameter) GetCheckbox(key string) *bool {
	if iv, ok := p.Inputs[key]; ok {
		if cb, ok := iv.(CheckboxValue); ok {
			return &cb.Value
		}
	}
	return nil
}

func (p *Parameter) GetTargetInput(key string) *[]model.Target {
	if iv, ok := p.Inputs[key]; ok {
		if tl, ok := iv.(TargetListValue); ok {
			return &tl.Value
		}
	}
	return nil
}
