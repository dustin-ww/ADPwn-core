package nmap

import (
	"encoding/xml"
)

// Hauptdatenstruktur für Nmap-Ergebnisse
type Result struct {
	XMLName          xml.Name   `xml:"nmaprun"`
	Text             string     `xml:",chardata"`
	Scanner          string     `xml:"scanner,attr"`
	Args             string     `xml:"args,attr"`
	Start            string     `xml:"start,attr"`
	Startstr         string     `xml:"startstr,attr"`
	Version          string     `xml:"version,attr"`
	Xmloutputversion string     `xml:"xmloutputversion,attr"`
	Scaninfo         ScanInfo   `xml:"scaninfo"`
	Verbose          Verbose    `xml:"verbose"`
	Debugging        Debugging  `xml:"debugging"`
	Hosthint         []Hosthint `xml:"hosthint"`
	Host             []Host     `xml:"host"`
	Postscript       Postscript `xml:"postscript"`
	Runstats         Runstats   `xml:"runstats"`
}

// ScanInfo enthält Informationen über den Scan
type ScanInfo struct {
	Text        string `xml:",chardata"`
	Type        string `xml:"type,attr"`
	Protocol    string `xml:"protocol,attr"`
	Numservices string `xml:"numservices,attr"`
	Services    string `xml:"services,attr"`
}

// Verbose enthält Informationen über die Verbose-Einstellung
type Verbose struct {
	Text  string `xml:",chardata"`
	Level string `xml:"level,attr"`
}

// Debugging enthält Informationen über die Debugging-Einstellung
type Debugging struct {
	Text  string `xml:",chardata"`
	Level string `xml:"level,attr"`
}

// Hosthint enthält Hinweise auf Hosts
type Hosthint struct {
	Text      string    `xml:",chardata"`
	Status    Status    `xml:"status"`
	Address   []Address `xml:"address"`
	Hostnames string    `xml:"hostnames"`
}

// Status enthält Statusinformationen
type Status struct {
	Text      string `xml:",chardata"`
	State     string `xml:"state,attr"`
	Reason    string `xml:"reason,attr"`
	ReasonTtl string `xml:"reason_ttl,attr"`
}

// Address enthält Adressinformationen
type Address struct {
	Text     string `xml:",chardata"`
	Addr     string `xml:"addr,attr"`
	Addrtype string `xml:"addrtype,attr"`
	Vendor   string `xml:"vendor,attr"`
}

// Host repräsentiert einen gescannten Host
type Host struct {
	Text       string     `xml:",chardata"`
	Starttime  string     `xml:"starttime,attr"`
	Endtime    string     `xml:"endtime,attr"`
	Status     Status     `xml:"status"`
	Address    []Address  `xml:"address"`
	Hostnames  string     `xml:"hostnames"`
	Ports      Ports      `xml:"ports"`
	Hostscript Hostscript `xml:"hostscript"`
	Times      Times      `xml:"times"`
}

// Ports enthält alle Port-Informationen eines Hosts
type Ports struct {
	Text       string     `xml:",chardata"`
	Extraports Extraports `xml:"extraports"`
	Port       []Port     `xml:"port"`
}

// Extraports enthält Informationen über nicht gescannte Ports
type Extraports struct {
	Text         string       `xml:",chardata"`
	State        string       `xml:"state,attr"`
	Count        string       `xml:"count,attr"`
	Extrareasons Extrareasons `xml:"extrareasons"`
}

// Extrareasons enthält Gründe für Extraports
type Extrareasons struct {
	Text   string `xml:",chardata"`
	Reason string `xml:"reason,attr"`
	Count  string `xml:"count,attr"`
	Proto  string `xml:"proto,attr"`
	Ports  string `xml:"ports,attr"`
}

// Port repräsentiert einen einzelnen gescannten Port
type Port struct {
	Text     string   `xml:",chardata"`
	Protocol string   `xml:"protocol,attr"`
	Portid   string   `xml:"portid,attr"`
	State    State    `xml:"state"`
	Service  Service  `xml:"service"`
	Script   []Script `xml:"script"`
}

// State enthält den Zustand eines Ports
type State struct {
	Text      string `xml:",chardata"`
	State     string `xml:"state,attr"`
	Reason    string `xml:"reason,attr"`
	ReasonTtl string `xml:"reason_ttl,attr"`
}

