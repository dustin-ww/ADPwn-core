package input

import "ADPwn-core/pkg/model"

type TargetListValue struct {
	CommonFields
	Value []model.Target `json:"value"`
}

func (TargetListValue) typeName() string { return "targetInput" }
