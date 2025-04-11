package deepseek

import (
    "context"
    "log"

    deepseek "github.com/cohesion-org/deepseek-go"
    "github.com/AykoSousa/uncle-cli/cmd/handlers"
)

func DeepseekPrompt(API_KEY string, prompt string) {
    apiKey := API_KEY
    client := deepseek.NewClient(apiKey)

    ctx := context.Background()

    req := &deepseek.ChatCompletionRequest{
        Model: "deepseek-chat",
        Messages: []deepseek.ChatCompletionMessage{
            {
                Role:    "user",
                Content: prompt,
            },
        },
    }

    resp, err := client.CreateChatCompletion(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    handlers.OpenAPIResponse(resp.Choices[0].Message.Content)
}