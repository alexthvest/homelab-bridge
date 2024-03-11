package telegram

import (
	"github.com/alexthvest/homelab-bridge/pkg/homelab"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Context struct {
	homelab.Context
	api     *tgbotapi.BotAPI
	message *tgbotapi.Message
	args    map[string]string
}

func (c Context) Argument(name string) {

}

func (c Context) Message() *tgbotapi.Message {
	return c.message
}

func (c Context) Reply(text string) error {
	message := tgbotapi.NewMessage(c.Message().Chat.ID, text)
	message.ReplyToMessageID = c.Message().MessageID

	_, err := c.api.Send(message)
	return err
}
