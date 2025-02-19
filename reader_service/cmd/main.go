package main

import (
	"flag"
	"log"

	"github.com/glu/video-real-time-ranking/core/pkg/logger"
	"github.com/glu/video-real-time-ranking/reader_service/config"
	"github.com/glu/video-real-time-ranking/reader_service/server"
)

func main() {
	flag.Parse()

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	appLogger := logger.NewAppLogger(cfg.Logger)
	appLogger.InitLogger()
	appLogger.WithName("ReaderService")

	s := server.NewServer(appLogger, cfg)
	appLogger.Fatal(s.Run())
}
