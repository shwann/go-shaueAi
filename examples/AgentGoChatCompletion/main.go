package main

import (
	"bufio"
	"fmt"
	"github.com/shwann/go-shaueAi/shaueai"
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
		resp, err := client.AgentGoChatCompletion(messages, shaueai.DOUBAO_FunctionCall)
		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err.Error())
			continue
		}
		messages = append(messages, shaueai.ChatCompletionMessage{
			Role:    shaueai.ChatMessageRoleAssistant,
			Content: resp,
		})
		fmt.Printf("%s\n\n", resp)
		fmt.Print("> ")
	}
}
