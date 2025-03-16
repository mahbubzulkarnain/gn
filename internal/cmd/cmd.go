package cmd

import (
	"context"
	"errors"
	"io"
	"os/exec"

	"github.com/spf13/cobra"

	"github.com/mahbubzulkarnain/gn/internal/config"
)

var rootCMD = cobra.Command{Use: `gn`, SilenceUsage: true}

func Run(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) int {
	defer config.App().PKGListClear()

	rootCMD.PersistentFlags().StringP("file", "f", "", `default: gn.yaml`)
	rootCMD.PersistentFlags().String(`table_name`, ``, `Table name (example: Entities)`)

	rootCMD.SetArgs(args)
	rootCMD.SetIn(stdin)
	rootCMD.SetOut(stdout)
	rootCMD.SetErr(stderr)

	ctx := context.Background()
	if err := rootCMD.ExecuteContext(ctx); err != nil {
		var errExit *exec.ExitError
		if errors.As(err, &errExit) {
			return errExit.ExitCode()
		}
		return 1
	}
	return 0
}
