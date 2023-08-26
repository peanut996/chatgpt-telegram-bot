package bot

import (
	"chatgpt-telegram-bot/bot/telegram/service"
	"chatgpt-telegram-bot/cfg"
)

var (
	Telegram = "telegram"
)

type Bot interface {
	Init(*cfg.Config) error
	Run()
}

func GetBot(botType string) Bot {
	switch botType {
	case Telegram:
		return service.NewTelegramBot()
	default:
		return service.NewTelegramBot()
	}
}
