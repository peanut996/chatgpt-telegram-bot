package model

import (
	"chatgpt-telegram-bot/constant/cmd"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
)

type ChatTask struct {
	Question  string
	Answer    string
	Chat      int64
	From      int64
	MessageID int
	UUID      string

	User          *User
	rawMessage    tgbotapi.Message
	IsGPT4Message bool
}

func (c *ChatTask) String() string {
	return fmt.Sprintf("[ChatTask] gpt-4: %t chat: %d, from: %d, message id: %d, question: %s, answer: %s,]",
		c.IsGPT4Message, c.Chat, c.From, c.MessageID, c.Question, c.Answer)
}

func (c *ChatTask) GetFormattedQuestion() string {
	if c.User != nil {
		return fmt.Sprintf("❓ from %s\n%s", c.User.String(), c.Question)
	}
	return fmt.Sprintf("❓ from %d\n%s", c.From, c.Question)
}

func (c *ChatTask) GetFormattedAnswer() string {
	if c.User != nil {
		return fmt.Sprintf("✅ to %s\n%s", c.User.String(), c.Answer)
	}
	return fmt.Sprintf("✅ to %d\n%s", c.From, c.Answer)
}

func NewChatTask(message tgbotapi.Message) *ChatTask {
	task := &ChatTask{
		Question:   message.Text,
		Chat:       message.Chat.ID,
		From:       message.From.ID,
		MessageID:  message.MessageID,
		UUID:       uuid.New().String(),
		rawMessage: message,
	}
	if message.IsCommand() &&
		(message.Command() == cmd.GPT || message.Command() == cmd.GPT4) {
		task.Question = message.CommandArguments()
		task.IsGPT4Message = message.Command() == cmd.GPT4
	}
	return task
}

func (c *ChatTask) GetRawMessage() tgbotapi.Message {
	return c.rawMessage
}
