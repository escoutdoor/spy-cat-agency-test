package env

import (
	"github.com/caarlos0/env/v11"
)

type catClientConfig struct {
	ClientApiKey string `env:"CAT_CLIENT_API_KEY,required"`
}

func NewCatClientConfig() (*catClientConfig, error) {
	config := new(catClientConfig)
	if err := env.Parse(config); err != nil {
		return nil, err
	}

	return config, nil
}

func (c *catClientConfig) ApiKey() string {
	return c.ClientApiKey
}
