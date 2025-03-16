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

type AppConfigDirImplement struct {
	Path       string
	Config     string
	DTO        string
	Delivery   string
	Domain     string
	Repository string
	Service    string
}

type AppConfigDirTemplate struct {
	Path       string
	Config     string
	DTO        string
	Delivery   string
	Domain     string
	Repository string
	Service    string
}

type AppConfigDir struct {
	Internal string
	Implment AppConfigDirImplement
	Template AppConfigDirTemplate
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

		appConfig.Dir.Implment.Path = path.Join("internal")
		appConfig.Dir.Implment.Config = path.Join("internal", "config")
		appConfig.Dir.Implment.DTO = path.Join("internal", "dto")
		appConfig.Dir.Implment.Delivery = path.Join("internal", "delivery")
		appConfig.Dir.Implment.Domain = path.Join("internal", "domain")
		appConfig.Dir.Implment.Repository = path.Join("internal", "repository")
		appConfig.Dir.Implment.Service = path.Join("internal", "service")

		appConfig.Dir.Template.Path = path.Join(appConfig.Dir.Internal, "templates")
		appConfig.Dir.Template.Config = path.Join(appConfig.Dir.Template.Path, "internal", "config")
		appConfig.Dir.Template.DTO = path.Join(appConfig.Dir.Template.Path, "internal", "dto", "entity")
		appConfig.Dir.Template.Delivery = path.Join(appConfig.Dir.Template.Path, "internal", "delivery", "echo")
		appConfig.Dir.Template.Domain = path.Join(appConfig.Dir.Template.Path, "internal", "domain", "entity")
		appConfig.Dir.Template.Repository = path.Join(appConfig.Dir.Template.Path, "internal", "repository", "entity")
		appConfig.Dir.Template.Service = path.Join(appConfig.Dir.Template.Path, "internal", "service", "entity")
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
