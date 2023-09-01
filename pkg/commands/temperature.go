package commands

import (
	"github.com/keybase/go-keybase-chat-bot/kbchat"
)

func HandleTemperatureCommand(api *kbchat.API, channel *kbchat.Channel, temperature float64) error {
	// Your logic for setting the temperature
	// ...

	// Send a response message to the user
	_, err := api.SendMessage(channel, "Temperature set to: "+string(temperature))
	return err
}

