package model

type Host struct {
	IP                 string `json:"ip"`
	IsDomaincontroller bool   `json:"isDomaincontroller"`
	Type               string `json:"type"`
}

func NewHost(IP string) *Host {
	return &Host{
		IP:                 IP,
		IsDomaincontroller: false,
		Type:               "Client",
	}
}
