package commands

import (
	"github.com/keybase/go-keybase-chat-bot/kbchat"
)

func HandleInfoCommand(api *kbchat.API, channel *kbchat.Channel) error {
	// Your logic for displaying the bot info
	// ...

	// Send a response message to the user
	_, err := api.SendMessage(channel, "Bot Info: ...") // Replace "..." with actual information
	return err
}

