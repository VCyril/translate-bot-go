package telegram

import (
	reversoapi "github.com/BRUHItsABunny/go-reverso-api"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const (
	updateTimeout = 60
)

type Bot struct {
	bot    *tgbotapi.BotAPI
	client *reversoapi.ReversoClient
}

func NewBot(bot *tgbotapi.BotAPI, client *reversoapi.ReversoClient) *Bot {
	return &Bot{
		bot:    bot,
		client: client,
	}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	updates, err := b.getUpdatesChan()
	if err != nil {
		return err
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if err := b.handleMessage(update.Message); err != nil {
			return err
		}

	}
	return nil
}

func (b *Bot) getUpdatesChan() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = updateTimeout

	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		return nil, err
	}
	return updates, nil
}
