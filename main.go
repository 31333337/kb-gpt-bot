package main

import (
    "strings"
    "path/filepath"
    "os"
    "fmt"
    "context"
	"github.com/fatih/color"
	"github.com/keybase/go-keybase-chat-bot/kbchat"
	openai "github.com/sashabaranov/go-openai"
	"os/exec"
    "flag"
)
var debug bool
var (
	errorStyle = color.New(color.FgHiRed, color.Bold)
)

func printStyled(c *color.Color, format string, a ...interface{}) {
	c.PrintfFunc()(format, a...)
}

func resolveSymlink(path string) (string, error) {
    resolvedPath, err := os.Readlink(path)
    if err != nil {
        return "", err
    }
    if !filepath.IsAbs(resolvedPath) {
        resolvedPath = filepath.Join(filepath.Dir(path), resolvedPath)
    }
    return resolvedPath, nil
}
func findKeybasePath() (string, error) {
    out, err := exec.Command("which", "keybase").Output()
    if err != nil {
        return "", err
    }
    path := string(out)
    path = strings.TrimSpace(path) // To remove any newline characters
    resolvedPath, err := resolveSymlink(path)
    if err != nil {
        return "", err
    }
    return resolvedPath, nil
}

func main() {
    flag.BoolVar(&debug, "debug", false, "Enable detailed error logging")
    flag.Parse()
    keybasePath, err := findKeybasePath()
	if err != nil {
		panic("Could not find Keybase binary: " + err.Error())
	}
client := openai.NewClient("OPENAI-API-KEY") // Hide this key in production!

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
			_, _ = kbc.SendMessage(msg.Message.Channel, "Error: %v", err)
			return
		}

		// Send GPT-3 response to the user via Keybase
		_, _ = kbc.SendMessage(msg.Message.Channel, resp.Choices[0].Message.Content)

		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: resp.Choices[0].Message.Content,
		})
	}
}
