package configuration

import (
	"git.workshop21.ch/go/abraxas/configuration/yaml"
)

type Config struct {
	Env          string
	GRPCPort     string
	GatewayPort  string
	ImageBaseUrl string
	HttpsPort    int32
}

func ReadConfig() (*Config, error) {
	config := &Config{}
	err := yaml.ReadConfig(config,
		"./config/config.yaml")
	return config, err
}
