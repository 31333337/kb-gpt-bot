package commands

import (
	"fmt"
	"github.com/keybase/go-keybase-chat-bot/kbchat"
)

func HandleInfo(api *kbchat.API, msg kbchat.SubscriptionMessage) error {
	resp := "Bot Info:\n- Version: 1.0\n- Commands: /role, /temperature, /settings, /info, /config"
	if _, err := api.SendMessageByConvID(msg.Conversation.Id, resp); err != nil {
		return err
	}
	return nil
}

