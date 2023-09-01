package commands

import (
	"fmt"
	"github.com/keybase/go-keybase-chat-bot/kbchat"
)

func HandleSettings(api *kbchat.API, msg kbchat.SubscriptionMessage, settings map[string]string) error {
	// Here, handle the 'settings' change logic.
	resp := "Settings updated."
	for key, value := range settings {
		resp += fmt.Sprintf("\n- %s: %s", key, value)
	}
	if _, err := api.SendMessageByConvID(msg.Conversation.Id, resp); err != nil {
		return err
	}
	return nil
}

