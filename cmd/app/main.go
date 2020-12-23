package main

import (
	"errors"
	"flag"
	"log"

	"github.com/myjupyter/simple-grpc-kv/internal/app"
	"github.com/myjupyter/simple-grpc-kv/pkg/config"
)

func main() {
	var path string
	flag.StringVar(&path, "configs", "configs/config.json", "config path")
	flag.Parse()

	cfg, err := config.NewConfigFromFile(path)
	if err != nil {
		log.Fatal(err)
	}

	application := app.NewApplication(cfg)
	if err = application.Run(); err != nil {
		if !errors.Is(err, app.ErrAppInterrupted) {
			log.Fatal(err)
		}
	}
	log.Println("application gracefully shutdowned")
}
