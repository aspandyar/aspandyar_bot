package bot

import (
	"log"
	"time"

	"github.com/aspandyar/aspandyar_bot/util"
	tele "gopkg.in/telebot.v3"
)

type ServerBot struct {
	bot *tele.Bot
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

	return &ServerBot{bot: b}, nil
}

func (server *ServerBot) SetupRoutes() error {
	server.bot.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	server.bot.Handle("/start", func(c tele.Context) error {
		return c.Send("Welcome to the bot! Use /hello to get greeted.")
	})

	return nil
}

func (server *ServerBot) Start() {
	log.Println("Starting the bot...")
	server.bot.Start()
}