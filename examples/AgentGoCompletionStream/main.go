package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/golang/go-shaueAi/shaueai"
	"io"
	"os"
)

func main() {
	client := shaueai.NewClient(os.Getenv("SHAUEAI_API_KEY"), os.Getenv("SHAUEAI_BASE_URL"))
	fmt.Println("---------------------")
	fmt.Print("> ")
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
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
		client.AgentGoCompletionStream("you are a helpful chatbot", s.Text(), shaueai.DOUBAO_FunctionCall, callback)
		fmt.Print("> ")
	}
}
