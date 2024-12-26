package configs

import "github.com/kelseyhightower/envconfig"

type ServerConfig struct {
	Port     string `envconfig:"PORT"`
	Hostname string `envconfig:"HOSTNAME"`
}

func CreateServerConfig() (*ServerConfig, error) {
	config := &ServerConfig{}
	err := envconfig.Process("", config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
