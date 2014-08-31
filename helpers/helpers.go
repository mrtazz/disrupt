package helpers

import (
	"fmt"
	"os"
)

var Debug bool = false

// Helper function to enable debug output if the variables is set
func PrintIfDebug(format string, args ...interface{}) {
	if Debug {
		fmt.Fprintf(os.Stderr, format, args...)
	}
}
