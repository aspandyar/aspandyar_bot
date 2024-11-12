package main

import (
	"log"

	"github.com/aspandyar/aspandyar_bot/api"
	"github.com/aspandyar/aspandyar_bot/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	server, err := api.NewServer(config)
	if err != nil {
		log.Fatal("Cannot create server: ", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
