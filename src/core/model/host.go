package model

import "errors"

type Host struct {
	UID                string    `json:"uid,omitempty"`
	IP                 string    `json:"ip"`
	Name               string    `json:"name"`
	HostProjectID      string    `json:"hostProjectID,omitempty"`
	IsDomaincontroller bool      `json:"isDomaincontroller"`
	Services           []Service `json:"has_service,omitempty"`
	DType              []string  `json:"dgraph.type,omitempty"`
}

type HostBuilder struct {
	host *Host
}

func NewHostBuilder() *HostBuilder {
	return &HostBuilder{
		host: &Host{
			DType: []string{"host"},
		},
	}
}

func (b *HostBuilder) WithIP(ip string) *HostBuilder {
	b.host.Name = ip
	b.host.IP = ip
	return b
}

func (b *HostBuilder) AsDomainController() *HostBuilder {
	b.host.IsDomaincontroller = true
	return b
}

func (b *HostBuilder) WithServices(services []Service) *HostBuilder {
	b.host.Services = services
	return b
}

func (b *HostBuilder) AddService(service Service) *HostBuilder {
	if b.host.Services == nil {
		b.host.Services = []Service{}
	}
	b.host.Services = append(b.host.Services, service)
	return b
}

func (b *HostBuilder) AddServices(service []Service) *HostBuilder {
	if b.host.Services == nil {
		b.host.Services = []Service{}
	}
	b.host.Services = append(b.host.Services, service...)
	return b
}

func (b *HostBuilder) Build() (*Host, error) {
	if b.host.IP == "" {
		return nil, errors.New("IP is a required field")
	}
	return b.host, nil
}
