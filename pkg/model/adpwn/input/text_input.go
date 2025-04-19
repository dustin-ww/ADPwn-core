package input

type TextInputValue struct {
	CommonFields
	Value string `json:"value"`
}

func (TextInputValue) typeName() string { return "textInput" }
