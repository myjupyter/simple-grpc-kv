package main

import (
	"errors"
	"flag"
	"os"

	"github.com/myjupyter/simple-grpc-kv/internal/app"
	"github.com/myjupyter/simple-grpc-kv/pkg/config"
	log "google.golang.org/grpc/grpclog"
)

func main() {
	var path string
	flag.StringVar(&path, "configs", "configs/config.json", "config path")
	flag.Parse()

	cfg, err := config.NewConfigFromFile(path)
	if err != nil {
		log.Fatal(err)
	}

	logger := log.NewLoggerV2(os.Stdout, os.Stdout, os.Stderr)

	application := app.NewApplication(cfg, logger)
	if err = application.Run(); err != nil {
		if !errors.Is(err, app.ErrAppInterrupted) {
			logger.Fatal(err)
		}
	}
	logger.Info("application gracefully shutdowned")
}
