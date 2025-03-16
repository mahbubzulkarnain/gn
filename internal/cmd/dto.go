package cmd

import (
	"path"

	"github.com/spf13/cobra"

	"github.com/mahbubzulkarnain/gn/internal/config"
	"github.com/mahbubzulkarnain/gn/internal/pkg/generator"
	"github.com/mahbubzulkarnain/gn/internal/pkg/gomod"
	"github.com/mahbubzulkarnain/gn/internal/pkg/str"
)

func init() {
	rootCMD.AddCommand(dtoCMD)
}

var dtoCMD = &cobra.Command{
	Use:   "dto",
	Short: `Create new dto`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		var entityName string
		if len(args) > 0 {
			entityName = args[0]
		} else {
			entityName = config.DTO().Name
		}

		config.App().ModuleName = gomod.Name()
		return dtoGenerator(dtoRequest{
			EntityNameSlug:     str.ToSlug(entityName),
			EntityNameLoweCase: str.ToLower(entityName),
		})
	},
}

type dtoRequest struct {
	EntityNameSlug     string
	EntityNameLoweCase string
	Version            string
}

func dtoGenerator(req dtoRequest) (err error) {
	if req.Version == "" {
		req.Version = "v1"
	}

	if err = generator.New(
		path.Join(config.App().Dir.Implment.DTO, req.EntityNameSlug, req.Version),
		path.Join(config.App().Dir.Template.DTO),
		map[string]interface{}{
			"EntityNameLoweCase": req.EntityNameLoweCase,
			"ModuleName":         config.App().ModuleName,
			"Version":            req.Version,
		},
	).Generate(); err != nil {
		return
	}
	return
}
