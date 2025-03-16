package cmd

import (
	"os"

	"github.com/mahbubzulkarnain/gn/internal/cmd"
)

func Execute() {
	os.Exit(cmd.Run(os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
}
