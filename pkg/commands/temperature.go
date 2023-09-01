package commands

import (
	"fmt"
	"github.com/keybase/go-keybase-chat-bot/kbchat"
)

func HandleTemperature(api *kbchat.API, msg kbchat.SubscriptionMessage, temp float64) error {
	// Here, handle the 'temperature' change logic.
	resp := fmt.Sprintf("Temperature changed to: %f", temp)
	if _, err := api.SendMessageByConvID(msg.Conversation.Id, resp); err != nil {
		return err
	}
	return nil
}

