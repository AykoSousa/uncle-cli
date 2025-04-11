package chatgpt

import (
	"context"
	"log"

	openai "github.com/sashabaranov/go-openai"
	"github.com/AykoSousa/uncle-cli/cmd/handlers"
)

type Model struct {
	*openai.Model
}

func ChatGPTPrompt(API_KEY string, prompt string) {
	apiKey := API_KEY
	client := openai.NewClient(apiKey)

	ctx := context.Background()

	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role: openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	}

	response, err := client.CreateChatCompletion(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	handlers.OpenAPIResponse(response.Choices[0].Message.Content)
}