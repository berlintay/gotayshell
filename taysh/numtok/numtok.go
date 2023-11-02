package main

import (
	"fmt"
)

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
