package main

import (
	"log"

	"github.com/aspandyar/aspandyar_bot/bot"
	"github.com/aspandyar/aspandyar_bot/bot/chat"
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

	prompt, err := util.LoadPromptByName("prompt")
	if err != nil {
		log.Fatal("Cannot load prompt: ", err)
	}

	err = chat.InitChatWithSystemRole(config.OpenaiToken, prompt)
	if err != nil {
		log.Fatal("Cannot init chatgpt, check openai token: ", err)
	}

	serverBot.Start(config)
}
