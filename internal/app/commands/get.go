package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	arg := "Product: "
	if args != "" {
		argument, err := strconv.Atoi(args)
		if err != nil {
			return
		}
		product, err := c.productService.Get(argument)
		if err != nil {
			return
		}
		arg += fmt.Sprintf("%v", product.Title)

	}
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		arg,
	)
	c.bot.Send(msg)
}

func init() {
	registeredCommand["get"] = (*Commander).Get
}
