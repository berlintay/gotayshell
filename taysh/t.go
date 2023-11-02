package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	godotenv "github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

func getEnvVar(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}
	return os.Getenv(key)
}

func main() {
	client := openai.NewClient(getEnvVar("OPENAI_API_KEY"))

	req := openai.ChatCompletionRequest{
		Model:            openai.GPT3Dot5Turbo,
		Messages:         []openai.ChatCompletionMessage{{Role: openai.ChatMessageRoleSystem, Content: "you are a helpful chatbot experienced in developing programs"}},
		MaxTokens:        0,
		Temperature:      1,
		TopP:             0,
		N:                0,
		Stream:           false,
		Stop:             []string{},
		PresencePenalty:  0,
		FrequencyPenalty: 0,
		LogitBias:        map[string]int{},
		User:             "",
		Functions:        []openai.FunctionDefinition{},
		FunctionCall:     nil,
	}
	fmt.Println("Conversation")
	fmt.Println("---------------------")
	fmt.Print("> ")
	s := bufio.NewScanner(os.Stdin)
	conversation := []map[string]string{} // To store conversation messages

	for s.Scan() {
		userMessage := s.Text()
		req.Messages = append(req.Messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: userMessage,
		})

		resp, err := client.CreateChatCompletion(context.Background(), req)
		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			continue
		}

		// Store the response message in the conversation
		conversation = append(conversation, map[string]string{"role": "user", "content": userMessage})
		conversation = append(conversation, map[string]string{"role": "assistant", "content": resp.Choices[0].Message.Content})

		fmt.Printf("Assistant: %s\n\n", resp.Choices[0].Message.Content)
		fmt.Print("> ")
		// Calculate the number of tokens in the conversation
		numTokens := numTokensFromMessages(conversation, "gpt-3.5-turbo-0613")
		fmt.Printf("Number of tokens in the conversation: %d\n", numTokens)
	}
}

func numTokensFromMessages(messages []map[string]string, model string) int {
	var tokensPerMessage, tokensPerName int
	switch model {
	case "gpt-3.5-turbo-0613", "gpt-3.5-turbo-16k-0613", "gpt-4-0314", "gpt-4-32k-0314", "gpt-4-0613", "gpt-4-32k-0613":
		tokensPerMessage = 3
		tokensPerName = 1
	case "gpt3.5-turbo-0301":
		tokensPerMessage = 4
		tokensPerName = -1
	default:
		panic(fmt.Sprintf("numTokensFromMessages() is not implemented for model %s.", model))
	}
	numTokens := 0
	for _, message := range messages {
		numTokens += tokensPerMessage
		for key, value := range message {
			numTokens += len([]rune(value))
			if key == "name" {
				numTokens += tokensPerName
			}
		}
	}
	numTokens += 3
	return numTokens
}
