package commands

import (
	"strings"

	"github.com/keybase/go-keybase-chat-bot/kbchat"
)

func HandleRoleCommand(api *kbchat.API, channel *kbchat.Channel, role string) error {
	// Your logic for setting the role
	// ...

	// Send a response message to the user
	_, err := api.SendMessage(channel, "Role set to: "+strings.ToUpper(role))
	return err
}

