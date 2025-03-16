package cmd

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"

	"github.com/spf13/cobra"

	"github.com/mahbubzulkarnain/gn/internal/config"
)

func init() {
	rootCMD.AddCommand(initCMD)
}

var initCMD = &cobra.Command{
	Use:   "init",
	Short: `Init project`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		var (
			c          config.Config
			ModuleName string
		)
		if len(args) > 0 {
			ModuleName = args[0]
			c.Module.Name = ModuleName
		}

		file := config.App().File
		if f := cmd.Flag("file"); f != nil && f.Changed {
			if file = f.Value.String(); file == "" {
				return fmt.Errorf("file flag is required")
			}
		}

		if _, err = os.Stat(file); !os.IsNotExist(err) {
			fmt.Printf("file %s already exists", file)
			return nil
		}

		return initGenerator(initRequest{
			File:   file,
			Config: c,
		})
	},
}

type initRequest struct {
	File   string
	Config config.Config
}

func initGenerator(req initRequest) (err error) {
	if req.File == "" {
		err = fmt.Errorf("file flag is required")
		return
	}

	if _, err = os.Stat(req.File); os.IsNotExist(err) {
		var blob []byte
		if blob, err = yaml.Marshal(req.Config); err != nil {
			return
		}

		if err = os.WriteFile(req.File, blob, 0644); err != nil {
			return
		}

		fmt.Printf("file %s created", req.File)
	}
	return
}
