package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/keybase/go-keybase-chat-bot/kbchat"
	openai "github.com/sashabaranov/go-openai"
	"os/exec"
	"strings"
	"github.com/31333337/kb-gpt-bot/pkg/commands" 
)

var debug bool

var (
	errorStyle = color.New(color.FgHiRed, color.Bold)
)

func printStyled(c *color.Color, format string, a ...interface{}) {
	c.PrintfFunc()(format, a...)
}

func findKeybasePath() (string, error) {
	out, err := exec.Command("which", "keybase").Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func main() {
	flag.BoolVar(&debug, "debug", false, "Enable detailed error logging")
	flag.Parse()

	keybasePath, err := findKeybasePath()
	if err != nil {
		panic("Could not find Keybase binary: " + err.Error())
	}

	client := openai.NewClient("sk-your-api-key") // Replace with your OpenAI API key

	kbc, err := kbchat.Start(kbchat.RunOptions{KeybaseLocation: keybasePath})
	if err != nil {
		if debug {
			fmt.Printf("Detailed Error: %v\n", err)
		} else {
			fmt.Println("Failed to start Keybase chat.")
		}
		return
	}

	sub, err := kbc.ListenForNewTextMessages()
	if err != nil {
		printStyled(errorStyle, "Failed to listen for new messages: %v", err)
		return
	}

	var messages []openai.ChatCompletionMessage

	for {
		msg, err := sub.Read()
		if err != nil {
			printStyled(errorStyle, "Failed to read message: %v", err)
			return
		}

		if msg.Message.Content.TypeName != "text" {
			continue
		}

		userInput := msg.Message.Content.Text.Body

		// Handle special commands
		if strings.HasPrefix(userInput, "/") {
			commandParts := strings.SplitN(userInput[1:], " ", 2)
			command := commandParts[0]
			argument := ""
			if len(commandParts) > 1 {
				argument = commandParts[1]
			}

			switch command {
			case "role":
				commands.HandleRoleCommand(kbc, &msg.Message.Channel, argument)
			case "temperature":
				// Convert argument to float and handle errors
				commands.HandleTemperatureCommand(kbc, &msg.Message.Channel, argument)
			case "settings":
				commands.HandleSettingsCommand(kbc, &msg.Message.Channel)
			case "info":
				commands.HandleInfoCommand(kbc, &msg.Message.Channel)
			case "config":
				// Split argument into key and value and handle errors
				commands.HandleConfigCommand(kbc, &msg.Message.Channel, key, value)
			}
			continue
		}

		if userInput == "exit" {
			break
		}

		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: userInput,
		})

		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model:    openai.GPT3Dot5Turbo,
				Messages: messages,
			},
		)

		if err != nil {
			_, _ = kbc.SendMessageByConvID(msg.Message.ConvID, fmt.Sprintf("Error: %v", err))
			return
		}

		_, _ = kbc.SendMessageByConvID(msg.Message.ConvID, resp.Choices[0].Message.Content)

		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: resp.Choices[0].Message.Content,
		})
	}
}

