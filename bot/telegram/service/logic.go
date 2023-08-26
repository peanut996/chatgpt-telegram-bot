package service

import (
	"chatgpt-telegram-bot/constant/cmd"
	"chatgpt-telegram-bot/constant/tip"
	"chatgpt-telegram-bot/model"
	"chatgpt-telegram-bot/utils"
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func IsGPTMessage(message tgbotapi.Message) bool {
	return message.IsCommand() && (message.Command() == cmd.GPT4 || message.Command() == cmd.GPT)
}

func (b *Bot) GetBotInviteLink(code string) string {
	return fmt.Sprintf("https://t.me/%s?start=%s", b.tgBot.Self.UserName, code)
}

func (b *Bot) GetUserInfo(userID int64) (*model.User, error) {
	user, err := b.tgBot.GetChat(tgbotapi.ChatInfoConfig{
		ChatConfig: tgbotapi.ChatConfig{
			ChatID: userID,
		}})
	if err != nil {
		return nil, err
	}
	return model.NewUser(user.UserName, user.FirstName, user.LastName), nil
}

func (b *Bot) sendTyping(chatID int64) {
	msg := tgbotapi.NewChatAction(chatID, tgbotapi.ChatTyping)
	_, _ = b.tgBot.Send(msg)
}

func (b *Bot) SafeSend(msg tgbotapi.MessageConfig) {
	if msg.Text == "" {
		return
	}
	if len(msg.Text) < 4096 {
		b.sendMessageSilently(msg)
		return
	}
	b.sendLargeMessage(msg)
}

func (b *Bot) SendAutoDeleteMessage(msg tgbotapi.MessageConfig, duration time.Duration) {
	newMsg, err := b.tgBot.Send(msg)
	if err != nil {
		log.Println("[SendAutoDeleteMessage]send message failed, err: " + err.Error())
		return
	}
	go func(bot *tgbotapi.BotAPI, message tgbotapi.Message, duration time.Duration) {
		time.Sleep(duration)
		deleteMessage := tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID)
		_, err := bot.Send(deleteMessage)
		if err != nil {
			log.Println("[SendAutoDeleteMessage]delete message failed, err: " + err.Error())
		}
	}(b.tgBot, newMsg, duration)
}

func (b *Bot) SafeSendWithoutPreview(msg tgbotapi.MessageConfig) {
	msg.DisableWebPagePreview = true
	b.SafeSend(msg)
}

func (b *Bot) sendLargeMessage(msg tgbotapi.MessageConfig) {
	msgs := utils.SplitMessageByMaxSize(msg.Text, 4000)
	for _, m := range msgs {
		msg.Text = m
		b.sendMessageSilently(msg)
	}
}

func (b *Bot) sendMessageSilently(msg tgbotapi.MessageConfig) {
	if msg.Text == "" {
		return
	}
	msg.ParseMode = tgbotapi.ModeMarkdown
	_, err := b.tgBot.Send(msg)
	if err != nil {
		log.Printf("[SendMessageSilently] send message failed, err: 【%s】, msg: 【%+v】", err, msg)
		msg.ParseMode = ""
		_, _ = b.tgBot.Send(msg)
	}
}

func (b *Bot) sendFromChatTask(task model.ChatTask) {
	msg := tgbotapi.NewMessage(task.Chat, task.Question)
	msg.ParseMode = tgbotapi.ModeMarkdown
	msg.Text = task.Answer
	msg.ReplyToMessageID = task.MessageID
	msgs := utils.SplitMessageByMaxSize(task.Answer, 4000)
	for _, m := range msgs {
		msg.Text = m
		b.SafeSend(msg)
	}
}

func (b *Bot) SafeSendMsg(chatID int64, text string) {
	b.SafeSend(tgbotapi.NewMessage(chatID, text))
}

func (b *Bot) SafeSendMsgAutoDelete(chatID int64, text string) {
	b.SafeSend(tgbotapi.NewMessage(chatID, text))
}

func (b *Bot) SafeSendMsgWithoutPreview(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.DisableWebPagePreview = true
	b.SafeSend(msg)
}

func (b *Bot) SafeReplyMsg(chatID int64, messageID int, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyToMessageID = messageID
	b.SafeSend(msg)
}

func (b *Bot) SafeReplyMsgWithoutPreview(chatID int64, messageID int, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyToMessageID = messageID
	msg.DisableWebPagePreview = true
	b.SafeSend(msg)
}

func (b *Bot) sendQueueToast(chatID int64, messageID int) {
	queue := len(b.chatTaskChannel)
	if queue < 3 {
		return
	}
	text := fmt.Sprintf(tip.QueueTipTemplate, queue, queue)
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyToMessageID = messageID
	msg.ParseMode = tgbotapi.ModeMarkdown
	b.SafeSend(msg)

}
