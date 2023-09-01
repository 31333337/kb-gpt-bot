package commands

import (
	"fmt"
	"github.com/keybase/go-keybase-chat-bot/kbchat"
)

func HandleConfig(api *kbchat.API, msg kbchat.SubscriptionMessage, config map[string]interface{}) error {
	// Here, handle the 'config' change logic.
	resp := "Configuration updated."
	for key, value := range config {
		resp += fmt.Sprintf("\n- %s: %v", key, value)
	}
	if _, err := api.SendMessageByConvID(msg.Conversation.Id, resp); err != nil {
		return err
	}
	return nil
}

