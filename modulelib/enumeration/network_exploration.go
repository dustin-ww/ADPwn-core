package enumeration

import (
	plugin "ADPwn-core/internal/module"
	"ADPwn-core/internal/sse"
	"ADPwn-core/pkg/adapter/serializable"
	adapter "ADPwn-core/pkg/adapter/tools"
	"ADPwn-core/pkg/model"
	"ADPwn-core/pkg/model/adpwn/input"
	"ADPwn-core/pkg/service"
	"fmt"
	"log"
	"time"
)

// INITIALIZE MODULE AS ADPWN PLUGIN
func init() {
	module := &NetworkExplorer{
		configKey: "NetworkExplorer",
	}
	plugin.RegisterPlugin(module)
}

type NetworkExplorer struct {
	// Internal
	configKey string
	// Services
	projectService *service.ProjectService
	// Tool Adapter
	nmapAdapter *adapter.NmapAdapter
}

func (n *NetworkExplorer) ConfigKey() string {
	return n.configKey
}

func (n *NetworkExplorer) ExecuteModule(params *input.Parameter, logger *sse.SSELogger) error {
	// Log start of module execution
	log.Printf("Executing module key: %s", n.ConfigKey)
	logger.Info(fmt.Sprintf("Starting module: %s", n.ConfigKey))

	logger.Event("scan_start", map[string]interface{}{
		"target_network": params.Inputs["network"],
		"ports":          "1-1024",
	})

	// Log completion of module
	logger.Event("module_complete", map[string]interface{}{
		"moduleKey": n.ConfigKey,
		"timestamp": time.Now().Unix(),
	})

	return nil
}

// Build Domain -> Hosts and Services

func (n *NetworkExplorer) buildDomain(project model.Project, result serializable.NmapResult) {

}

func (n *NetworkExplorer) buildHosts(project model.Project, result serializable.NmapResult) {
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

func (n *NetworkExplorer) buildServices(host serializable.Host) ([]model.Service, error) {
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

func (n *NetworkExplorer) isDomainController(ports serializable.Ports) bool {
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
