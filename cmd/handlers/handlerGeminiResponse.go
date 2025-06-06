package handlers

import (
	"fmt"
	
	"github.com/google/generative-ai-go/genai"
)

func GeminiResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
	}
	fmt.Println("*****")
}