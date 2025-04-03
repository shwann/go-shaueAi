package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/golang/go-shaueAi/shaueai"
	"io"
	"os"
)

func main() {
	client := shaueai.NewClient(os.Getenv("SHAUEAI_API_KEY"), os.Getenv("SHAUEAI_BASE_URL"))

	req := shaueai.ChatCompletionRequest{
		Stream: true,
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
		stream, err := client.CreateChatCompletionStream(
			context.Background(),
			req,
		)
		if err != nil {
			fmt.Printf("ChatCompletionStream error: %v\n", err)
			return
		}
		defer stream.Close()
		fmt.Print("Stream response: ")
		for {
			var response shaueai.ChatCompletionStreamResponse
			response, err = stream.Recv()
			if errors.Is(err, io.EOF) {
				fmt.Println("\nStream finished")
				break
			}
			if err != nil {
				fmt.Printf("\nStream error: %v\n", err)
				break
			}
			fmt.Print(response.Choices[0].Delta.Content)
		}
		fmt.Print("> ")
	}
}
