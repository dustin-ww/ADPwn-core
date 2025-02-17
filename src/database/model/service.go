package model

type Service struct {
	UID   string   `json:"uid,omitempty"`
	Name  string   `json:"name"`
	Port  string   `json:"port"`
	DType []string `json:"dgraph.type,omitempty"`
}

func NewService(name string, port string) *Service {

	return &Service{
		Name:  name,
		Port:  port,
		DType: []string{"service"},
	}
}
