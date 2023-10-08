package main

import (
	"context"

	cfg "github.com/aniljaiswalcs/image-batch-download/config"
	"github.com/aniljaiswalcs/image-batch-download/core/app"
)

const configFilePath = "./config.yml"

func main() {
	config := cfg.NewConfig(configFilePath)
	ctx := context.Background()
	app.StartImageDownloaderApp(ctx, config)
}
