package main

import (
	"fmt"
	"github.com/tkandal/fodselsnr"
	"os"
	"strings"
)

func main() {
	arg := strings.TrimSpace(os.Args[1])
	if len(arg) == 0 {
		_, _ = fmt.Fprintf(os.Stdout, "fødselsnummer is empty\n")
		os.Exit(1)
	}
	if len(arg) < (fodselsnr.FodselsnrLength-1) || len(arg) > fodselsnr.FodselsnrLength {
		_, _ = fmt.Fprintf(os.Stdout, "fødselsnummer has incorrect length\n")
		os.Exit(1)
	}

	if !fodselsnr.Check(arg) {
		_, _ = fmt.Fprintf(os.Stdout, "fødselsnummer %s is not legal\n", arg)
		os.Exit(1)
	}
	_, _ = fmt.Fprintf(os.Stdout, "fødselsnummer %s is legal\n", arg)
	os.Exit(0)
}
