package model

import (
	"fmt"
	"net"
	"time"
)

type Project struct {
	UID         string    `json:"uid,omitempty"`
	Name        string    `json:"name,omitempty"`
	Tags        []string  `json:"tags,omitempty"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty"`
	HasTarget   []Target  `json:"has_target,omitempty"`
	HasDomain   []Domain  `json:"has_domain,omitempty"`
	DType       []string  `json:"dgraph.type,omitempty"`
}

func (p *Project) TargetsAsAddressList() ([]string, error) {
	var unifiedIPs []string

	for _, item := range p.HasTarget {
		if _, ipNet, err := net.ParseCIDR(item.IPRange); err == nil {
			for ip := ipNet.IP.Mask(ipNet.Mask); ipNet.Contains(ip); p.incrementIP(ip) {
				unifiedIPs = append(unifiedIPs, ip.String())
			}
		} else if ip := net.ParseIP(item.IPRange); ip != nil {
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
