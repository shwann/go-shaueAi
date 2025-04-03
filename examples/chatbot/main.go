package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/golang/go-shaueAi/shaueai"
	"os"
)

func main() {
	client := shaueai.NewClient(os.Getenv("SHAUEAI_API_KEY"), os.Getenv("SHAUEAI_BASE_URL"))

	req := shaueai.ChatCompletionRequest{
		Stream: false,
		Model:  shaueai.DOUBAO_FunctionCall,
		Messages: []shaueai.ChatCompletionMessage{
			{
				Role:    shaueai.ChatMessageRoleSystem,
				Content: "you are a helpful chatbot",
			},
		},
	}
	fmt.Println("Conversation")
	fmt.Println("---------------------")
	fmt.Print("> ")
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		req.Messages = append(req.Messages, shaueai.ChatCompletionMessage{
			Role:    shaueai.ChatMessageRoleUser,
			Content: s.Text(),
		})
		resp, err := client.CreateChatCompletion(context.Background(), req)
		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			continue
		}
		fmt.Printf("%s\n\n", resp.Choices[0].Message.Content)
		req.Messages = append(req.Messages, resp.Choices[0].Message)
		fmt.Print("> ")
	}
}
