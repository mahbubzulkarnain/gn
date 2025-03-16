package config

type SQLEngine string

const (
	SQLEngineMySQL      SQLEngine = "mysql"
	SQLEnginePostgreSQL SQLEngine = "postgresql"
	SQLEngineSQLite     SQLEngine = "sqlite"
)

type SQLConfigDatabase struct {
	URI     string `json:"uri" yaml:"uri"`
	Managed bool   `json:"managed" yaml:"managed"`
}

type SQLConfig struct {
	Name     string             `json:"name" yaml:"name"`
	Engine   SQLEngine          `json:"engine" yaml:"engine"`
	Schema   PathsConfig        `json:"schema,omitempty" yaml:"schema,omitempty"`
	Database *SQLConfigDatabase `json:"database,omitempty" yaml:"database,omitempty"`
}
