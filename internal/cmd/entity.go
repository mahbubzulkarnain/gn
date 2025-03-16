package cmd

import (
	"github.com/spf13/cobra"

	"github.com/mahbubzulkarnain/gn/internal/config"
	"github.com/mahbubzulkarnain/gn/internal/pkg/gomod"
	"github.com/mahbubzulkarnain/gn/internal/pkg/str"
)

func init() {
	rootCMD.AddCommand(entityCMD)
}

var entityCMD = &cobra.Command{
	Use:   "entity",
	Short: `Create new entity`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		var entityName string
		if entityName, err = cmd.Flags().GetString("name"); err != nil {
			return
		}

		var (
			entityNameSlug       = str.ToSlug(entityName)
			entityNameLowerCase  = str.ToLower(entityName)
			entityNamePascalCase = str.ToPascal(entityName)

			tableName string
		)
		if tableName, err = cmd.Flags().GetString("table_name"); err != nil {
			return
		}

		if tableName == "" {
			tableName = entityNameLowerCase
		}

		config.App().ModuleName = gomod.Name()
		return entityGenerator(entityRequest{
			EntityNameSlug:       entityNameSlug,
			EntityNameLoweCase:   entityNameLowerCase,
			EntityNamePascalCase: entityNamePascalCase,
			ModuleName:           gomod.Name(),
			TableName:            tableName,
		})
	},
}

type entityRequest struct {
	EntityNameSlug       string
	EntityNameLoweCase   string
	EntityNamePascalCase string
	TableName            string
	ModuleName           string
}

func entityGenerator(req entityRequest) (err error) {
	if err = repositoryGenerator(repositoryRequest{
		EntityNameSlug:       req.EntityNameSlug,
		EntityNameLoweCase:   req.EntityNameLoweCase,
		EntityNamePascalCase: req.EntityNamePascalCase,
		TableName:            req.TableName,
	}); err != nil {
		return
	}

	if err = serviceGenerator(serviceRequest{
		EntityNameSlug:       req.EntityNameSlug,
		EntityNameLoweCase:   req.EntityNameLoweCase,
		EntityNamePascalCase: req.EntityNamePascalCase,
	}); err != nil {
		return
	}

	return
}
