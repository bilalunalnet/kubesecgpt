package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

func CheckVulnerability(spec []byte) (string, error) {
	openaiApiKey := os.Getenv("OPENAI_API_KEY")
	client := openai.NewClient(openaiApiKey)

	prompt := `I want you to act as a Kubernetes Security Expert.
	I will give you a deployment manifest file in YAML format, and you will review the file for security best practices and provide me with a list of vulnerabilities in the file.
	I want you to only reply with the list and nothing else, do not write explanations.
	
	`
	prompt += strings.TrimSpace(string(spec))

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
