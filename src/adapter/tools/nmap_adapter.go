package tools

import (
	"ADPwn/adapter/serializable"
	"context"
	"errors"
	"log"
	"os/exec"
	"time"
)

// NmapOption definiert die verfügbaren Nmap-Optionen als Enum.
type NmapOption int

const (
	OutputXML               NmapOption = iota // -oX
	ServiceVersionDetection                   // -sV
	CommonPorts                               // --top-ports 1000
	ScriptScan                                // -sC
)

// String konvertiert das Enum in den entsprechenden Nmap-Options-String.
func (o NmapOption) String() string {
	switch o {
	case OutputXML:
		return "-oX"
	case ServiceVersionDetection:
		return "-sV"
	case CommonPorts:
		return "--top-ports 1000" // Kombiniert --top-ports mit dem Wert 1000
	case ScriptScan:
		return "-sC"
	default:
		return ""
	}
}

// RunCommand führt den Nmap-Befehl mit den angegebenen Zieladressen und Optionen aus.
func RunCommand(targetAddresses []string, options []NmapOption) (serializable.NmapResult, error) {
	// Setze ein Timeout von 30 Sekunden für den Nmap-Befehl
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Konvertiere die Enum-Optionen in Strings
	var nmapArgs []string
	for _, opt := range options {
		nmapArgs = append(nmapArgs, opt.String())
	}
	nmapArgs = append(nmapArgs, targetAddresses...) // Füge die Zieladressen hinzu

	log.Printf("Executing nmap with args: %v", nmapArgs)

	// Erstelle den Nmap-Befehl
	cmd := exec.CommandContext(ctx, "nmap", nmapArgs...)

	// Führe den Befehl aus und erfasse die Ausgabe
	out, err := cmd.Output()
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			log.Println("Error: Nmap command timed out")
		} else {
			log.Printf("Error executing nmap command: %v", err)
		}
		return serializable.Nmaprun{}, err
	}

	log.Println("Nmap command executed successfully")

	// Parse die XML-Ausgabe in ein Nmaprun-Objekt
	var nmapRun serializable.Nmaprun
	if err := nmapRun.NewFromXML(out); err != nil {
		log.Printf("Error parsing nmap output: %v", err)
		return serializable.Nmaprun{}, err
	}

	log.Println("Nmap output parsed successfully")
	return nmapRun, nil
}
