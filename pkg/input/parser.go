package input

import (
	"ADPwn-core/pkg/model"
	"ADPwn-core/pkg/model/adpwn/input"
	"encoding/json"
	"fmt"
)

type Parser struct{}

func UnmarshalInput(raw json.RawMessage, inputType string) (input.InputValue, error) {
	switch inputType {
	case "checkbox":
		var c input.CheckboxValue
		if err := json.Unmarshal(raw, &c); err != nil {
			return nil, err
		}
		return c, nil

	case "targetInput":
		var tmp struct {
			input.CommonFields
			Value json.RawMessage `json:"value"`
		}
		if err := json.Unmarshal(raw, &tmp); err != nil {
			return nil, err
		}
		var list []model.Target
		if err := json.Unmarshal(tmp.Value, &list); err != nil {
			return nil, err
		}
		return input.TargetListValue{tmp.CommonFields, list}, nil

	default:
		var t input.TextInputValue
		if err := json.Unmarshal(raw, &t); err != nil {
			return nil, err
		}
		return t, nil
	}
}

func ParseParameters(jsonData []byte) (input.Parameter, error) {
	var raw struct {
		RunID    string                     `json:"runId"`
		Metadata map[string]string          `json:"metadata"`
		Inputs   map[string]json.RawMessage `json:"inputs"`
	}
	if err := json.Unmarshal(jsonData, &raw); err != nil {
		return input.Parameter{}, err
	}

	result := input.Parameter{
		RunID:    raw.RunID,
		Metadata: raw.Metadata,
		Inputs:   make(map[string]input.InputValue),
	}

	for key, rawMsg := range raw.Inputs {
		var typeHolder struct {
			Type string `json:"type"`
		}
		if err := json.Unmarshal(rawMsg, &typeHolder); err != nil {
			typeHolder.Type = "textInput"
		}
		inputType := typeHolder.Type
		if inputType == "" {
			inputType = "textInput"
		}

		val, err := UnmarshalInput(rawMsg, inputType)
		if err != nil {
			return result, fmt.Errorf("error while parsing '%s': %w", key, err)
		}
		result.Inputs[key] = val
	}

	return result, nil
}

/*jsonData := []byte(`{
	"runId": "scan-123",
	"inputs": {
		"fullscan": { "type": "checkbox", "value": true },
		"udp": { "type": "textInput", "value": "yes" },
		"additional": { "type": "textInput", "value": "-sV -O" },
		"targetInput": { "type": "targetInput", "value": [
			{"uid":"0x1","name":"Server 1","ip_range":"192.168.1.1-10","dgraph.type":["Target"]},
			{"uid":"0x2","name":"Server 2","ip_range":"10.0.0.5","dgraph.type":["Target"]}
		] }
	},
	"metadata": {"creator":"admin","priority":"high"}
}`)*/
