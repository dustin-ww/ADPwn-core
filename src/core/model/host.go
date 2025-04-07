package model

import (
	"errors"
	"time"
)

type Host struct {
	// Internal
	UID                string    `json:"uid,omitempty"`
	IP                 string    `json:"ip,omitempty"`
	IsDomainController bool      `json:"is_domain_controller,omitempty"`
	BelongsToDomain    Domain    `json:"belongs_to_domain,omitempty"`
	HasService         []Service `json:"has_service,omitempty"`
	DType              []string  `json:"dgraph.type,omitempty"`
	InternalCreatedAt  time.Time `json:"internal_created_at,omitempty"`
	// AD related
	DistinguishedName      string    `json:"distinguishedName"`
	ObjectGUID             string    `json:"objectGUID"`
	ObjectSid              string    `json:"objectSid"`
	SAMAccountName         string    `json:"sAMAccountName"`
	DNSHostName            string    `json:"dNSHostName"`
	OperatingSystem        string    `json:"operatingSystem"`
	OperatingSystemVersion string    `json:"operatingSystemVersion"`
	LastLogonTimestamp     time.Time `json:"lastLogonTimestamp"`
	WhenCreated            time.Time `json:"whenCreated"`
	WhenChanged            time.Time `json:"whenChanged"`
	UserAccountControl     int       `json:"userAccountControl"`
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
