package chat

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

var (
	client      *openai.Client
	chatHistory []openai.ChatCompletionMessage
)

func InitChatWithSystemRole(token, systemMessage string) error {
	client = openai.NewClient(token)
	if client == nil {
		return fmt.Errorf("cannot create OpenAI client")
	}

	chatHistory = append(chatHistory, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: systemMessage,
	})

	return nil
}

func AddSystemMessageToChatGPT(message string) {
	chatHistory = append(chatHistory, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: message,
	})
}

func SendMessageToChatGPT(message string) (response openai.ChatCompletionResponse, err error) {
	chatHistory = append(chatHistory, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: message,
	})

	response, err = client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT4oMini,
			Messages: chatHistory,
		},
	)

	chatHistory = append(chatHistory, response.Choices[0].Message)

	return
}

func GetChatHistory() []openai.ChatCompletionMessage {
	return chatHistory
}
