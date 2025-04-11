package groq

import (
	"context"
	"log"

	"github.com/AykoSousa/uncle-cli/cmd/handlers"
	"github.com/conneroisu/groq-go"
)

func Groq(GROQ_API_KEY string, model string, userPrompt string) {
	apiKey := GROQ_API_KEY
	if apiKey == "" {
		log.Fatal("GROQ_API_KEY is required")
	}

	client, err := groq.NewClient(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	response, err := client.ChatCompletion(ctx, groq.ChatCompletionRequest{
		Model: groq.ChatModel(model),
		Messages: []groq.ChatCompletionMessage{
			{
				Role: "system",
				Content: "You are a highly intelligent assistant and you must answer the user's question directly, without formatting. Just with text.",
			},
			{
				Role: "user",
				Content: userPrompt,
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	handlers.OpenAPIResponse(response.Choices[0].Message.Content)
}