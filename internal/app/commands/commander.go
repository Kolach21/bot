package commands

import (
	"bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var registeredCommand = map[string]func(c *Commander, msg *tgbotapi.Message){}

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, productService *product.Service) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) {
	if update.Message == nil { // If we got a message
		return
	}
	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Println(panicValue)
		}
	}()

	command, ok := registeredCommand[update.Message.Command()]
	if ok {
		command(c, update.Message)
	} else {
		c.Default(update.Message)
	}
}
