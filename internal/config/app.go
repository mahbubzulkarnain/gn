package config

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"sync"
)

type AppConfigPKGListMap struct {
	Name    string
	Version string
}

type AppConfigDirFunc func(entitySlug, entityVersion string) string

type AppConfigDirApplication struct {
	Delivery   string
	Config     string
	Repository string
	Service    string
}

type AppConfigDirPKG struct {
	Path       string
	DTO        AppConfigDirFunc
	Domain     AppConfigDirFunc
	Entity     AppConfigDirFunc
	Repository AppConfigDirFunc
	Service    AppConfigDirFunc
}

type AppConfigDirTemplate struct {
	Path       string
	Config     string
	DTO        string
	Domain     string
	Delivery   string
	Entity     string
	Repository string
	Service    string
}

type AppConfigDir struct {
	Internal    string
	Application AppConfigDirApplication
	PKG         AppConfigDirPKG
	Template    AppConfigDirTemplate
}

type AppConfig struct {
	ModuleName string
	Dir        AppConfigDir
	File       string
	UseSQL     bool

	pkgList    []string
	pkgListMap map[string]*AppConfigPKGListMap
}

var (
	appConfig *AppConfig
	appOnce   sync.Once
)

func App() *AppConfig {
	appOnce.Do(func() {
		appConfig = new(AppConfig)
		appConfig.File = `gn.yaml`

		var (
			_, b, _, _ = runtime.Caller(0)
			dir        = strings.Split(path.Dir(b), "/")
		)
		for i := len(dir) - 1; i > 0; i-- {
			if dir[i] == "internal" {
				appConfig.Dir.Internal = strings.TrimSuffix(path.Dir(b), path.Join(dir[i+1:]...))
				break
			}
		}

		appConfig.Dir.Template.Path = path.Join(appConfig.Dir.Internal, "templates")
		appConfig.Dir.Template.Config = path.Join(appConfig.Dir.Template.Path, "internal", "config")
		appConfig.Dir.Template.DTO = path.Join(appConfig.Dir.Template.Path, "internal", "dto", "entity")
		appConfig.Dir.Template.Domain = path.Join(appConfig.Dir.Template.Path, "internal", "domain", "entity")
		appConfig.Dir.Template.Delivery = path.Join(appConfig.Dir.Template.Path, "internal", "delivery", "echo")
		appConfig.Dir.Template.Entity = path.Join(appConfig.Dir.Template.Path, "pkg", "entity")
		appConfig.Dir.Template.Repository = path.Join(appConfig.Dir.Template.Path, "internal", "repository", "entity")
		appConfig.Dir.Template.Service = path.Join(appConfig.Dir.Template.Path, "internal", "service", "entity")

		appConfig.Dir.Application.Delivery = path.Join("internal", "delivery")
		appConfig.Dir.Application.Config = path.Join("internal", "config")
		appConfig.Dir.Application.Repository = path.Join("internal", "repository")
		appConfig.Dir.Application.Service = path.Join("internal", "service")

		appConfig.Dir.PKG.Path = path.Join("pkg")
		appConfig.Dir.PKG.DTO = func(entitySlug, entityVersion string) string {
			return path.Join(appConfig.Dir.PKG.Path, entitySlug, entityVersion, "dto")
		}
		appConfig.Dir.PKG.Repository = func(entitySlug, entityVersion string) string {
			return path.Join(appConfig.Dir.PKG.Path, entitySlug, entityVersion, "repository")
		}

		appConfig.Dir.PKG.Domain = func(entitySlug, entityVersion string) string {
			return path.Join(appConfig.Dir.PKG.Path, entitySlug, entityVersion, "domain")
		}
		appConfig.Dir.PKG.Entity = func(entitySlug, entityVersion string) string {
			return path.Join(appConfig.Dir.PKG.Path, entitySlug, entityVersion)
		}
		appConfig.Dir.PKG.Service = func(entitySlug, entityVersion string) string {
			return path.Join(appConfig.Dir.PKG.Path, entitySlug, entityVersion, "service")
		}
	})
	return appConfig
}

// Close ...
func (c *AppConfig) Close() {
	c.PKGListClear()
	c.UseSQL = false
}

// PKGListAdd ...
func (c *AppConfig) PKGListAdd(i string) {
	i = strings.TrimSpace(i)
	var iSplit []string
	for _, iSplitBy := range []string{`@`, ` `} {
		if strings.Contains(i, iSplitBy) {
			iSplit = strings.Split(i, iSplitBy)
		}
	}

	var name, version string
	switch len(iSplit) {
	case 2:
		name, version = iSplit[0], iSplit[1]
	default:
		return
	}

	if c.pkgListMap == nil {
		c.pkgListMap = make(map[string]*AppConfigPKGListMap)
	}

	if c.pkgListMap[name] != nil {
		if c.pkgListMap[name].Version == version {
			return
		}
		c.pkgListMap[name].Version = version
		for _, v := range c.pkgListMap {
			c.pkgList = append(c.pkgList, fmt.Sprintf(`%v %v`, v.Name, v.Version))
		}
		return
	}
	c.pkgListMap[name] = &AppConfigPKGListMap{
		Name:    name,
		Version: version,
	}
	c.pkgList = append(c.pkgList, fmt.Sprintf(`%v %v`, name, version))
	return
}

// PKGListClear ...
func (c *AppConfig) PKGListClear() {
	c.pkgListMap = make(map[string]*AppConfigPKGListMap)
	c.pkgList = make([]string, 0)
}

// PKGList ...
func (c *AppConfig) PKGList() []string {
	return c.pkgList
}
