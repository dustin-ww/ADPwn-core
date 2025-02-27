package adapter

import (
	"ADPwn/adapter/serializable/nmap"
	_ "ADPwn/adapter/serializable/nmap"
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"time"
)

type NmapAdapter struct{}

type NmapOption int

const (
	ServiceVersionDetection NmapOption = iota
	CommonPorts
	ScriptScan
	FullScan
	UDPScan
)

var NmapOpt = struct {
	ServiceVersionDetection NmapOption
	CommonPorts             NmapOption
	ScriptScan              NmapOption
	FullScan                NmapOption
	UDPScan                 NmapOption
}{
	ServiceVersionDetection: ServiceVersionDetection,
	CommonPorts:             CommonPorts,
	ScriptScan:              ScriptScan,
	FullScan:                FullScan,
	UDPScan:                 UDPScan,
}

func (o NmapOption) String() string {
	switch o {
	case ServiceVersionDetection:
		return "-sV"
	case CommonPorts:
		return "--top-ports 1000"
	case ScriptScan:
		return "-sC"
	case FullScan:
		return "-sVC"
	case UDPScan:
		return "-sU"
	default:
		return ""
	}
}

var ErrNmapTimeout = errors.New("nmap command timed out")

func (n *NmapAdapter) RunCommand(targetAddresses []string, options []NmapOption) (nmap.Result, error) {
	if len(targetAddresses) == 0 {
		return nmap.Result{}, errors.New("no target addresses provided")
	}

	// 50 min timeout
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Minute)
	defer cancel()

	var nmapArgs []string
	for _, opt := range options {
		nmapArgs = append(nmapArgs, opt.String())
	}
	nmapArgs = append(nmapArgs, "-oX -")
	nmapArgs = append(nmapArgs, targetAddresses...)

	log.Printf("Executing nmap with args: %v", nmapArgs)

	cmd := exec.CommandContext(ctx, "nmap", nmapArgs...)

	out, err := cmd.CombinedOutput()
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			log.Println("Error: Nmap command timed out")
			return nmap.Result{}, ErrNmapTimeout
		}
		log.Printf("Error executing nmap command: %v, output: %s", err, out)
		return nmap.Result{}, fmt.Errorf("nmap execution failed: %w", err)
	}

	log.Println("Nmap command executed successfully")

	nmapRun, err := n.parseXML(out)
	if err != nil {
		log.Printf("Error parsing nmap output: %v", err)
		return nmap.Result{}, fmt.Errorf("parse xml: %w", err)
	}

	log.Println("Nmap output parsed successfully")
	return nmapRun, nil
}

func (n *NmapAdapter) parseXML(nmapXML []byte) (nmap.Result, error) {
	var nmapResult nmap.Result

	err := xml.Unmarshal(nmapXML, &nmapResult)
	if err != nil {
		return nmap.Result{}, fmt.Errorf("unmarshal XML: %w", err)
	}
	return nmapResult, nil
}

func NewNmapAdapter() *NmapAdapter {
	return &NmapAdapter{}
}
