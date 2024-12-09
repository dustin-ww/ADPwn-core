package cmdutils

import (
	tm "github.com/buger/goterm"
)

func ClearCMD() {
	tm.Clear()
	tm.MoveCursor(1, 1)
	tm.Flush()
}
