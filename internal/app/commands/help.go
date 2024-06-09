package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *Commander) Help(inoutMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inoutMessage.Chat.ID, "/help\n/list")
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	registeredCommand["help"] = (*Commander).Help
}
