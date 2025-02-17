package model

type Domain struct {
	UID        string            `json:"uid,omitempty"`
	Name       string            `json:"name"`
	Hosts      []Host            `json:"hosts,omitempty"`
	Users      []User            `json:"users,omitempty"`
	ModuleRuns map[string]string `json:"metadata,omitempty"`
	DType      []string          `json:"dgraph.type,omitempty"`
}

func NewDomain(name string) *Domain {
	return &Domain{
		Name:  name,
		DType: []string{"domain"},
	}
}
