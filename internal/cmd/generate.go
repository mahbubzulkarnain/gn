package cmd

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/mahbubzulkarnain/gn/internal/pkg/generator"
	"github.com/spf13/cobra"

	"github.com/mahbubzulkarnain/gn/internal/config"
	"github.com/mahbubzulkarnain/gn/internal/pkg/str"
)

func init() {
	rootCMD.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: `Generate project from config gn.yaml`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		var (
			dir  string
			file = config.App().File
		)
		if f := cmd.Flag("file"); f != nil && f.Changed {
			if dir = f.Value.String(); dir == "" {
				return fmt.Errorf("file flag is required")
			}
			if dir, err = filepath.Abs(dir); err != nil {
				return err
			}
			file = filepath.Base(dir)
			dir = filepath.Dir(dir)
		} else {
			if dir, err = os.Getwd(); err != nil {
				return err
			}
		}

		var c config.Config
		if c, err = generateReadConfigFile(dir, file); err != nil {
			return
		}

		config.App().ModuleName = c.Module.Name

		var (
			entityNameSlug       string
			entityNameLowerCase  string
			entityNamePascalCase string

			tableName string
		)
		for _, entityConfig := range c.Module.Entities {
			entityNameSlug = str.ToSlug(entityConfig.Name)
			entityNameLowerCase = str.ToLower(entityConfig.Name)
			entityNamePascalCase = str.ToPascal(entityConfig.Name)

			if entityConfig.SQL.Name != "" {
				tableName = entityConfig.SQL.Name
			} else {
				tableName = entityNamePascalCase
			}

			// Repository
			if err = repositoryGenerator(repositoryRequest{
				EntityNameSlug:       entityNameSlug,
				EntityNameLoweCase:   entityNameLowerCase,
				EntityNamePascalCase: entityNamePascalCase,
				TableName:            tableName,
				Config:               entityConfig,
				Version:              entityConfig.Version,
			}); err != nil {
				return
			}

			// Service
			if err = serviceGenerator(serviceRequest{
				EntityNameSlug:       entityNameSlug,
				EntityNameLoweCase:   entityNameLowerCase,
				EntityNamePascalCase: entityNamePascalCase,
				Version:              entityConfig.Version,
			}); err != nil {
				return
			}

			switch c.Module.Framework.Name {
			case config.FrameworkEngineEcho:
				if err = generator.New(
					path.Join(config.App().Dir.Application.Delivery, entityNameSlug, entityConfig.Version),
					path.Join(config.App().Dir.Template.Delivery, `handler`),
					map[string]interface{}{
						"ModuleName":           config.App().ModuleName,
						"EntityNamePascalCase": entityNamePascalCase,
						"EntityNameLoweCase":   entityNameLowerCase,
						"Version":              entityConfig.Version,
					},
				).Generate(); err != nil {
					return
				}
			}
		}

		// Project
		if err = projectGenerator(projectRequest{
			File:   file,
			Config: c,
		}); err != nil {
			return
		}
		return
	},
}

func generateReadConfigFile(dir, name string) (c config.Config, err error) {
	var configPath string
	if name != "" {
		configPath = filepath.Join(dir, name)
	} else {
		var (
			filePaths []string

			fileJSONPath = filepath.Join(dir, `gn.json`)
			fileYAMLPath = filepath.Join(dir, `gn.yaml`)
			fileYMLPath  = filepath.Join(dir, `gn.yml`)
		)

		if _, err = os.Stat(fileJSONPath); !os.IsNotExist(err) {
			filePaths = append(filePaths, fileJSONPath)
		}
		if _, err = os.Stat(fileYAMLPath); !os.IsNotExist(err) {
			filePaths = append(filePaths, fileYAMLPath)
		}
		if _, err = os.Stat(fileYMLPath); !os.IsNotExist(err) {
			filePaths = append(filePaths, fileYMLPath)
		}

		filePathsLen := len(filePaths)
		if filePathsLen == 0 {
			err = errors.New(`file config not found. ( gn.json | gn.yaml | gn.yml )`)
			return
		} else if filePathsLen > 1 {
			err = fmt.Errorf(`which file ? ( %v ) `, strings.Join(filePaths, " | "))
			return
		}
		configPath = filePaths[0]
	}

	var file *os.File
	if file, err = os.Open(configPath); err != nil {
		return
	}
	defer file.Close()

	if c, err = config.ConfigParse(file); err != nil {
		return
	}
	return

}
