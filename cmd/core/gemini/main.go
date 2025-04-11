package gemini

import (
    "context"
    "log"

	"github.com/AykoSousa/uncle-cli/cmd/handlers"
    "github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func Gemini(API_KEY string, modelName string, prompt string) {
    ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(API_KEY))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel(modelName)
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
	}

	handlers.GeminiResponse(resp)
}