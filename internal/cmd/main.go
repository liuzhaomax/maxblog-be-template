package main

import (
	"context"
	"maxblog-be-template/internal/app"
)

func main() {
	const ConfigDir = "env/raw"
	const ConfigFile = "dev.yaml"
	ctx := context.Background()
	app.Launch(
		ctx,
		app.SetConfigDir(ConfigDir),
		app.SetConfigFile(ConfigFile),
	)
}