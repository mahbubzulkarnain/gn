package config

import "sync"

type EntityConfig struct {
	Name     string    `json:"name" yaml:"name"`
	Version  string    `json:"version" yaml:"version"`
	Endpoint *string   `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`
	SQL      SQLConfig `json:"sql" yaml:"sql"`
}

var (
	entityConfig *EntityConfig
	entityOnce   sync.Once
)

func Entity() *EntityConfig {
	entityOnce.Do(func() {
		entityConfig = new(EntityConfig)
		entityConfig.Name = "Entity"
	})
	return entityConfig
}
