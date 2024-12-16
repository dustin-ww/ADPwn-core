package worker

import (
	"ADPwn/tools/tools"
	"fmt"
)

type Enumerator struct {
}

func (e Enumerator) Run() {
	fmt.Println("STARTING enumeration")
	nmap := tools.Nmap{}

	nmap.Execute()
}
