package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/keybase/go-keybase-chat-bot/kbchat"
	"github.com/31333337/kb-gpt-bot/pkg/commands"
	openai "github.com/sashabaranov/go-openai"
)

var debug bool
var errorStyle = color.New(color.FgHiRed, color.Bold)

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
	path := strings.TrimSpace(string(out))
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

	api := openai.NewAPI("your-openai-api-key")

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
		fmt.Printf("Failed to listen for new messages: %v", err)
		return
	}

	for {
		msg, err := sub.Read()
		if err != nil {
			fmt.Printf("Failed to read message: %v", err)
			return
		}

		if msg.Message.Content.TypeName != "text" {
			continue
		}

		text := msg.Message.Content.Text.Body
		switch {
		case strings.HasPrefix(text, "/info"):
			err = commands.HandleInfo(kbc, msg)
		case strings.HasPrefix(text, "/role"):
			role := strings.TrimSpace(strings.TrimPrefix(text, "/role"))
			err = commands.HandleRole(kbc, msg, role)
		case strings.HasPrefix(text, "/temperature"):
			err = commands.HandleTemperature(kbc, msg, 0.7)
		case strings.HasPrefix(text, "/settings"):
			err = commands.HandleSettings(kbc, msg, nil)
		case strings.HasPrefix(text, "/config"):
			err = commands.HandleConfig(kbc, msg, nil)
		default:
			resp, err := api.ChatCompletion(context.Background(), text)
			if err == nil {
				_, _ = kbc.SendMessageByConvID(msg.Conversation.Id, resp)
			}
		}

		if err != nil {
			fmt.Printf("Command failed: %v", err)
		}
	}
}

