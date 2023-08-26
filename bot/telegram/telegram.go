package telegram

import (
	"chatgpt-telegram-bot/cfg"
	"chatgpt-telegram-bot/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"time"
)

type TelegramBot interface {
	Init(*cfg.Config) error
	Run()

	SelfID() int64
	Config() *cfg.Config
	TGBot() *tgbotapi.BotAPI

	GetBotInviteLink(code string) string
	GetUserInfo(userID int64) (*model.User, error)

	SafeSend(msg tgbotapi.MessageConfig)
	SafeReplyMsgWithoutPreview(chatID int64, messageID int, text string)
	SafeSendMsg(chatID int64, text string)
	SendAutoDeleteMessage(msg tgbotapi.MessageConfig, duration time.Duration)
}
