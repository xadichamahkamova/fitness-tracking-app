package load

import (
	"github.com/xadichamahkamova/config-helper/function"
)

type PostgresConfig struct {
	Host     string
	Port     string
	Password string
	Database string
}

type ServiceConfig struct {
	Host string
	Port string
}

type Config struct {
	Service  ServiceConfig
	Postgres PostgresConfig
}

func Load(path string) (*Config, error) {

	config := Config{}
	if err := function.LoadYAMLConfig("config/config.yaml", &config); err != nil {
		return nil, err
	}
	return &config, nil
}
