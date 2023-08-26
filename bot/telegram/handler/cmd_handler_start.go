package handler

import (
	"chatgpt-telegram-bot/bot/telegram"
	"chatgpt-telegram-bot/constant/cmd"
	"chatgpt-telegram-bot/constant/tip"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type StartCommandHandler struct {
}

func (h *StartCommandHandler) Cmd() BotCmd {
	return cmd.START
}

func (h *StartCommandHandler) Run(bot telegram.TelegramBot, message tgbotapi.Message) error {
	log.Println(fmt.Printf("get args: [%h]", message.CommandArguments()))
	bot.SafeSendMsg(message.Chat.ID, tip.BotStartTip)
	return nil
}

func NewStartCommandHandler() *StartCommandHandler {
	return &StartCommandHandler{}
}
