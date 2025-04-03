package main

import (
	"bufio"
	"fmt"
	"github.com/golang/go-shaueAi/shaueai"
	"os"
)

func main() {
	client := shaueai.NewClient(os.Getenv("SHAUEAI_API_KEY"), os.Getenv("SHAUEAI_BASE_URL"))
	fmt.Println("---------------------")
	fmt.Print("> ")
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		resp, err := client.AgentGoCompletion("you are a helpful chatbot", s.Text(), shaueai.DOUBAO_FunctionCall)
		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err.Error())
			continue
		}
		fmt.Printf("%s\n\n", resp)
		fmt.Print("> ")
	}
}
