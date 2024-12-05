package bot

import (
	"log"
	"time"

	"github.com/aspandyar/aspandyar_bot/bot/handlers"
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

	server.bot.Handle("/ping", func(c tele.Context) error {
		payload := c.Message().Payload

		if payload == "" {
			return c.Send("Pong!")
		}

		if payload[0] == '@' {
			c.Send("Got you!")
			for i := 0; i < 100; i++ {
				c.Send(payload)
			}
		}

		return c.Send("Pong!")

	})

	server.bot.Handle("/finish", func(c tele.Context) error {
		return nil
	})

	server.bot.Handle("/begin", func(c tele.Context) error {
		c.Send("Let's start the conversation!")
		serverBot := &handlers.ServerBot{Bot: server.bot}
		serverBot.StartConversation(c)
		return nil
	})

	return nil
}

func (server *ServerBot) Start() {
	log.Println("Starting the bot...")
	server.bot.Start()
}
