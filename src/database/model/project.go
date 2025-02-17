package model

type Project struct {
	UID     string   `json:"uid,omitempty"`
	Name    string   `json:"name"`
	Domains []Domain `json:"domains"`
	Targets []string `json:"targets"`
	DType   []string `json:"dgraph.type,omitempty"`
}

func NewProject(name string) *Project {

	// todo: remove default domain
	domain := NewDomain("main")

	return &Project{
		Name:    name,
		Domains: []Domain{*domain},
		DType:   []string{"project"},
	}
}
