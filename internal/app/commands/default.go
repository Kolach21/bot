package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *Commander) Default(inoutMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inoutMessage.From.UserName, inoutMessage.Text)
	msg := tgbotapi.NewMessage(inoutMessage.Chat.ID, "Ты написал "+inoutMessage.Text)
	msg.ReplyToMessageID = inoutMessage.MessageID
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Fatal(err)
	}
}
