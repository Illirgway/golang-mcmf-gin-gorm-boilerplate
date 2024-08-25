/**
 *
 * Copyright (c) 2024 Illirgway
 *
 */

package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/alexflint/go-arg"
)

var (
	description = "backend REST API server"
)

type CliArgs struct {
	Config string `arg:"-c,--config" placeholder:"PATH" help:"path to config file"`
}

func (cli *CliArgs) validate() (err error) {

	fp, err := os.Open(cli.Config)

	if err != nil {
		return fmt.Errorf("config file open error: %w", err)
	}

	return fp.Close()
}

func (c *Config) init() (err error) {

	cli := new(CliArgs)

	arg.MustParse(cli)

	if err = cli.validate(); err != nil {
		return err
	}

	return c.load(cli.Config)
}

func (c *Config) validate() error {
	// TODO validate config values
	return nil
}

// Description implements arg.Described
func (c *Config) Description() string {
	return description
}

func (c *Config) load(cfg string) error {

	data, err := os.ReadFile(cfg)

	if err != nil {
		return fmt.Errorf("config file %q read error: %w", cfg, err)
	}

	if err = json.Unmarshal(data, c); err != nil {
		return fmt.Errorf("config file %q parse error: %w", cfg, err)
	}

	return c.validate()
}
