package tools

import (
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

func Execute() {
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

func main() {
	/* if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		Execute()
	} */
	fmt.Printf("%v", DomainControllers())
}

func readNmapXML() serializable.Nmaprun {
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
	return domainControllersByPort88()
}

func domainControllersByPort88() []model.DomainController {
	nmapRun := readNmapXML()
	var domainControllers []model.DomainController

	for _, host := range nmapRun.Host {
		for _, port := range host.Ports.Port {
			if port.Portid == "88" {
				domainControllers = append(domainControllers, model.DomainController{Ipv4: host.Address[0].Addr, Hostname: host.Hostnames, Reliablity: model.Safe})
			}
		}
	}
	fmt.Println("Namap found 2 Domain Controllers by Port 88")
	return domainControllers
}
