package model

type Project struct {
	UID     string   `json:"uid,omitempty"`
	Name    string   `json:"name"`
	Domains []Domain `json:"domains"`
	DType   []string `json:"dgraph.type,omitempty"`
}

func NewProject(name string) *Project {

	domain := NewDomain("main")

	return &Project{
		Name:    name,
		Domains: []Domain{*domain},
		DType:   []string{"project"},
	}
}
