package adapter

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
)

type NetExec struct {
	string Result
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
	// here we perform the pwd command.
	// we can store the output of this in our out variable
	// and catch any errors in err
	out, err := exec.Command("nxc", "smb").Output()

	// if there is an error with our execution
	// handle it here
	if err != nil {
		fmt.Printf("%s", err)
	}
	// as the out variable defined above is of type []byte we need to convert
	// this to a string or else we will see garbage printed out in our console
	// this is how we convert it to a string
	fmt.Println("Command Successfully Executed")
	output := string(out[:])
	fmt.Println(output)

	/* // let's try the pwd command herer
	out, err = exec.Command("pwd").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println("Command Successfully Executed")
	output = string(out[:])
	fmt.Println(output) */
	e, err := os.Executable()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(path.Dir(e))

}
