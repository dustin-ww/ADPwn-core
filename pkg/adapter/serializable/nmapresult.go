package serializable

import (
	"encoding/xml"
)

type NmapResult struct {
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

type ScanInfo struct {
	Text        string `xml:",chardata"`
	Type        string `xml:"type,attr"`
	Protocol    string `xml:"protocol,attr"`
	Numservices string `xml:"numservices,attr"`
	Services    string `xml:"services,attr"`
}

type Verbose struct {
	Text  string `xml:",chardata"`
	Level string `xml:"level,attr"`
}

type Debugging struct {
	Text  string `xml:",chardata"`
	Level string `xml:"level,attr"`
}

type Hosthint struct {
	Text      string    `xml:",chardata"`
	Status    Status    `xml:"status"`
	Address   []Address `xml:"address"`
	Hostnames string    `xml:"hostnames"`
}

type Status struct {
	Text      string `xml:",chardata"`
	State     string `xml:"state,attr"`
	Reason    string `xml:"reason,attr"`
	ReasonTtl string `xml:"reason_ttl,attr"`
}

type Address struct {
	Text     string `xml:",chardata"`
	Addr     string `xml:"addr,attr"`
	Addrtype string `xml:"addrtype,attr"`
	Vendor   string `xml:"vendor,attr"`
}

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

type Ports struct {
	Text       string     `xml:",chardata"`
	Extraports Extraports `xml:"extraports"`
	Port       []Port     `xml:"port"`
}

type Extraports struct {
	Text         string       `xml:",chardata"`
	State        string       `xml:"state,attr"`
	Count        string       `xml:"count,attr"`
	Extrareasons Extrareasons `xml:"extrareasons"`
}

type Extrareasons struct {
	Text   string `xml:",chardata"`
	Reason string `xml:"reason,attr"`
	Count  string `xml:"count,attr"`
	Proto  string `xml:"proto,attr"`
	Ports  string `xml:"ports,attr"`
}

type Port struct {
	Text     string   `xml:",chardata"`
	Protocol string   `xml:"protocol,attr"`
	Portid   string   `xml:"portid,attr"`
	State    State    `xml:"state"`
	Service  Service  `xml:"service"`
	Script   []Script `xml:"script"`
}

type State struct {
	Text      string `xml:",chardata"`
	State     string `xml:"state,attr"`
	Reason    string `xml:"reason,attr"`
	ReasonTtl string `xml:"reason_ttl,attr"`
}

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

type Script struct {
	Text   string        `xml:",chardata"`
	ID     string        `xml:"id,attr"`
	Output string        `xml:"output,attr"`
	Table  []ScriptTable `xml:"table"`
	Elem   []ScriptElem  `xml:"elem"`
}

type ScriptTable struct {
	Text  string        `xml:",chardata"`
	Key   string        `xml:"key,attr"`
	Elem  []ScriptElem  `xml:"elem"`
	Table []NestedTable `xml:"table"`
}

type NestedTable struct {
	Text  string       `xml:",chardata"`
	Key   string       `xml:"key,attr"`
	Elem  []ScriptElem `xml:"elem"`
	Table DeepTable    `xml:"table"`
}

type DeepTable struct {
	Text string   `xml:",chardata"`
	Key  string   `xml:"key,attr"`
	Elem []string `xml:"elem"`
}

type ScriptElem struct {
	Text string `xml:",chardata"`
	Key  string `xml:"key,attr"`
}

type Hostscript struct {
	Text   string           `xml:",chardata"`
	Script []HostScriptItem `xml:"script"`
}

type HostScriptItem struct {
	Text   string           `xml:",chardata"`
	ID     string           `xml:"id,attr"`
	Output string           `xml:"output,attr"`
	Table  HostScriptTable  `xml:"table"`
	Elem   []HostScriptElem `xml:"elem"`
}

type HostScriptTable struct {
	Text string `xml:",chardata"`
	Key  string `xml:"key,attr"`
	Elem string `xml:"elem"`
}

type HostScriptElem struct {
	Text string `xml:",chardata"`
	Key  string `xml:"key,attr"`
}

type Times struct {
	Text   string `xml:",chardata"`
	Srtt   string `xml:"srtt,attr"`
	Rttvar string `xml:"rttvar,attr"`
	To     string `xml:"to,attr"`
}

type Postscript struct {
	Text   string           `xml:",chardata"`
	Script PostscriptScript `xml:"script"`
}

type PostscriptScript struct {
	Text   string          `xml:",chardata"`
	ID     string          `xml:"id,attr"`
	Output string          `xml:"output,attr"`
	Table  PostscriptTable `xml:"table"`
}

type PostscriptTable struct {
	Text string   `xml:",chardata"`
	Key  string   `xml:"key,attr"`
	Elem []string `xml:"elem"`
}

type Runstats struct {
	Text     string   `xml:",chardata"`
	Finished Finished `xml:"finished"`
	Hosts    Hosts    `xml:"hosts"`
}

type Finished struct {
	Text    string `xml:",chardata"`
	Time    string `xml:"time,attr"`
	Timestr string `xml:"timestr,attr"`
	Summary string `xml:"summary,attr"`
	Elapsed string `xml:"elapsed,attr"`
	Exit    string `xml:"exit,attr"`
}

type Hosts struct {
	Text  string `xml:",chardata"`
	Up    string `xml:"up,attr"`
	Down  string `xml:"down,attr"`
	Total string `xml:"total,attr"`
}
