package model

import "errors"

type Service struct {
	UID   string   `json:"uid,omitempty"`
	Name  string   `json:"name"`
	Port  string   `json:"port"`
	DType []string `json:"dgraph.type,omitempty"`
}

type ServiceBuilder struct {
	service *Service
}

func NewServiceBuilder() *ServiceBuilder {
	return &ServiceBuilder{
		service: &Service{
			DType: []string{"service"},
		},
	}
}

func (b *ServiceBuilder) WithName(name string) *ServiceBuilder {
	b.service.Name = name
	return b
}

func (b *ServiceBuilder) WithPort(port string) *ServiceBuilder {
	b.service.Port = port
	return b
}

func (b *ServiceBuilder) Build() (*Service, error) {
	if b.service.Name == "" {
		return nil, errors.New("name is a required field")
	}
	if b.service.Port == "" {
		return nil, errors.New("port is a required field")
	}
	return b.service, nil
}
