package routes

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/aspandyar/aspandyar_bot/bot"
	"github.com/aspandyar/aspandyar_bot/bot/handlers"
	tele "gopkg.in/telebot.v3"
)

type ServerBotWrapper struct {
	*bot.ServerBot
}

func (server *ServerBotWrapper) SetupRoutes() error {

	server.Bot.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	server.Bot.Handle("/chatid", func(c tele.Context) error {
		chatID := c.Chat().ID
		server.ChatID = chatID
		return c.Send(fmt.Sprintf("Chat ID: %d", chatID))
	})

	server.Bot.Handle("/userid", func(c tele.Context) error {
		userID := c.Sender().ID

		if userID == 0 {
			return c.Send("User ID not found.")
		}
		return c.Send(fmt.Sprintf("User ID: %d", userID))
	})

	server.Bot.Handle("/start", func(c tele.Context) error {
		return c.Send("Welcome to the bot! Use /hello to get greeted.")
	})

	server.Bot.Handle("/ping", func(c tele.Context) error {
		payload := c.Message().Payload

		if payload == "" {
			return c.Send("Pong!")
		}

		fmt.Println("Payload: ", payload)
		ans_array := strings.Split(payload, " ")
		fmt.Println(ans_array)

		if len(ans_array) > 0 && ans_array[0][0] == '@' {
			c.Send("Got you!")

			count_number := 100

			if len(ans_array) == 2 {
				count, err := strconv.Atoi(ans_array[1])
				if err != nil {
					return c.Send("2nd argument should be number.")
				}
				count_number = count
			} else if len(ans_array) > 2 {
				return c.Send("Too many arguments!")
			}

			if err := bot.SendMessagesInBatches(c, ans_array[0], count_number); err != nil {
				log.Printf("Error sending messages: %v", err)
				return c.Send("Something went wrong while sending messages.")
			}

		} else {
			c.Send("Payload must start with '@'!")
		}

		return c.Send("Pong!")
	})

	server.Bot.Handle("/finish", func(c tele.Context) error {
		return nil
	})

	server.Bot.Handle("/begin", func(c tele.Context) error {
		c.Send("Let's start the conversation!")
		serverBot := &handlers.ServerBot{Bot: server.Bot}
		serverBot.StartConversation(c)
		return nil
	})

	server.Bot.Handle("/help", func(c tele.Context) error {
		payload := c.Message().Payload

		if payload == "" {
			helpMessage := "Available commands:\n" +
				"/hello - Greet the bot\n" +
				"/chatid - Get your chat ID\n" +
				"/start - Start the bot\n" +
				"/ping - Ping the bot\n" +
				"/finish - Finish the conversation (chat gpt) \n" +
				"/begin - Start a new conversation (chat gpt) \n" +
				"/help - Show this help message or detailed help for a specific command\n" +
				"/food - Get food-related information\n" +
				"/history - Get chat history\n" +
				"Use '/help [command]' for detailed information about a specific command."
			return c.Send(helpMessage)
		}

		switch strings.ToLower(payload) {
		case "ping":
			detailedHelp := "Detailed help for /ping:\n" +
				"/ping - Responds with 'Pong!'\n" +
				"/ping @[nickname] - Sends '@[nickname]' 100 times\n" +
				"/ping @[nickname] [count] - Sends '@[nickname]' [count] times (count must be a number)"
			return c.Send(detailedHelp)
		case "hello":
			detailedHelp := "Detailed help for /hello:\n" +
				"/hello - The bot will greet you with 'Hello!'"
			return c.Send(detailedHelp)
		case "chatid":
			detailedHelp := "Detailed help for /chatid:\n" +
				`/chatid - The bot will respond with your chat ID. 
				That chat ID would stored, and if bot would stop, he woudl send his last words to that chat`
			return c.Send(detailedHelp)
		case "start":
			detailedHelp := "Detailed help for /start:\n" +
				"/start - The bot will welcome you and provide basic instructions."
			return c.Send(detailedHelp)
		case "finish":
			detailedHelp := "Detailed help for /finish:\n" +
				"/finish - Ends the current conversation or session."
			return c.Send(detailedHelp)
		case "begin":
			detailedHelp := "Detailed help for /begin:\n" +
				"/begin - Starts a new conversation or session."
			return c.Send(detailedHelp)
		case "food":
			detailedHelp := "Detailed help for /food:\n" +
				"/food - Provides food-related information (details depend on implementation)."
			return c.Send(detailedHelp)
		case "history":
			detailedHelp := "Detailed help for /history:\n" +
				"/history - Retrieves the chat history (details depend on implementation)."
			return c.Send(detailedHelp)
		default:
			return c.Send("Unknown command. Use '/help' to see the list of available commands.")
		}
	})

	return nil
}
