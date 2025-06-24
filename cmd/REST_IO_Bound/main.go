package main

import (
	"fmt"
	"log"
	"rest_io_bound/internal/config"
	"rest_io_bound/internal/server"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg)

	Server := server.NewWebServer(cfg.HTTP.Port)

	Server.Launch()
}
