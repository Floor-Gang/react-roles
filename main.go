package main

import (
	"github.com/Floor-Gang/react-roles/internal"
	util "github.com/Floor-Gang/utilpkg"
)

const (
	configLocation = "config.yml"
)

func main() {
	config := internal.GetConfig(configLocation)
	internal.Start(config, configLocation)

	util.KeepAlive()
}
