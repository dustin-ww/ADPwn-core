package adapter

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
)

type NetExec struct {
}

func (worker NetExec) testAnonymousLogin() {
}

func (worker NetExec) Execute() {
	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	} else {
		log.Println(dir)
		fmt.Println(dir)
	}

	out, err := exec.Command("nxc", "smb").Output()

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
