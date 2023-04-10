package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
)

type Command = string

const (
	TelegramApiToken         = "TELEGRAM_API_TOKEN"
	AddCommand       Command = "add"
	UpdateCommand    Command = "update"
	ListCommand      Command = "list"
	SearchCommand    Command = "search"
)

var bot *tgbotapi.BotAPI

func StartTGBot() {
	b, err := tgbotapi.NewBotAPI(os.Getenv(TelegramApiToken))
	if err != nil {
		panic(err)
	}
	bot = b
	bot.Debug = true

	cfg := tgbotapi.NewUpdate(0)
	cfg.Timeout = 30
	for update := range bot.GetUpdatesChan(cfg) {
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			commandManager[update.Message.Command()](update)
			return
		}
		ExecDefault(update)
	}
}
