package upload

import (
	"fmt"
	"os"
)

/**
 *
 */
func ExitErrorf(msg string, args ...interface{}) {
	if args[0] != nil {
		fmt.Fprintf(os.Stderr, msg+"\n", args...)
		os.Exit(1)
	}
}
