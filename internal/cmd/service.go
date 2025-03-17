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
	rootCMD.AddCommand(serviceCMD)
}

var serviceCMD = &cobra.Command{
	Use:   "service",
	Short: `Create new service`,
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
		)

		config.App().ModuleName = gomod.Name()
		return serviceGenerator(serviceRequest{
			EntityNamePascalCase: entityNamePascalCase,
			EntityNameLoweCase:   entityNameLowerCase,
			EntityNameSlug:       entityNameSlug,
		})
	},
}

type serviceRequest struct {
	EntityNamePascalCase string
	EntityNameLoweCase   string
	EntityNameSlug       string
	Version              string
}

func serviceGenerator(req serviceRequest) (err error) {
	if req.Version == "" {
		req.Version = "v1"
	}

	if err = dtoGenerator(dtoRequest{
		EntityNameSlug:     req.EntityNameSlug,
		EntityNameLoweCase: req.EntityNameLoweCase,
		Version:            req.Version,
	}); err != nil {
		return
	}

	if err = generator.New(
		path.Join(config.App().Dir.PKG.Service(req.EntityNameSlug, req.Version)),
		path.Join(config.App().Dir.Template.Service),
		map[string]interface{}{
			"EntityNamePascalCase": req.EntityNamePascalCase,
			"EntityNameLoweCase":   req.EntityNameLoweCase,
			"ModuleName":           config.App().ModuleName,
			"Version":              req.Version,
		},
	).Generate(); err != nil {
		return
	}

	if err = generator.New(
		path.Join(config.App().Dir.PKG.Entity(req.EntityNameSlug, req.Version), "service.go"),
		path.Join(config.App().Dir.Template.Entity, `service.go.tmpl`),
		map[string]interface{}{
			"EntityNamePascalCase": req.EntityNamePascalCase,
			"EntityNameLoweCase":   req.EntityNameLoweCase,
			"ModuleName":           config.App().ModuleName,
			"Version":              req.Version,
		},
	).Generate(); err != nil {
		return
	}
	return
}
