package main

import (
	"flag"
	"log"

	"github.com/Mihalych32/byte-battle/tree/main/byte-battle_backend/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.yaml", "path to the apiserver config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
