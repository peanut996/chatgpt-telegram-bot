package service

import (
	"chatgpt-telegram-bot/bot/telegram/handler"
	"chatgpt-telegram-bot/cfg"
	"chatgpt-telegram-bot/constant/cmd"
	botError "chatgpt-telegram-bot/constant/error"
	"chatgpt-telegram-bot/constant/tip"
	"chatgpt-telegram-bot/engine"
	"chatgpt-telegram-bot/model"
	"chatgpt-telegram-bot/utils"
	"fmt"
	"log"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/mattn/go-sqlite3"
)

type Bot struct {
	config *cfg.Config
	tgBot  *tgbotapi.BotAPI
	engine engine.Engine

	chatTaskChannel chan model.ChatTask
	handlers        map[handler.BotCmd]handler.CommandHandler
}

func (b *Bot) SelfID() int64 {
	return b.tgBot.Self.ID
}

func (b *Bot) Config() *cfg.Config {
	return b.config
}

func (b *Bot) TGBot() *tgbotapi.BotAPI {
	return b.tgBot
}

func (b *Bot) Init(cfg *cfg.Config) error {
	b.config = cfg

	bot, err := tgbotapi.NewBotAPI(cfg.BotConfig.TelegramBotToken)
	if err != nil {
		return err
	}
	b.tgBot = bot
	b.engine = engine.GetEngine(cfg.EngineConfig.EngineType)
	err = b.engine.Init(cfg)
	if err != nil {
		return err
	}

	b.chatTaskChannel = make(chan model.ChatTask, 100)

	b.handlers = make(map[handler.BotCmd]handler.CommandHandler)

	b.registerCommandHandler(
		handler.NewStartCommandHandler(),
		handler.NewPingCommandHandler())

	go b.loopAndFinishChatTask()

	log.Printf("[Init] telegram bot init success, bot name: %s", b.tgBot.Self.UserName)
	return nil
}

func NewTelegramBot() *Bot {
	return &Bot{}
}

func (b *Bot) Run() {
	log.Println("[Run] start telegram bot")
	go b.fetchUpdates()
}

func (b *Bot) fetchUpdates() {
	config := tgbotapi.NewUpdate(0)
	config.Timeout = 60
	config.AllowedUpdates = []string{"message", "edited_message", "channel_post", "edited_channel_post", "chat_member"}

	botChannel := b.tgBot.GetUpdatesChan(config)
	for {
		select {
		case update, ok := <-botChannel:
			if !ok {
				b.tgBot.StopReceivingUpdates()
				botChannel = b.tgBot.GetUpdatesChan(config)
				log.Println("[FetchUpdates] channel closed, fetch again")
				continue
			}
			go b.handleUpdate(update)
		case <-time.After(30 * time.Second):
		}
	}
}

func (b *Bot) loopAndFinishChatTask() {
	for {
		select {
		case task := <-b.chatTaskChannel:
			b.finishChatTask(task)
		case <-time.After(30 * time.Second):
		}

	}
}

func (b *Bot) finishChatTask(task model.ChatTask) {
	log.Printf("[finishChatTask] start chat task %s", task.String())
	b.sendTyping(task.Chat)

	chatCtx := model.NewChatContext(task.Question, utils.Int64ToString(task.From), "")
	if task.IsGPT4Message {
		chatCtx.Model = "gpt-4"
	}
	res, err := b.engine.Chat(chatCtx)
	if err != nil {
		task.Answer = err.Error()
	} else {
		task.Answer = res
	}
	b.sendTyping(task.Chat)
	b.sendFromChatTask(task)

	log.Printf("[finishChatTask] end chat task: %s", task.String())
}

func (b *Bot) handleUpdate(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}
	log.Printf("[Update] 【msg】:%s", utils.ToJson(update.Message))

	log.Printf("[Update] 【type】:%s, 【from】:%s 【text】: %s",
		update.Message.Chat.Type, model.From(update.Message.From).String(),
		update.Message.Text)

	if update.Message.IsCommand() && !IsGPTMessage(*update.Message) {
		b.execCommand(*update.Message)
		return
	}

	if IsGPTMessage(*update.Message) && strings.Trim(update.Message.CommandArguments(), " ") == "" {
		b.SafeReplyMsg(update.Message.Chat.ID, update.Message.MessageID, fmt.Sprintf(tip.GPTLackTextTipTemplate,
			update.Message.Command(), update.Message.Command()))
		return
	}

	b.handleMessage(*update.Message)

}

func (b *Bot) handleMessage(message tgbotapi.Message) {
	b.publishChatTask(message)
}

func (b *Bot) publishChatTask(message tgbotapi.Message) {
	log.Printf("[publishChatTask] with message %s", utils.ToJson(message))
	chatTask := model.NewChatTask(message)
	user, err := b.GetUserInfo(message.From.ID)
	if err == nil {
		chatTask.User = user
	}

	chatTask.IsGPT4Message = false
	b.chatTaskChannel <- *chatTask

	b.sendTyping(chatTask.Chat)
}

func (b *Bot) registerCommandHandler(handlers ...handler.CommandHandler) {
	for _, commandHandler := range handlers {
		b.handlers[commandHandler.Cmd()] = commandHandler
	}
}

func (b *Bot) execCommand(message tgbotapi.Message) {
	command := message.Command()
	if !cmd.IsBotCmd(command) {
		return
	}
	commandHandler, ok := b.handlers[command]
	if !ok {
		b.SafeSend(tgbotapi.NewMessage(message.Chat.ID, tip.UnknownCmdTip))
		return
	}

	err := commandHandler.Run(b, message)
	if err != nil {
		log.Println("[CommandHandler]exec handler encounter error: " + err.Error())
		b.SafeReplyMsg(message.Chat.ID, message.MessageID, botError.InternalError)
	}
}
