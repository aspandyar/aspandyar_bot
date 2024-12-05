package handlers

import (
	"fmt"

	"github.com/aspandyar/aspandyar_bot/bot/chat"
	"github.com/aspandyar/aspandyar_bot/util"
	tele "gopkg.in/telebot.v3"
)

type ServerBot struct {
	Bot *tele.Bot
}

func (server *ServerBot) StartConversation(c tele.Context) error {
	server.Bot.Handle(tele.OnAddedToGroup, func(c tele.Context) error {
		c.Send("Hello! I'm a bot that can help you with your questions. Just type your question and I'll do my best to help you. To start a new conversation, type /begin.")
		return nil
	})

	server.Bot.Handle("/history", func(c tele.Context) error {
		history := chat.GetChatHistory()

		var historyMessage string
		for _, message := range history {
			historyMessage += fmt.Sprintf("%s: %s\n", message.Role, message.Content)
		}
		c.Send(historyMessage)

		return nil
	})

	server.Bot.Handle("/food", func(c tele.Context) error {
		foodPrompt, err := util.LoadPromptByName("food")
		if err != nil {
			return err
		}

		chat.AddSystemMessageToChatGPT(foodPrompt)

		generateMessage := "Explain what can you do now"

		response, err := chat.SendMessageToChatGPT(generateMessage)
		if err != nil {
			c.Send("Sorry, I'm having trouble understanding you right now.")
		}

		if len(response.Choices) > 0 && response.Choices[0].Message.Content != "" {
			err = c.Send(response.Choices[0].Message.Content)
			if err != nil {
				return err
			}
		} else {
			c.Send("Hmm, I didn't quite catch that. Can you try again?")
		}

		return nil
	})

	server.Bot.Handle(tele.OnText, func(c tele.Context) error {
		fmt.Println("Received message:", c.Text())
		userMessage := c.Text()
		response, err := chat.SendMessageToChatGPT(userMessage)
		if err != nil {
			c.Send("Sorry, I'm having trouble understanding you right now.")
		}

		if len(response.Choices) > 0 && response.Choices[0].Message.Content != "" {
			err = c.Send(response.Choices[0].Message.Content)
			if err != nil {
				return err
			}
		} else {
			c.Send("Hmm, I didn't quite catch that. Can you try again?")
		}

		return nil

	})

	return nil
}
