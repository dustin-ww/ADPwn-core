package model

import (
	"fmt"
	"net"
)

type Project struct {
	UID     string   `json:"uid,omitempty"`
	Name    string   `json:"name"`
	Domains []Domain `json:"has_domain"`
	Targets []string `json:"has_target"`
	DType   []string `json:"dgraph.type,omitempty"`
}

func NewProject(name string) *Project {

	return &Project{
		Name:  name,
		DType: []string{"project"},
	}
}

func (p *Project) TargetsAsAddressList() ([]string, error) {
	var unifiedIPs []string

	for _, item := range p.Targets {
		if _, ipNet, err := net.ParseCIDR(item); err == nil {
			for ip := ipNet.IP.Mask(ipNet.Mask); ipNet.Contains(ip); p.incrementIP(ip) {
				unifiedIPs = append(unifiedIPs, ip.String())
			}
		} else if ip := net.ParseIP(item); ip != nil {
			unifiedIPs = append(unifiedIPs, ip.String())
		} else {
			return nil, fmt.Errorf("invalid IP or CIDR: %s", item)
		}
	}
	return unifiedIPs, nil
}
func (p *Project) incrementIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}

}
