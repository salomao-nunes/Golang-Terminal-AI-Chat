package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"github.com/pterm/pterm"
	"google.golang.org/api/option"
)

func loadEnv() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	API_KEY := os.Getenv("GEMINI_API_KEY")

	return API_KEY
}

func main() {
	GEMINI_API_KEY := loadEnv()

	inputUser := userInput()

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(GEMINI_API_KEY))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")

	response, err := model.GenerateContent(ctx, genai.Text(inputUser))
	if err != nil {
		log.Fatal(err)
	}

	printResponse(response)
}

func userInput() string {
	result, _ := pterm.DefaultInteractiveTextInput.Show(pterm.Magenta("Me"))
	fmt.Println()

	return result
}

func printResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				pterm.DefaultBasicText.Println(pterm.LightBlue("AI: "), part)
				// fmt.Println(part)
			}
		}
	}
}
