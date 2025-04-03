package bot

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aspandyar/aspandyar_bot/util"
	tele "gopkg.in/telebot.v3"
)

type ServerBot struct {
	Bot    *tele.Bot
	ChatID int64
}

func NewServerBot(config util.Config) (*ServerBot, error) {
	pref := tele.Settings{
		Token:  config.TelegramToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		return nil, err
	}

	return &ServerBot{Bot: b, ChatID: config.TelegramChatID}, nil
}

func (server *ServerBot) Start(config util.Config) {
	log.Println("Starting the bot...")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		server.Bot.Start()
	}()

	<-stop
	log.Println("Shutting down the bot...")

	_, err := server.Bot.Send(tele.ChatID(server.ChatID), "I am duying... Goodbye! :(")
	if err != nil {
		log.Printf("Failed to send shutdown message: %v", err)
	}

	server.Bot.Stop()
	log.Println("Bot stopped.")
}
