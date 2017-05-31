package toutput

import (
	"fmt"
	"os"
)

// PrintFatal print msg and die
func PrintFatal(err string, msg string) {
	if msg != "" {
		fmt.Fprintf(os.Stderr, "%s\n", msg)
	}
	if err != "" {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	os.Exit(1)
}

// CheckError print msg and die
func CheckError(e error, msg string, fail bool) bool {
	if e != nil {
		if fail {
			PrintFatal(e.Error(), msg)
		} else {
			StderrLog(msg, 0, 0)
			fmt.Fprintf(os.Stderr, "%s\n", e.Error())
			return true
		}
	}
	return false
}
