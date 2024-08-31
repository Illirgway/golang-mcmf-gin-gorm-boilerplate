/**
 *
 * Copyright (c) 2024 Illirgway
 *
 */

package config

type Config struct {
	Listen   string `json:"listen"`
	Port     uint   `json:"port"`
	DB       string `json:"db"` // DSN
	Sessions struct {
		Secret  string `json:"secret"` // cookie (sessions) secret
		Timeout uint   `json:"timeout"`
	} `json:"sessions"`
	Paths struct {
		Templates string `json:"templates"`
		Assets    string `json:"assets"`
	} `json:"paths"`
	Debug bool `json:"debug"`
}

func New() (c *Config, err error) {

	c = new(Config)

	if err = c.init(); err != nil {
		return nil, err
	}

	return c, nil
}
