package tools

import (
	app_model "ADPwn/database/project/model"
	"ADPwn/database/project/service"
	"ADPwn/tools/serializable"
	"context"
	"errors"
	"log"
	"os/exec"
	"time"
)

type Nmap struct {
	service service.ProjectService
}

func (n *Nmap) ExecuteFullRecon(project app_model.Project) {

	nmapRun := n.runCommand(project)
	n.AddHosts(nmapRun, project)

}

func (n *Nmap) runCommand(project app_model.Project) serializable.Nmaprun {
	ctx, cancel := context.WithTimeout(context.Background(), 30000*time.Second)
	defer cancel()

	//options := []string{"-oX", "-", "-sVC"}
	options := []string{"-oX", "-"}
	args := append(options, project.Targets...)
	log.Println(project.Targets)

	log.Println(args)

	cmd := exec.CommandContext(ctx, "nmap", append(options, project.Targets...)...)

	out, err := cmd.Output()
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			log.Println("Error: Command timed out")
		} else {
			log.Fatalf("Error executing command: %s\n", err)
		}
		return serializable.Nmaprun{}
	}
	log.Println("Command Successfully Executed")

	var nmapRun serializable.Nmaprun
	log.Println("FINISH")
	log.Println(string(out))
	return nmapRun.NewFromXML(out)
}

func (n *Nmap) AddHosts(nmapRun serializable.Nmaprun, project app_model.Project) {
	for _, host := range nmapRun.Host {
		newHost := app_model.NewHost(host.Address[0].Addr, project.UID, project.Name)

		for _, port := range host.Ports.Port {
			newService := app_model.NewService(port.Service.Name, port.Portid)
			newHost.Services = append(newHost.Services, *newService)
			// Domaincontrollers
			if port.Portid == "88" {
				newHost.IsDomaincontroller = true
			}
		}
		_, err := n.service.AddHost(context.Background(), project, *newHost)
		if err != nil {
			return
		}
	}
	log.Println("")
}
