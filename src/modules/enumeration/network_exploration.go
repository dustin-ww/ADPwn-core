package enumeration

import (
	"ADPwn/adapter/serializable/nmap"
	"ADPwn/adapter/tools"
	"ADPwn/core/model"
	"ADPwn/core/plugin" // Import plugin statt registry
	"ADPwn/core/service"
	"fmt"
)

type NetworkExplorer struct {
	Name            string
	Description     string
	Version         string
	Author          string
	ExecutionMetric string
	Dependencies    []string
	Modes           []string
	projectService  *service.ProjectService
	nmapAdapter     *adapter.NmapAdapter
}

func (n *NetworkExplorer) GetExecutionMetric() string {
	return n.ExecutionMetric
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

	nmapAdapter := adapter.NewNmapAdapter()
	nmapOptions := []adapter.NmapOption{
		adapter.FullScan,
	}

	projectAddressList, err := project.TargetsAsAddressList()
	if err != nil {
		fmt.Println("Error getting target addresses")
		return err
	}
	_, err = nmapAdapter.RunCommand(projectAddressList, nmapOptions)

	if err != nil {
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
			serviceBuilder := model.NewServiceBuilder()
			serviceBuilder.WithName(port.Service.Name)
			serviceBuilder.WithPort(port.Portid)
			buildService, err := serviceBuilder.Build()
			if err != nil {
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
		Name:            "Network Exploration",
		Description:     "Network Exploration",
		Version:         "0.1",
		Author:          "dw-sec",
		ExecutionMetric: "1h",
	}
	plugin.RegisterEnumeration(module)
}
