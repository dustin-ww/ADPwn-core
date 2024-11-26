package main

import (
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
	out, err := exec.Command("nmap", "-oX", "192.168.57.0/24").Output()

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

func readNmapXML() Nmaprun {
	path, _ := os.Getwd()
	nmapXML, err := os.Open(path + "/out/nmap.xml")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened nmap.xml")
	byteValue, _ := io.ReadAll(nmapXML)

	var nmapRun Nmaprun
	xml.Unmarshal(byteValue, &nmapRun)

	defer nmapXML.Close()
	return nmapRun
}

func DomainControllers() []DomainController {
	return domainControllersByPort88()
}

func domainControllersByPort88() []DomainController {
	nmapRun := readNmapXML()
	var domainControllers []DomainController

	for _, host := range nmapRun.Host {
		for _, port := range host.Ports.Port {
			if port.Portid == "88" {
				domainControllers = append(domainControllers, DomainController{host.Address.Addr, "not implemented"})
			}
		}
	}
	fmt.Println("Namap found 2 Domain Controllers by Port 88")
	return domainControllers
}

type Nmaprun struct {
	XMLName          xml.Name `xml:"nmaprun"`
	Text             string   `xml:",chardata"`
	Scanner          string   `xml:"scanner,attr"`
	Args             string   `xml:"args,attr"`
	Start            string   `xml:"start,attr"`
	Startstr         string   `xml:"startstr,attr"`
	Version          string   `xml:"version,attr"`
	Xmloutputversion string   `xml:"xmloutputversion,attr"`
	Scaninfo         struct {
		Text        string `xml:",chardata"`
		Type        string `xml:"type,attr"`
		Protocol    string `xml:"protocol,attr"`
		Numservices string `xml:"numservices,attr"`
		Services    string `xml:"services,attr"`
	} `xml:"scaninfo"`
	Verbose struct {
		Text  string `xml:",chardata"`
		Level string `xml:"level,attr"`
	} `xml:"verbose"`
	Debugging struct {
		Text  string `xml:",chardata"`
		Level string `xml:"level,attr"`
	} `xml:"debugging"`
	Hosthint []struct {
		Text   string `xml:",chardata"`
		Status struct {
			Text      string `xml:",chardata"`
			State     string `xml:"state,attr"`
			Reason    string `xml:"reason,attr"`
			ReasonTtl string `xml:"reason_ttl,attr"`
		} `xml:"status"`
		Address struct {
			Text     string `xml:",chardata"`
			Addr     string `xml:"addr,attr"`
			Addrtype string `xml:"addrtype,attr"`
		} `xml:"address"`
		Hostnames string `xml:"hostnames"`
	} `xml:"hosthint"`
	Host []struct {
		Text      string `xml:",chardata"`
		Starttime string `xml:"starttime,attr"`
		Endtime   string `xml:"endtime,attr"`
		Status    struct {
			Text      string `xml:",chardata"`
			State     string `xml:"state,attr"`
			Reason    string `xml:"reason,attr"`
			ReasonTtl string `xml:"reason_ttl,attr"`
		} `xml:"status"`
		Address struct {
			Text     string `xml:",chardata"`
			Addr     string `xml:"addr,attr"`
			Addrtype string `xml:"addrtype,attr"`
		} `xml:"address"`
		Hostnames string `xml:"hostnames"`
		Ports     struct {
			Text       string `xml:",chardata"`
			Extraports struct {
				Text         string `xml:",chardata"`
				State        string `xml:"state,attr"`
				Count        string `xml:"count,attr"`
				Extrareasons struct {
					Text   string `xml:",chardata"`
					Reason string `xml:"reason,attr"`
					Count  string `xml:"count,attr"`
					Proto  string `xml:"proto,attr"`
					Ports  string `xml:"ports,attr"`
				} `xml:"extrareasons"`
			} `xml:"extraports"`
			Port []struct {
				Text     string `xml:",chardata"`
				Protocol string `xml:"protocol,attr"`
				Portid   string `xml:"portid,attr"`
				State    struct {
					Text      string `xml:",chardata"`
					State     string `xml:"state,attr"`
					Reason    string `xml:"reason,attr"`
					ReasonTtl string `xml:"reason_ttl,attr"`
				} `xml:"state"`
				Service struct {
					Text   string `xml:",chardata"`
					Name   string `xml:"name,attr"`
					Method string `xml:"method,attr"`
					Conf   string `xml:"conf,attr"`
				} `xml:"service"`
			} `xml:"port"`
		} `xml:"ports"`
		Times struct {
			Text   string `xml:",chardata"`
			Srtt   string `xml:"srtt,attr"`
			Rttvar string `xml:"rttvar,attr"`
			To     string `xml:"to,attr"`
		} `xml:"times"`
	} `xml:"host"`
	Runstats struct {
		Text     string `xml:",chardata"`
		Finished struct {
			Text    string `xml:",chardata"`
			Time    string `xml:"time,attr"`
			Timestr string `xml:"timestr,attr"`
			Summary string `xml:"summary,attr"`
			Elapsed string `xml:"elapsed,attr"`
			Exit    string `xml:"exit,attr"`
		} `xml:"finished"`
		Hosts struct {
			Text  string `xml:",chardata"`
			Up    string `xml:"up,attr"`
			Down  string `xml:"down,attr"`
			Total string `xml:"total,attr"`
		} `xml:"hosts"`
	} `xml:"runstats"`
}
