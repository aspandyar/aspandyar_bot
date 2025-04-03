package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aspandyar/aspandyar_bot/bot/handlers"
	"github.com/aspandyar/aspandyar_bot/util"
	tele "gopkg.in/telebot.v3"
)

type ServerBot struct {
	bot    *tele.Bot
	chatID int64
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

	return &ServerBot{bot: b, chatID: config.TelegramChatID}, nil
}

func (server *ServerBot) SetupRoutes() error {
	server.bot.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	server.bot.Handle("/chatid", func(c tele.Context) error {
		chatID := c.Chat().ID
		server.chatID = chatID
		return c.Send(fmt.Sprintf("Chat ID: %d", chatID))
	})

	server.bot.Handle("/start", func(c tele.Context) error {
		return c.Send("Welcome to the bot! Use /hello to get greeted.")
	})

	server.bot.Handle("/ping", func(c tele.Context) error {
		payload := c.Message().Payload

		if payload == "" {
			return c.Send("Pong!")
		}

		if len(payload) > 0 && payload[0] == '@' {
			c.Send("Got you!")

			if err := sendMessagesInBatches(c, payload, 100); err != nil {
				log.Printf("Error sending messages: %v", err)
				return c.Send("Something went wrong while sending messages.")
			}

		} else {
			c.Send("Payload must start with '@'!")
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

func (server *ServerBot) Start(config util.Config) {
	log.Println("Starting the bot...")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		server.bot.Start()
	}()

	<-stop
	log.Println("Shutting down the bot...")

	_, err := server.bot.Send(tele.ChatID(server.chatID), "I am duying... Goodbye! :(")
	if err != nil {
		log.Printf("Failed to send shutdown message: %v", err)
	}

	server.bot.Stop()
	log.Println("Bot stopped.")
}
