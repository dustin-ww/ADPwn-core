package model

type Project struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func NewProject(ID string, name string) *Project {
	return &Project{
		ID:   ID,
		Name: name,
		Type: "Project",
	}
}
