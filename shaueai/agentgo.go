package shaueai

import (
	"context"
	"errors"
	"io"
)

func (client *Client) AgentGoCompletion(systemPrompt string, userPrompt string, model string) (resp string, err error) {
	req := ChatCompletionRequest{
		Stream:   false,
		Model:    model,
		Messages: []ChatCompletionMessage{},
	}
	if systemPrompt != "" {
		req.Messages = append(req.Messages, ChatCompletionMessage{
			Role:    ChatMessageRoleSystem,
			Content: systemPrompt,
		})
	}
	if userPrompt != "" {
		req.Messages = append(req.Messages, ChatCompletionMessage{
			Role:    ChatMessageRoleUser,
			Content: userPrompt,
		})
	}
	res, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return resp, err
	}
	return res.Choices[0].Message.Content, err
}

func (client *Client) AgentGoChatCompletion(messages []ChatCompletionMessage, model string) (resp string, err error) {
	req := ChatCompletionRequest{
		Stream:   false,
		Model:    model,
		Messages: messages,
	}
	res, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return resp, err
	}
	return res.Choices[0].Message.Content, err
}

func (client *Client) AgentGoCompletionStream(systemPrompt string, userPrompt string, model string, callback func(resp string, err error)) {
	req := ChatCompletionRequest{
		Stream:   true,
		Model:    model,
		Messages: []ChatCompletionMessage{},
	}
	if systemPrompt != "" {
		req.Messages = append(req.Messages, ChatCompletionMessage{
			Role:    ChatMessageRoleSystem,
			Content: systemPrompt,
		})
	}
	if userPrompt != "" {
		req.Messages = append(req.Messages, ChatCompletionMessage{
			Role:    ChatMessageRoleUser,
			Content: userPrompt,
		})
	}
	stream, err := client.CreateChatCompletionStream(
		context.Background(),
		req,
	)
	if err != nil {
		callback("", err)
		return
	}
	defer stream.Close()
	for {
		var response ChatCompletionStreamResponse
		response, err = stream.Recv()
		if errors.Is(err, io.EOF) {
			//fmt.Println("\nStream finished")
			callback("[DONE]", nil)
			return
		}
		if err != nil {
			//fmt.Printf("\nStream error: %v\n", err)
			callback("", err)
			return
		}
		callback(response.Choices[0].Delta.Content, nil)
	}
}

func (client *Client) AgentGoChatCompletionStream(messages []ChatCompletionMessage, model string, callback func(resp string, err error)) {
	req := ChatCompletionRequest{
		Stream:   false,
		Model:    model,
		Messages: messages,
	}
	stream, err := client.CreateChatCompletionStream(
		context.Background(),
		req,
	)
	if err != nil {
		callback("", err)
		return
	}
	defer stream.Close()
	for {
		var response ChatCompletionStreamResponse
		response, err = stream.Recv()
		if errors.Is(err, io.EOF) {
			callback("", err)
			return
		}
		if err != nil {
			callback("", err)
			return
		}
		callback(response.Choices[0].Delta.Content, nil)
	}
}
