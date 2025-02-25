package serializable

import "encoding/xml"

type NmapResult struct {
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
		Address []struct {
			Text     string `xml:",chardata"`
			Addr     string `xml:"addr,attr"`
			Addrtype string `xml:"addrtype,attr"`
			Vendor   string `xml:"vendor,attr"`
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
		Address []struct {
			Text     string `xml:",chardata"`
			Addr     string `xml:"addr,attr"`
			Addrtype string `xml:"addrtype,attr"`
			Vendor   string `xml:"vendor,attr"`
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
				} `xml:"service"`
				Script []struct {
					Text   string `xml:",chardata"`
					ID     string `xml:"id,attr"`
					Output string `xml:"output,attr"`
					Table  []struct {
						Text string `xml:",chardata"`
						Key  string `xml:"key,attr"`
						Elem []struct {
							Text string `xml:",chardata"`
							Key  string `xml:"key,attr"`
						} `xml:"elem"`
						Table []struct {
							Text string `xml:",chardata"`
							Key  string `xml:"key,attr"`
							Elem []struct {
								Text string `xml:",chardata"`
								Key  string `xml:"key,attr"`
							} `xml:"elem"`
							Table struct {
								Text string   `xml:",chardata"`
								Key  string   `xml:"key,attr"`
								Elem []string `xml:"elem"`
							} `xml:"table"`
						} `xml:"table"`
					} `xml:"table"`
					Elem []struct {
						Text string `xml:",chardata"`
						Key  string `xml:"key,attr"`
					} `xml:"elem"`
				} `xml:"script"`
			} `xml:"port"`
		} `xml:"ports"`
		Hostscript struct {
			Text   string `xml:",chardata"`
			Script []struct {
				Text   string `xml:",chardata"`
				ID     string `xml:"id,attr"`
				Output string `xml:"output,attr"`
				Table  struct {
					Text string `xml:",chardata"`
					Key  string `xml:"key,attr"`
					Elem string `xml:"elem"`
				} `xml:"table"`
				Elem []struct {
					Text string `xml:",chardata"`
					Key  string `xml:"key,attr"`
				} `xml:"elem"`
			} `xml:"script"`
		} `xml:"hostscript"`
		Times struct {
			Text   string `xml:",chardata"`
			Srtt   string `xml:"srtt,attr"`
			Rttvar string `xml:"rttvar,attr"`
			To     string `xml:"to,attr"`
		} `xml:"times"`
	} `xml:"host"`
	Postscript struct {
		Text   string `xml:",chardata"`
		Script struct {
			Text   string `xml:",chardata"`
			ID     string `xml:"id,attr"`
			Output string `xml:"output,attr"`
			Table  struct {
				Text string   `xml:",chardata"`
				Key  string   `xml:"key,attr"`
				Elem []string `xml:"elem"`
			} `xml:"table"`
		} `xml:"script"`
	} `xml:"postscript"`
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

func (n *NmapResult) NewFromXML(nmapXML []byte) NmapResult {
	var nmapRun NmapResult

	err := xml.Unmarshal(nmapXML, &nmapRun)
	if err != nil {
		return NmapResult{}
	}
	return nmapRun
}
