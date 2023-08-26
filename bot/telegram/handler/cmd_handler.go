package handler

import (
	"chatgpt-telegram-bot/bot/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotCmd = string

type CommandHandler interface {
	Cmd() BotCmd
	Run(t telegram.TelegramBot, message tgbotapi.Message) error
}
