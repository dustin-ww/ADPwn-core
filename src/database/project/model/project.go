package model

type Project struct {
	UID   string `json:"uid"`
	ID    string `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Hosts []Host `json:"nodes,omitempty"`
}

func NewProject(ID string, name string) *Project {
	return &Project{
		ID:   ID,
		Name: name,
		Type: "Project",
	}
}
