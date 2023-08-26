package engine

import (
	"chatgpt-telegram-bot/cfg"
	"chatgpt-telegram-bot/model"
)

var (
	CHATGPT = "chatgpt"
)

type Engine interface {
	Init(*cfg.Config) error
	Chat(ctx model.ChatContext) (string, error)
	Alive() bool
}

func GetEngine(engineType string) Engine {
	switch engineType {
	case CHATGPT:
		return NewChatGPTEngine()
	default:
		return NewChatGPTEngine()
	}
}
