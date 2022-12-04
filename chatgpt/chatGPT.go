package chatgpt

import (
	"context"
	"log"
	"os"

	gogpt "github.com/sashabaranov/go-gpt3"
)

func SendMessageToGPT(msg string) string {
	var OPENAI_KEY string = os.Getenv("OPENAI_API_KEY")
	c := gogpt.NewClient(OPENAI_KEY)
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:     "text-davinci-003",
		MaxTokens: 2000,
		Prompt:    msg,
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		log.Println("Unable to send message to ChatGPT,", err)
	}
	return resp.Choices[0].Text
}
