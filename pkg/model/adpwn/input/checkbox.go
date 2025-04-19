package input

type CheckboxValue struct {
	CommonFields
	Value bool `json:"value"`
}

func (CheckboxValue) typeName() string { return "checkbox" }
