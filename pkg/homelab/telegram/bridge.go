package telegram

import (
	"log"

	"github.com/alexthvest/homelab-bridge/pkg/homelab"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bridge struct {
	api    *tgbotapi.BotAPI
	router *Router
}

func NewBridge(token string, router *Router) (Bridge, error) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return Bridge{}, err
	}

	return Bridge{
		api:    api,
		router: router,
	}, nil
}

func (b Bridge) Listen(ctx homelab.Context) error {
	updateCfg := tgbotapi.NewUpdate(0)
	updateChan := b.api.GetUpdatesChan(updateCfg)

	for update := range updateChan {
		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() {
			continue
		}

		go b.handleCommand(ctx, update.Message)
	}

	return nil
}

func (b Bridge) handleCommand(ctx homelab.Context, message *tgbotapi.Message) {
	routerCtx := Context{
		Context: ctx,
		api:     b.api,
		message: message,
		args:    make(map[string]string),
	}
	if err := b.router.Execute(routerCtx); err != nil {
		errMessage := tgbotapi.NewMessage(message.Chat.ID, err.Error())
		if _, err := b.api.Send(errMessage); err != nil {
			log.Fatal(err)
		}
	}
}
