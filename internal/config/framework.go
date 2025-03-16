package config

type (
	FrameworkEngine string
	FrameworkConfig struct {
		Name FrameworkEngine `json:"name" yaml:"name"`
	}
)

const (
	FrameworkEngineEcho FrameworkEngine = "echo"
)
