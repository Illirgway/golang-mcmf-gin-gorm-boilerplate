/**
 *
 * Copyright (c) 2024 Illirgway
 *
 */

package config

import (
	"net"
	"strconv"
)

func (c *Config) ListenAddr() string {

	port := strconv.FormatUint(uint64(c.Port), 10)

	return net.JoinHostPort(c.Listen, port)
}

func (c *Config) IsDebug() bool {
	return c.Debug
}
