package commands

import (
	"github.com/keybase/go-keybase-chat-bot/kbchat"
)

func HandleRole(api *kbchat.API, msg kbchat.SubscriptionMessage, role string) error {
	// Here, handle the 'role' change logic.
	resp := fmt.Sprintf("Role changed to: %s", role)
	if _, err := api.SendMessageByConvID(msg.Conversation.Id, resp); err != nil {
		return err
	}
	return nil
}

