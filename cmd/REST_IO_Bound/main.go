package main

import (
	"fmt"
	"log"
	"rest_io_bound/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg)
}
