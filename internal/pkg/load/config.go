package load

import (
	"github.com/xadichamahkamova/config-helper/function"
)

type Config struct {
	ServiceHost   string `yaml:"service_host"`
	ServicePost   string `yaml:"service_port"`
	Postgres      string `yaml:"postgres"`
	TokenKey      string `yaml:"token_key"`
	EmailFrom     string `yaml:"my_email"`
	EmailPassword string `yaml:"my_password"`
}

func Load(path string) (*Config, error) {

	config := Config{}
	if err := function.LoadYAMLConfig("config/config.yaml", &config); err != nil {
		return nil, err
	}
	return &config, nil
}
