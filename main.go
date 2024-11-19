package main

import (
	"log"
	"sync"

	"github.com/aspandyar/aspandyar_bot/api"
	"github.com/aspandyar/aspandyar_bot/bot"
	"github.com/aspandyar/aspandyar_bot/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	serverBot, err := bot.NewServerBot(config)
	if err != nil {
		log.Fatal("Cannot create server bot: ", err)
	}

	err = serverBot.SetupRoutes()
	if err != nil {
		log.Fatal("Cannot setup bots routes: ", err)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		serverBot.Start()
	}()

	server, err := api.NewServer(config)
	if err != nil {
		log.Fatal("Cannot create server: ", err)
	}

	go func() {
		defer wg.Done()
		err = server.Start(config.ServerAddress)
		if err != nil {
			log.Fatal("Cannot start server: ", err)
		}
	}()

	wg.Wait()
}
