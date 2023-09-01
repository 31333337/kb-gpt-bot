package commands

import (
	"github.com/keybase/go-keybase-chat-bot/kbchat"
)

func HandleConfigCommand(api *kbchat.API, channel *kbchat.Channel, configKey string, configValue string) error {
	// Your logic for setting the configuration
	// ...

	// Send a response message to the user
	_, err := api.SendMessage(channel, "Configuration updated.")
	return err
}

