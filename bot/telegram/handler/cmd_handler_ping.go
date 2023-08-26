package handler

import (
	"chatgpt-telegram-bot/bot/telegram"
	"chatgpt-telegram-bot/constant/cmd"
	"chatgpt-telegram-bot/constant/tip"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type PingCommandHandler struct {
}

func (c *PingCommandHandler) Cmd() BotCmd {
	return cmd.PING
}

func (c *PingCommandHandler) Run(b telegram.TelegramBot, message tgbotapi.Message) error {
	b.SafeReplyMsgWithoutPreview(message.Chat.ID, message.MessageID, tip.BotPingTip)
	return nil
}
func NewPingCommandHandler() *PingCommandHandler {
	return &PingCommandHandler{}
}
