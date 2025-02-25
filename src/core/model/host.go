package model

type Host struct {
	UID                string    `json:"uid,omitempty"`
	IP                 string    `json:"ip"`
	Name               string    `json:"name"`
	HostProjectID      string    `json:"hostProjectID"`
	IsDomaincontroller bool      `json:"isDomaincontroller"`
	Services           []Service `json:"has_service,omitempty"`
	DType              []string  `json:"dgraph.type,omitempty"`
}

func NewHost(IP string, projectUID string, projectName string) *Host {
	return &Host{
		IP:                 IP,
		HostProjectID:      projectUID + "_" + IP,
		Name:               projectName + "_" + IP,
		IsDomaincontroller: false,
		DType:              []string{"host"},
	}
}

func (h Host) AddService(service Service) {
	h.Services = append(h.Services, service)
}
