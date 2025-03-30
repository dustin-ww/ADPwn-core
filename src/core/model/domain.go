package model

type Domain struct {
	UID              string   `json:"uid,omitempty"`
	Name             string   `json:"name,omitempty"`
	BelongsToProject Project  `json:"belongs_to_project,omitempty"`
	HasHost          []Host   `json:"has_host,omitempty"`
	HasUser          []User   `json:"has_user,omitempty"`
	DType            []string `json:"dgraph.type,omitempty"`
}
