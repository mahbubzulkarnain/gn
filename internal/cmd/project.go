package cmd

import (
	"fmt"
	"path"
	"strings"

	"github.com/spf13/cobra"

	"github.com/mahbubzulkarnain/gn/internal/config"
	"github.com/mahbubzulkarnain/gn/internal/pkg/generator"
)

func init() {
	rootCMD.AddCommand(projectCMD)
}

var projectCMD = &cobra.Command{
	Use:   `new`,
	Short: `Create new project`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if len(args) == 0 {
			return cmd.Usage()
		}

		config.App().ModuleName = args[0]
		return projectGenerator(projectRequest{
			File: config.App().File,
			Config: config.Config{
				Module: config.ModuleConfig{
					Name: config.App().ModuleName,
				},
			},
		})
	},
}

type projectRequest struct {
	File   string
	Config config.Config
}

func projectGenerator(req projectRequest) (err error) {
	if err = initGenerator(initRequest{
		File:   req.File,
		Config: req.Config,
	}); err != nil {
		return
	}

	switch req.Config.Module.Framework.Name {
	case config.FrameworkEngineEcho:
		config.App().PKGListAdd(`github.com/hiko1129/echo-pprof@v1.0.1`)
		config.App().PKGListAdd(`github.com/joho/godotenv@v1.5.1`)
		config.App().PKGListAdd(`github.com/labstack/echo/v4@v4.12.0`)

		if err = generator.New(
			`.env`,
			path.Join(config.App().Dir.Template.Delivery, `.env.tmpl`),
			nil,
		).Generate(); err != nil {
			return
		}

		if err = generator.New(
			path.Join(config.App().Dir.Implment.Delivery, `http.go`),
			path.Join(config.App().Dir.Template.Delivery, `server.go.tmpl`),
			map[string]interface{}{
				"ModuleName": config.App().ModuleName,
			},
		).Generate(); err != nil {
			return
		}

		if err = generator.New(
			`main.go`,
			path.Join(config.App().Dir.Template.Path, `main.go.tmpl`),
			map[string]interface{}{
				`ModuleName`: config.App().ModuleName,
				`ImportList`: strings.Join([]string{
					path.Join(config.App().ModuleName, config.App().Dir.Implment.Delivery),
				}, "\n"),
				`FuncBody`: `delivery.HTTP(nil)`,
			},
		).Generate(); err != nil {
			return
		}
	}

	if err = generator.New(
		path.Join(config.App().Dir.Implment.Config),
		path.Join(config.App().Dir.Template.Config),
		nil,
	).Generate(); err != nil {
		return
	}

	if err = generator.New(
		path.Join(config.App().Dir.Implment.Repository, "repository.go"),
		path.Join(config.App().Dir.Template.Path, "internal", `repository`, "repository.go.tmpl"),
		map[string]interface{}{
			"ModuleName": config.App().ModuleName,
		},
	).Generate(); err != nil {
		return
	}

	if err = generator.New(
		path.Join(config.App().Dir.Implment.Service, "service.go"),
		path.Join(config.App().Dir.Template.Path, "internal", `service`, "service.go.tmpl"),
		map[string]interface{}{
			"ModuleName": config.App().ModuleName,
		},
	).Generate(); err != nil {
		return
	}

	if config.App().UseSQL {
		if err = generator.New(
			path.Join("pkg", "database"),
			path.Join(config.App().Dir.Template.Path, "pkg", `database`),
			nil,
		).Generate(); err != nil {
			return
		}
	}

	var pkgList string
	if len(config.App().PKGList()) > 0 {
		pkgList = fmt.Sprintf("require (\n\t%v\n)", strings.Join(config.App().PKGList(), "\n\t"))
	}

	if err = generator.New(
		`go.mod`,
		path.Join(config.App().Dir.Template.Path, `go.mod.tmpl`),
		map[string]interface{}{
			"ModuleName": config.App().ModuleName,
			"PKGList":    pkgList,
		},
	).Generate(); err != nil {
		return
	}
	return
}
