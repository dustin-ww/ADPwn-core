package utils

import (
	"fmt"
	"log"
	"net"
)

func GenerateIPs(subnet string) ([]string, error) {
	ip, ipNet, err := net.ParseCIDR(subnet)
	if err != nil {
		return nil, fmt.Errorf("ungültiges Subnetz: %v", err)
	}

	var ips []string
	for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); incrementIP(ip) {
		if !isPrivateIP(ip.String()) {
			log.Fatal("This tool is only valid for private ip adres")
		}

		ips = append(ips, ip.String())
	}

	// remove broadcast adress
	if len(ips) > 0 {
		ips = ips[:len(ips)-1]
	}

	return ips, nil
}

func incrementIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func isPrivateIP(ip string) bool {
	privateBlocks := []string{
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
	}

	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}

	for _, block := range privateBlocks {
		_, subnet, _ := net.ParseCIDR(block)
		if subnet.Contains(parsedIP) {
			return true
		}
	}

	return false
}
