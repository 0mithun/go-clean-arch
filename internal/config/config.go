package config

import (
	"fmt"

	"github.com/kkyr/fig"
)

type Config struct {
	Mongo Mongo `fig:"mongo" validate:"required"`
	Http  Http  `fig:"http" validate:"required"`
}

type Mongo struct {
	AppName     string `fig:"app_name" validate:"required"`
	Host        string `fig:"host" validate:"required"`
	Port        string `fig:"port" validate:"required"`
	Username    string `fig:"username"`
	Password    string `fig:"password"`
	Params      string `fig:"params" `
	Database    string `fig:"database" validate:"required"`
	MinPoolSize int    `fig:"min_pool_size" validate:"required"`
	MaxPoolSize int    `fig:"max_pool_size" validate:"required"`
}

func (c *Mongo) GetConnectionString() string {
	result := fmt.Sprintf("mongodb://%s:%s", c.Host, c.Port)

	if c.Username != "" && c.Password != "" {
		result = fmt.Sprintf("mongodb://%s:%s@%s:%s", c.Username, c.Password, c.Host, c.Port)
	}

	if c.Params != "" {
		result = fmt.Sprintf("%s/%s", result, c.Params)
	}

	return result
}

type Http struct {
	Port string `fig:"port" validate:"required"`
}

func GetConfig() (*Config, error) {
	var config Config
	if err := fig.Load(&config, fig.UseEnv("")); err != nil {
		return nil, fmt.Errorf("error loading config: %v", err)
	}

	return &config, nil
}
