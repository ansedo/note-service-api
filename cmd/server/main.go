package main

import (
	"context"
	"flag"
	"log"

	"github.com/ansedo/note-service-api/internal/app"
)

var pathConfig string

func init() {
	flag.StringVar(&pathConfig, "config", "config/config.json", "path to configuration file")
}

func main() {
	flag.Parse()

	ctx := context.Background()

	a, err := app.NewApp(ctx, pathConfig)
	if err != nil {
		log.Fatalf("failed to create app: %s", err.Error())
	}

	if err = a.Run(); err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
