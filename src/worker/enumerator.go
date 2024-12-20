package worker

import (
	"ADPwn/database/project/model"
	"ADPwn/tools/tools"
	"fmt"
)

type Enumerator struct {
}

func (e Enumerator) Run(project model.Project) {
	fmt.Println("STARTING enumeration")

	// nmap for initial recon
	nmap := tools.Nmap{}
	nmap.Execute(project)
}
