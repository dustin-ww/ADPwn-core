package model

import "errors"

type Host struct {
	UID                string    `json:"uid,omitempty"`
	IP                 string    `json:"ip,omitempty"`
	IsDomainController bool      `json:"is_domain_controller,omitempty"`
	BelongsToDomain    Domain    `json:"belongs_to_domain,omitempty"`
	HasService         []Service `json:"has_service,omitempty"`
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
	b.host.IP = ip
	return b
}

func (b *HostBuilder) AsDomainController() *HostBuilder {
	b.host.IsDomainController = true
	return b
}

func (b *HostBuilder) WithServices(services []Service) *HostBuilder {
	b.host.HasService = services
	return b
}

func (b *HostBuilder) AddService(service Service) *HostBuilder {
	if b.host.HasService == nil {
		b.host.HasService = []Service{}
	}
	b.host.HasService = append(b.host.HasService, service)
	return b
}

func (b *HostBuilder) AddServices(service []Service) *HostBuilder {

	return b
}

func (b *HostBuilder) Build() (*Host, error) {
	if b.host.IP == "" {
		return nil, errors.New("IP is a required field")
	}
	return b.host, nil
}
