package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
	"unicode/utf8"
)

type Handler func(update tgbotapi.Update)

var commandManager = map[Command]Handler{
	AddCommand:    ExecAddCommand,
	UpdateCommand: ExecUpdateCommand,
	ListCommand:   ExecListCommand,
	SearchCommand: ExecSearchCommand,
}

func ExecAddCommand(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	msg.ReplyToMessageID = update.Message.MessageID

	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func ExecUpdateCommand(update tgbotapi.Update) {
}

func ExecListCommand(update tgbotapi.Update) {
}

func ExecSearchCommand(update tgbotapi.Update) {
}

func ExecDefault(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.Text = handleMsg(update.Message.Text)
	msg.Entities = textLinkEntity(msg.Text)
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func handleMsg(msg string) string {
	msg = strings.ReplaceAll(msg, "大小：", "📁 大小：")
	msg = strings.ReplaceAll(msg, "关键词：", "🔍 关键词：")
	msg = msg + `

📢 发布频道：阿里云盘吧
🔗 侵删联系：Delete Contact / DMCA`
	return msg
}

func textLinkEntity(text string) []tgbotapi.MessageEntity {
	l := utf8.RuneCountInString(text)
	var e []tgbotapi.MessageEntity
	e = append(e, tgbotapi.MessageEntity{
		Type:     "text_link",
		Offset:   l - 26 - 11,
		Length:   11,
		URL:      "https://t.me/Q66Share",
		User:     nil,
		Language: "",
	}, tgbotapi.MessageEntity{
		Type:     "text_link",
		Offset:   l - 22,
		Length:   26,
		URL:      "https://t.me/Q66Share/13",
		User:     nil,
		Language: "",
	})
	return e
}
