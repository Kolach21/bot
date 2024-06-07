package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *Commander) List(inoutMessage *tgbotapi.Message) {
	outputMsg := ""
	for _, p := range c.productService.List() {
		outputMsg += p.Title + "\n"
	}
	msg := tgbotapi.NewMessage(inoutMessage.Chat.ID, outputMsg)
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Fatal(err)
	}
}
