package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"github.com/pterm/pterm"
	"google.golang.org/api/option"
)

type chatHistory struct {
	id       int
	response string
}

type History struct {
	ChatHistory []chatHistory
}

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

	lh := loadHistory()
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(GEMINI_API_KEY))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	for {
		response, err := model.GenerateContent(ctx, genai.Text(userInput()))
		if err != nil {
			log.Fatal(err)
		}
		spinnerSuccess, _ := pterm.DefaultSpinner.Start("Processing input... (will succeed)")
		time.Sleep(time.Second * 2) // Simulate 3 seconds of processing something.
		spinnerSuccess.Success()

		printResponse(response, lh)

	}
}

func userInput() string {
	result, _ := pterm.DefaultInteractiveTextInput.Show(pterm.Magenta("Me"))
	fmt.Println()

	return result
}

func printResponse(resp *genai.GenerateContentResponse, ch History) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				pterm.DefaultBasicText.Println(pterm.LightBlue("AI: "), part)
				getResponseAddHistory(&ch, part)
				// fmt.Println(part)
				// pterm.DefaultBox.WithRightPadding(1).WithLeftPadding(1).WithTopPadding(2).WithBottomPadding(2).Println(pterm.LightBlue("AI: "), part)
			}
		}
	}
}

func getResponseAddHistory(chatHis *History, responses interface{}) {
	chatAdd := chatHistory{
		id:       len(chatHis.ChatHistory) + 1,
		response: fmt.Sprintf("%v", responses),
	}
	chatHis.ChatHistory = append(chatHis.ChatHistory, chatAdd)
}

func loadHistory() History {
	var history History

	return history
}
