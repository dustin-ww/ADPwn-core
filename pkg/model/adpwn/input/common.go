package input

type CommonFields struct {
	Key         string `json:"key,omitempty"`
	Label       string `json:"label,omitempty"`
	Placeholder string `json:"placeholder,omitempty"`
	Required    bool   `json:"required,omitempty"`
}
