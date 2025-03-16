package config

type ModuleConfig struct {
	Name      string          `json:"name" yaml:"name"`
	Framework FrameworkConfig `json:"framework" yaml:"framework"`
	Entities  []EntityConfig  `json:"entities" yaml:"entities"`
}
