package main

import (
	"bufio"
	"fmt"
	"log"
	"openai-wrapper/config"
	"openai-wrapper/openai"
	"os"
	"strings"
)

func main() {
	cfg, err := config.LoadConfig("config/config.json")
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	reader := bufio.NewReader(os.Stdin)
	client := openai.NewClient(cfg.OpenAIKey, cfg.OpenAIURL)

	fmt.Println("OpenAI CLI Test Tool")
	fmt.Println("--------------------")
	fmt.Println("Enter 'exit' to quit.")
	fmt.Println()

	// Initialize an empty conversation history
	var conversationHistory []openai.Message

	for {
		fmt.Print("You: ")
		input, _ := reader.ReadString('\n')

		text := strings.TrimSpace(input)

		if text == "exit" {
			break
		}

		// Append the new user message to the conversation history
		conversationHistory = append(conversationHistory, openai.Message{
			Role:    "user",
			Content: text,
		})

		// Create a chat request with the updated conversation history
		chatRequest := openai.ChatRequest{
			Model:    "gpt-4-1106-preview",
			Messages: conversationHistory,
		}

		// Check if the model is supported in the configuration
		if _, exists := cfg.ModelsList[chatRequest.Model]; !exists {
			fmt.Println("Model not supported")
			continue
		}

		// Send the chat request
		chatResponse, err := client.Chat(&chatRequest, log.New(os.Stderr, "LOG: ", log.LstdFlags))
		if err != nil {
			fmt.Printf("Error during chat: %v\n", err)
			continue
		}

		// Print out the AI's response and append it to the conversation history
		for _, choice := range chatResponse.Choices {
			fmt.Printf("AI: %s\n", choice.Message.Content)
			conversationHistory = append(conversationHistory, choice.Message)
		}
	}
}
