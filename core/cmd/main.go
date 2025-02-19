package main

import (
	"flag"
	"log"

	"github.com/glu/video-real-time-ranking/core/config"
)

func main() {
	flag.Parse()

	_, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}
}
