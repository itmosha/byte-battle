package main

import (
	"log"

	"github.com/Mihalych32/byte-battle/tree/main/byte-battle_backend/internal/app/apiserver"
)

func main() {
	config := apiserver.NewConfig()

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
