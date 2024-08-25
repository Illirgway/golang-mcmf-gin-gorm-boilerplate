/**
 *
 * Copyright (c) 2024 Illirgway
 *
 */

package config

type Config struct {
	Listen string `json:"listen"`
	Port   uint   `json:"port"`
	DB     string `json:"db"`     // DSN
	Secret string `json:"secret"` // cookie (sessions) secret
	Dirs   struct {
		Templates string `json:"templates"`
		Assets    string `json:"assets"`
	} `json:"dirs"`
	Debug bool `json:"debug"`
}

func New() (c *Config, err error) {

	c = new(Config)

	if err = c.init(); err != nil {
		return nil, err
	}

	return c, nil
}
