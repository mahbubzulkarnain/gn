package config

import "sync"

type DTOConfig struct {
	Name string
}

var (
	dtoConfig *DTOConfig
	dtoOnce   sync.Once
)

func DTO() *DTOConfig {
	dtoOnce.Do(func() {
		dtoConfig = new(DTOConfig)
		dtoConfig.Name = `Entity`
	})
	return dtoConfig
}
