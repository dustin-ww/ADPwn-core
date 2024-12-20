package tools

import (
	model2 "ADPwn/database/project/model"
	"ADPwn/tools/model"
	"ADPwn/tools/serializable"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
)

type Nmap struct {
}

func (n *Nmap) Execute(project model2.Project) {
	n.runCommand()
	nmapResult := n.parseResultXML()
	n.filterPort88DomainController(nmapResult)
	// projectService.add
}

func (n *Nmap) runCommand() {
	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	} else {
		log.Println(dir)
		fmt.Println(dir)
	}
	out, err := exec.Command("nmap", "-oX", "localhost").Output()

	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println("Command Successfully Executed")
	output := string(out[:])
	fmt.Println(output)

	e, err := os.Executable()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(path.Dir(e))

}

func (n *Nmap) parseResultXML() serializable.Nmaprun {
	path, _ := os.Getwd()
	nmapXML, err := os.Open(path + "/out/nmap.xml")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened nmap.xml")
	byteValue, _ := io.ReadAll(nmapXML)

	var nmapRun serializable.Nmaprun
	xml.Unmarshal(byteValue, &nmapRun)

	defer nmapXML.Close()
	return nmapRun
}

func DomainControllers() []model.DomainController {
	return nil
}

func (n *Nmap) filterHosts(nmapRun serializable.Nmaprun) {

}

func (n *Nmap) filterPort88DomainController(nmapResult serializable.Nmaprun) []model.DomainController {
	var domainControllers []model.DomainController

	for _, host := range nmapResult.Host {
		for _, port := range host.Ports.Port {
			if port.Portid == "88" {
				domainControllers = append(domainControllers, model.DomainController{Ipv4: host.Address[0].Addr, Hostname: host.Hostnames, Reliablity: model.Safe})
			}
		}
	}
	fmt.Println("Namap found 2 Domain Controllers by Port 88")
	return domainControllers
}
