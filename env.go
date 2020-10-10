package main

import (
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/shitpostingio/randomapi/config"
)

var (
	c              *config.Config
	debug          bool
	allowedOrigins []string

	path string
)

func envSetup() error {

	if _, err := toml.DecodeFile(path, &c); err != nil {
		return err
	}

	if c.Port == 0 {
		c.Port = 34378
	}

	if !strings.HasSuffix(c.MemeFolder, "/") {
		c.MemeFolder = c.MemeFolder + "/"
	}

	if !strings.HasSuffix(c.MemeSymlinkFolder, "/") {
		c.MemeSymlinkFolder = c.MemeSymlinkFolder + "/"
	}

	return nil
}
