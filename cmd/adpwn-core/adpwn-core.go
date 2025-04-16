package main

import (
	"ADPwn-core/internal/rest"
	"ADPwn-core/internal/sse"
	"fmt"

	// To load modules
	_ "ADPwn-core/modulelib/attacks"
	_ "ADPwn-core/modulelib/enumeration"
)

func main() {
	fmt.Printf("\n....WELCOME TO....\n\n\n")
	fmt.Println(
		".----------------.  .----------------.  .----------------.  .----------------.  .-----------------.\n" +
			"| .--------------. || .--------------. || .--------------. || .--------------. || .--------------. | \n" +
			"| |      __      | || |  ________    | || |   ______     | || | _____  _____ | || | ____  _____  | | \n" +
			"| |     /  \\     | || | |_   ___ `.  | || |  |_   __ \\   | || ||_   _||_   _|| || ||_   \\|_   _| | | \n" +
			"| |    / /\\ \\    | || |   | |   `. \\ | || |    | |__) |  | || |  | | /\\ | |  | || |  |   \\ | |   | | \n" +
			"| |   / ____ \\   | || |   | |    | | | || |    |  ___/   | || |  | |/  \\| |  | || |  | |\\ \\| |   | |\n" +
			"| | _/ /    \\ \\_ | || |  _| |___.' / | || |   _| |_      | || |  |   /\\   |  | || | _| |_\\   |_  | | \n" +
			"| ||____|  |____|| || | |________.'  | || |  |_____|     | || |  |__/  \\__|  | || ||_____|\\____| | | \n" +
			"| |              | || |              | || |              | || |              | || |              | | \n" +
			"| '--------------' || '--------------' || '--------------' || '--------------' || '--------------' | \n" +
			"'----------------'  '----------------'  '----------------'  '----------------'  '----------------'")
	go sse.StartServer("8082")
	rest.StartServer("8081")

}
