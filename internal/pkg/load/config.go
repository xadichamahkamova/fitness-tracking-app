package load

import (
	"github.com/xadichamahkamova/config-helper/function"
)

type Config struct {
	ServiceHost string `yaml:"service_host"`
	ServicePost string `yaml:"service_port"`
	Postgres    string `yaml:"postgres"`
}

func Load(path string) (*Config, error) {

	config := Config{}
	if err := function.LoadYAMLConfig("config/config.yaml", &config); err != nil {
		return nil, err
	}
	return &config, nil
}
