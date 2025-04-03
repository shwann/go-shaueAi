package main

import (
	"context"
	"fmt"
	"github.com/golang/go-shaueAi/jsonschema"
	"github.com/golang/go-shaueAi/shaueai"
	"os"
)

func main() {
	ctx := context.Background()
	client := shaueai.NewClient(os.Getenv("SHAUEAI_API_KEY"), os.Getenv("SHAUEAI_BASE_URL"))

	// describe the function & its inputs
	params := jsonschema.Definition{
		Type: jsonschema.Object,
		Properties: map[string]jsonschema.Definition{
			"location": {
				Type:        jsonschema.String,
				Description: "城市或市，例如：上海",
			},
			"unit": {
				Type: jsonschema.String,
				Enum: []string{"摄氏度", "华氏度"},
			},
		},
		Required: []string{"location"},
	}
	f := shaueai.FunctionDefinition{
		Name:        "get_current_weather",
		Description: "Get the current weather in a given location",
		Parameters:  params,
	}
	t := shaueai.Tool{
		Type:     shaueai.ToolTypeFunction,
		Function: &f,
	}

	// simulate user asking a question that requires the function
	dialogue := []shaueai.ChatCompletionMessage{
		{Role: shaueai.ChatMessageRoleUser, Content: "上海今天天气怎么样？"},
	}
	fmt.Printf("Asking OpenAI '%v' and providing it a '%v()' function...\n",
		dialogue[0].Content, f.Name)
	resp, err := client.CreateChatCompletion(ctx,
		shaueai.ChatCompletionRequest{
			Model:    shaueai.DOUBAO_FunctionCall,
			Messages: dialogue,
			Tools:    []shaueai.Tool{t},
		},
	)
	if err != nil || len(resp.Choices) != 1 {
		fmt.Printf("Completion error: err:%v len(choices):%v\n", err,
			len(resp.Choices))
		return
	}
	msg := resp.Choices[0].Message
	if len(msg.ToolCalls) != 1 {
		fmt.Printf("Completion error: len(toolcalls): %v\n", len(msg.ToolCalls))
		return
	}

	// simulate calling the function & responding to OpenAI
	dialogue = append(dialogue, msg)
	fmt.Printf("OpenAI called us back wanting to invoke our function '%v' with params '%v'\n",
		msg.ToolCalls[0].Function.Name, msg.ToolCalls[0].Function.Arguments)
	dialogue = append(dialogue, shaueai.ChatCompletionMessage{
		Role:       shaueai.ChatMessageRoleTool,
		Content:    "天气晴朗，10-22摄氏度.",
		Name:       msg.ToolCalls[0].Function.Name,
		ToolCallID: msg.ToolCalls[0].ID,
	})
	fmt.Printf("Sending OpenAI our '%v()' function's response and requesting the reply to the original question...\n",
		f.Name)
	resp, err = client.CreateChatCompletion(ctx,
		shaueai.ChatCompletionRequest{
			Model:    shaueai.DOUBAO_FunctionCall,
			Messages: dialogue,
			Tools:    []shaueai.Tool{t},
		},
	)
	if err != nil || len(resp.Choices) != 1 {
		fmt.Printf("2nd completion error: err:%v len(choices):%v\n", err,
			len(resp.Choices))
		return
	}

	// display OpenAI's response to the original question utilizing our function
	msg = resp.Choices[0].Message
	fmt.Printf("OpenAI answered the original request with: %v\n",
		msg.Content)
}
