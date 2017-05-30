package toutput

import (
	"fmt"
	"os"
)

// TVerbosity verbosity level
var TVerbosity = map[string]int{
	"ERROR":   0,
	"WARNING": 1,
	"INFO":    2,
	"DEBUG":   3,
}

// StderrLog print errors to stderr
func StderrLog(msg string, UserSelectedLevel int, printLevel int) {
	var prefix string
	if printLevel == TVerbosity["INFO"] {
		prefix = "I> "
	} else if printLevel == TVerbosity["DEBUG"] {
		prefix = "D> "
	} else {
		prefix = ""
	}
	if UserSelectedLevel >= printLevel {
		fmt.Fprintf(os.Stderr, "%s%s\n", prefix, msg)
	}
}
