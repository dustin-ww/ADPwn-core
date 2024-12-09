package model

type Host struct {
	UID                string   `json:"uid,omitempty"`
	IP                 string   `json:"ip"`
	HostProjectID      string   `json:"host_project_id"`
	IsDomaincontroller bool     `json:"is_domaincontroller"`
	DType              []string `json:"dgraph.type,omitempty"`
}

func NewHost(IP string, ProjectUID string) *Host {
	return &Host{
		IP:                 IP,
		HostProjectID:      ProjectUID + "|" + IP,
		IsDomaincontroller: false,
		DType:              []string{"Host"},
	}
}
