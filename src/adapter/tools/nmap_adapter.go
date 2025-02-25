package tools

import (
	"ADPwn/adapter/serializable"
	"context"
	"errors"
	"log"
	"os/exec"
	"time"
)

type NmapOption int

const (
	OutputXML               NmapOption = iota // -oX
	ServiceVersionDetection                   // -sV
	CommonPorts                               // --top-ports 1000
	ScriptScan                                // -sC
)

func (o NmapOption) String() string {
	switch o {
	case OutputXML:
		return "-oX"
	case ServiceVersionDetection:
		return "-sV"
	case CommonPorts:
		return "--top-ports 1000"
	case ScriptScan:
		return "-sC"
	default:
		return ""
	}
}

func RunCommand(targetAddresses []string, options []NmapOption) (serializable.NmapResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3000*time.Second)
	defer cancel()

	var nmapArgs []string
	for _, opt := range options {
		nmapArgs = append(nmapArgs, opt.String())
	}
	nmapArgs = append(nmapArgs, targetAddresses...)

	log.Printf("Executing nmap with args: %v", nmapArgs)

	cmd := exec.CommandContext(ctx, "nmap", nmapArgs...)

	out, err := cmd.Output()
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			log.Println("Error: Nmap command timed out")
		} else {
			log.Printf("Error executing nmap command: %v", err)
		}
		return serializable.NmapResult{}, err
	}

	log.Println("Nmap command executed successfully")

	// Parse die XML-Ausgabe in ein Nmaprun-Objekt
	var nmapRun serializable.NmapResult
	if err := nmapRun.NewFromXML(out); err != nil {
		log.Printf("Error parsing nmap output: %v", err)
		return serializable.Nmaprun{}, err
	}

	log.Println("Nmap output parsed successfully")
	return nmapRun, nil
}
