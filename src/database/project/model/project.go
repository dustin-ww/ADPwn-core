package model

type Project struct {
	UID   string   `json:"uid,omitempty"`
	Name  string   `json:"name"`
	Hosts []Host   `json:"hosts,omitempty"`
	Users []User   `json:"users,omitempty"`
	DType []string `json:"dgraph.type,omitempty"`
}

func NewProject(name string) *Project {
	return &Project{
		Name:  name,
		DType: []string{"project"},
	}
}
