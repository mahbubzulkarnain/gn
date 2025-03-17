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
	rootCMD.AddCommand(repositoryCMD)
}

var repositoryCMD = &cobra.Command{
	Use:     "repository",
	Short:   `Create new repository`,
	Example: `gn repository Users`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		var entityName string
		if len(args) > 0 {
			entityName = args[0]
		} else {
			entityName = config.Entity().Name
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
		return repositoryGenerator(repositoryRequest{
			EntityNameSlug:       entityNameSlug,
			EntityNameLoweCase:   entityNameLowerCase,
			EntityNamePascalCase: entityNamePascalCase,
			TableName:            tableName,
		})
	},
}

type repositoryRequest struct {
	EntityNameSlug       string
	EntityNameLoweCase   string
	EntityNamePascalCase string
	TableName            string
	Version              string
	Config               config.EntityConfig
}

func repositoryGenerator(req repositoryRequest) (err error) {
	if req.Version == "" {
		req.Version = "v1"
	}

	switch req.Config.SQL.Engine {
	case config.SQLEngineMySQL:
		config.App().UseSQL = true
	case config.SQLEnginePostgreSQL:
		config.App().UseSQL = true
		config.App().PKGListAdd(`gorm.io/driver/postgres@v1.5.9`)
		config.App().PKGListAdd(`github.com/lib/pq@v1.10.2`)
	case config.SQLEngineSQLite:
		config.App().UseSQL = true
	}
	if config.App().UseSQL {
		config.App().PKGListAdd(`go.elastic.co/apm/module/apmgormv2/v2@v2.6.2`)
		config.App().PKGListAdd(`gorm.io/plugin/dbresolver@v1.2.1`)
		config.App().PKGListAdd(`gorm.io/gorm@v1.23.4`)
	}

	if err = generator.New(
		path.Join(config.App().Dir.PKG.Repository(req.EntityNameSlug, req.Version)),
		path.Join(config.App().Dir.Template.Repository),
		map[string]interface{}{
			"EntityNameLoweCase":   req.EntityNameLoweCase,
			"EntityNamePascalCase": req.EntityNamePascalCase,
			"ModuleName":           config.App().ModuleName,
			"TableName":            req.TableName,
			"Version":              req.Version,
		},
	).Generate(); err != nil {
		return
	}

	if err = generator.New(
		path.Join(config.App().Dir.PKG.Entity(req.EntityNameSlug, req.Version), "repository.go"),
		path.Join(config.App().Dir.Template.Entity, `repository.go.tmpl`),
		map[string]interface{}{
			"EntityNamePascalCase": req.EntityNamePascalCase,
			"EntityNameLoweCase":   req.EntityNameLoweCase,
			"ModuleName":           config.App().ModuleName,
			"Version":              req.Version,
			"TableName":            req.TableName,
		},
	).Generate(); err != nil {
		return
	}

	return
}
