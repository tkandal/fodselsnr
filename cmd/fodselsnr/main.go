package main

import (
	"errors"
	"fmt"
	"github.com/tkandal/fodselsnr"
	"os"
	"path/filepath"
	"strings"
)

var (
	ErrIllegalArgument = errors.New("illegal argument")
	ErrIllegalNIN      = errors.New("nin is illegal")
)

func main() {
	if err := realMain(); err != nil {
		os.Exit(1)
	}
}

func realMain() error {
	if len(os.Args) < 2 || len(os.Args) > 2 {
		_, _ = fmt.Fprintf(os.Stdout, "Usage: %s fnr\n", filepath.Base(os.Args[0]))
		return ErrIllegalArgument
	}

	arg := strings.TrimSpace(os.Args[1])
	if len(arg) == 0 {
		_, _ = fmt.Fprintf(os.Stdout, "fødselsnummer is empty\n")
		return ErrIllegalArgument
	}

	if len(arg) < (fodselsnr.FodselsnrLength-1) || len(arg) > fodselsnr.FodselsnrLength {
		_, _ = fmt.Fprintf(os.Stdout, "fødselsnummer has incorrect length\n")
		return ErrIllegalArgument
	}

	if !fodselsnr.Check(arg) {
		_, _ = fmt.Fprintf(os.Stdout, "fødselsnummer %s is not legal\n", arg)
		return ErrIllegalNIN
	}
	_, _ = fmt.Fprintf(os.Stdout, "fødselsnummer %s is legal\n", arg)
	return nil
}
