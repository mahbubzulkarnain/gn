package config

import (
	"io"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Module  ModuleConfig         `json:"module" yaml:"module"`
	Feature FeatureConfig        `json:"features,omitempty" yaml:"features,omitempty"`
	PKG     []string             `json:"pkg" yaml:"pkg"`
	Options map[string]yaml.Node `json:"options,omitempty" yaml:"options,omitempty"`
}

func (c Config) Validate() (err error) {
	return
}

func ConfigParse(rd io.Reader) (c Config, err error) {
	dec := yaml.NewDecoder(rd)
	dec.KnownFields(true)
	if err = dec.Decode(&c); err != nil {
		return
	}
	if err = c.Validate(); err != nil {
		return
	}
	return
}
