package internal

import (
	util "github.com/Floor-Gang/utilpkg"
	"log"
	"strings"
)

type Config struct {
	Token               string   `yaml:"token"`
	Prefix              string   `yaml:"prefix"`
	Auth                string   `yaml:"auth_server"`
}

// This will get the current configuration file. If it doesn't exist then a
// new one will be made.
func GetConfig(location string) (config Config) {
	config = Config{
		Token:               "",
		Prefix:              ".rr",
		Auth:                "",
	}
	err := util.GetConfig(location, &config)

	if err != nil && strings.Contains(err.Error(), "default configuration") {
		log.Fatalln("A default configuration has been made.")
	}

	return config
}
