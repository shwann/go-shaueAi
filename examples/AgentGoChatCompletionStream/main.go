package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/shwann/go-shaueAi/shaueai"
	"io"
	"os"
)

func main() {
	client := shaueai.NewClient(os.Getenv("SHAUEAI_API_KEY"), os.Getenv("SHAUEAI_BASE_URL"))
	messages := []shaueai.ChatCompletionMessage{}
	messages = append(messages, shaueai.ChatCompletionMessage{
		Role:    shaueai.ChatMessageRoleSystem,
		Content: "you are a helpful chatbot",
	})
	fmt.Println("---------------------")
	fmt.Print("> ")
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		messages = append(messages, shaueai.ChatCompletionMessage{
			Role:    shaueai.ChatMessageRoleUser,
			Content: s.Text(),
		})
		callback := func(res string, err error) {
			if errors.Is(err, io.EOF) {
				fmt.Println("\nStream finished")
				return
			}
			if err != nil {
				println(fmt.Sprintf("Async Request failed: %v", err))
				return
			}
			fmt.Print(res)
		}
		fmt.Print("Stream response: ")
		client.AgentGoChatCompletionStream(messages, shaueai.DOUBAO_FunctionCall, callback)
		fmt.Print("> ")
	}
}
