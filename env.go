package main

import (
	"strings"

	"github.com/BurntSushi/toml"
)

func envSetup() error {

	if _, err := toml.DecodeFile(configFilePath, &c); err != nil {
		return err
	}

	if c.Port == 0 {
		c.Port = 34378
	}

	if !strings.HasSuffix(c.MemeFolder, "/") {
		c.MemeFolder = c.MemeFolder + "/"
	}

	if strings.HasSuffix(c.Endpoint, "/") {
		c.Endpoint = strings.TrimSuffix(c.Endpoint, "/")
	}

	return nil
}
