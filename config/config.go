package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

// Represent database server and credential
type Config struct {
	Server   string
	Database string
}

// Read and parse the configuration file
func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
