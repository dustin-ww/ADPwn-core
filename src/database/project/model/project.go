package model

type Project struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"project,omitempty"`
}

func NewProject(ID int, name string) *Project {
	return &Project{
		ID:   ID,
		Name: name,
		Type: "Project",
	}
}
