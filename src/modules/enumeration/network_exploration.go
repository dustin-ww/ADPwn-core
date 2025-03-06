package enumeration

import (
	"ADPwn/adapter/serializable/nmap"
	"ADPwn/adapter/tools"
	"ADPwn/cmd/logger"
	"ADPwn/core/model"
	"ADPwn/core/service"
	"ADPwn/modules/internal/base"
	"fmt"
)

type NetworkExplorer struct {
	Name           string
	Description    string
	Version        string
	Author         string
	Dependencies   []string
	Modes          []string
	Logger         *logger.ADPwnLogger
	projectService *service.ProjectService
	nmapAdapter    *adapter.NmapAdapter
}

func (n *NetworkExplorer) DependsOn() int {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkExplorer) GetName() string {
	return n.Name
}

func (n *NetworkExplorer) GetDescription() string {
	return n.Description
}

func (n *NetworkExplorer) GetVersion() string {
	return n.Version
}

func (n *NetworkExplorer) GetAuthor() string {
	return n.Author
}

func (n *NetworkExplorer) Execute(project model.Project, options []string) error {

	n.Logger.Log("[*] Starting AD network enumeration")
	n.Logger.Log(fmt.Sprintf("[*] Scanning project: %s", project.Name))
	n.Logger.Log(fmt.Sprintf("[*] Options: %v", options))

	nmapAdapter := adapter.NewNmapAdapter()
	nmapOptions := []adapter.NmapOption{
		adapter.FullScan,
	}

	projectAddressList, err := project.TargetsAsAddressList()
	if err != nil {
		n.Logger.Log(fmt.Sprintf("[*] Error getting target addresses: %s", err.Error()))
		fmt.Println("Error getting target addresses")
		return err
	}
	_, err = nmapAdapter.RunCommand(projectAddressList, nmapOptions)

	if err != nil {
		n.Logger.Log(fmt.Sprintf("[*] Error running command: %s", err.Error()))
		fmt.Println("Error running command")
	}

	return nil
}

// Build Domain -> Hosts and Services

func (n *NetworkExplorer) buildDomain(project model.Project, result nmap.Result) {

}

func (n *NetworkExplorer) buildHosts(project model.Project, result nmap.Result) {
	var hosts []model.Host

	for _, host := range result.Host {

		n.Logger.Log("[*] New Host discovered!")
		hostBuilder := model.NewHostBuilder()

		if n.isDomainController(host.Ports) {
			hostBuilder.AsDomainController()
		}

		services, err := n.buildServices(host)
		hostBuilder.AddServices(services)

		host, err := hostBuilder.Build()
		if err != nil {
			return
		}

		hosts = append(hosts, *host)

	}
}

func (n *NetworkExplorer) buildServices(host nmap.Host) ([]model.Service, error) {
	var services []model.Service

	for _, port := range host.Ports.Port {
		if port.State.State == "open" {
			n.Logger.Log("[*] New Service discovered!")
			serviceBuilder := model.NewServiceBuilder()
			serviceBuilder.WithName(port.Service.Name)
			serviceBuilder.WithPort(port.Portid)
			buildService, err := serviceBuilder.Build()
			if err != nil {
				n.Logger.Log(fmt.Sprintf("[*] Error building service: %s", err.Error()))
				return nil, err
			}
			services = append(services, *buildService)
		}
	}
	return services, nil
}

func (n *NetworkExplorer) isDomainController(ports nmap.Ports) bool {
	dcPorts := map[string]bool{
		"53":   true, // DNS
		"88":   true, // Kerberos
		"389":  true, // LDAP
		"445":  true, // SMB
		"464":  true, // Kerberos password change
		"636":  true, // LDAPS
		"3268": true, // Global Catalog
		"3269": true, // Global Catalog over SSL
	}

	matchCount := 0

	for _, port := range ports.Port {
		if port.State.State == "open" {
			if dcPorts[port.Portid] {
				matchCount++
			}

			// 88 is a reliable sign for domain controller
			if port.Portid == "88" {
				return true
			}

			// Search for other known dc services
			if port.Service.Name == "ldap" ||
				port.Service.Name == "kerberos" ||
				port.Service.Name == "msrpc" {
				matchCount++
			}
		}
	}

	return matchCount >= 3
}

// INIT
func init() {
	module := &NetworkExplorer{
		Name:        "NetworkExploration",
		Description: "ADPwn Module to enumerate ad network",
		Version:     "0.1",
		Author:      "Dustin Wickert",
	}
	base.GlobalRegistry.RegisterEnumerationModule(module)
}