// Service enthält Informationen über einen Dienst
type Service struct {
	Text      string   `xml:",chardata"`
	Name      string   `xml:"name,attr"`
	Product   string   `xml:"product,attr"`
	Ostype    string   `xml:"ostype,attr"`
	Method    string   `xml:"method,attr"`
	Conf      string   `xml:"conf,attr"`
	Version   string   `xml:"version,attr"`
	Extrainfo string   `xml:"extrainfo,attr"`
	Hostname  string   `xml:"hostname,attr"`
	Tunnel    string   `xml:"tunnel,attr"`
	Cpe       []string `xml:"cpe"`
}

// Script enthält Informationen über ein Nmap-Script
type Script struct {
	Text   string        `xml:",chardata"`
	ID     string        `xml:"id,attr"`
	Output string        `xml:"output,attr"`
	Table  []ScriptTable `xml:"table"`
	Elem   []ScriptElem  `xml:"elem"`
}

// ScriptTable enthält Tabellendaten eines Scripts
type ScriptTable struct {
	Text  string        `xml:",chardata"`
	Key   string        `xml:"key,attr"`
	Elem  []ScriptElem  `xml:"elem"`
	Table []NestedTable `xml:"table"`
}

// NestedTable enthält verschachtelte Tabellen
type NestedTable struct {
	Text  string       `xml:",chardata"`
	Key   string       `xml:"key,attr"`
	Elem  []ScriptElem `xml:"elem"`
	Table DeepTable    `xml:"table"`
}

// DeepTable ist eine noch tiefere Tabelle
type DeepTable struct {
	Text string   `xml:",chardata"`
	Key  string   `xml:"key,attr"`
	Elem []string `xml:"elem"`
}

// ScriptElem enthält Elemente eines Scripts
type ScriptElem struct {
	Text string `xml:",chardata"`
	Key  string `xml:"key,attr"`
}

// Hostscript enthält Host-spezifische Scripts
type Hostscript struct {
	Text   string           `xml:",chardata"`
	Script []HostScriptItem `xml:"script"`
}

// HostScriptItem enthält ein einzelnes Host-Script
type HostScriptItem struct {
	Text   string           `xml:",chardata"`
	ID     string           `xml:"id,attr"`
	Output string           `xml:"output,attr"`
	Table  HostScriptTable  `xml:"table"`
	Elem   []HostScriptElem `xml:"elem"`
}

// HostScriptTable enthält Tabellendaten eines Host-Scripts
type HostScriptTable struct {
	Text string `xml:",chardata"`
	Key  string `xml:"key,attr"`
	Elem string `xml:"elem"`
}

// HostScriptElem enthält Elemente eines Host-Scripts
type HostScriptElem struct {
	Text string `xml:",chardata"`
	Key  string `xml:"key,attr"`
}

// Times enthält Zeitmessungen
type Times struct {
	Text   string `xml:",chardata"`
	Srtt   string `xml:"srtt,attr"`
	Rttvar string `xml:"rttvar,attr"`
	To     string `xml:"to,attr"`
}

// Postscript enthält abschließende Scripts
type Postscript struct {
	Text   string           `xml:",chardata"`
	Script PostscriptScript `xml:"script"`
}

// PostscriptScript enthält ein einzelnes Postscript-Script
type PostscriptScript struct {
	Text   string          `xml:",chardata"`
	ID     string          `xml:"id,attr"`
	Output string          `xml:"output,attr"`
	Table  PostscriptTable `xml:"table"`
}

// PostscriptTable enthält Tabellendaten eines Postscripts
type PostscriptTable struct {
	Text string   `xml:",chardata"`
	Key  string   `xml:"key,attr"`
	Elem []string `xml:"elem"`
}

// Runstats enthält abschließende Statistiken
type Runstats struct {
	Text     string   `xml:",chardata"`
	Finished Finished `xml:"finished"`
	Hosts    Hosts    `xml:"hosts"`
}

// Finished enthält Informationen über den Abschluss
type Finished struct {
	Text    string `xml:",chardata"`
	Time    string `xml:"time,attr"`
	Timestr string `xml:"timestr,attr"`
	Summary string `xml:"summary,attr"`
	Elapsed string `xml:"elapsed,attr"`
	Exit    string `xml:"exit,attr"`
}

// Hosts enthält Hoststatistiken
type Hosts struct {
	Text  string `xml:",chardata"`
	Up    string `xml:"up,attr"`
	Down  string `xml:"down,attr"`
	Total string `xml:"total,attr"`
}
