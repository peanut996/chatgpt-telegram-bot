package app

import (
	"chatgpt-telegram-bot/bot"
	"chatgpt-telegram-bot/cfg"
	"chatgpt-telegram-bot/utils"
	"log"
	"os"
	"os/signal"
	"sync"
)

type App struct {
	b bot.Bot
}

func NewApp() *App {
	return &App{}
}

var (
	app  *App
	once sync.Once
)

func GetApp() *App {
	once.Do(func() {
		app = NewApp()
	})
	return app
}

func (a *App) Init(cfg *cfg.Config) error {
	log.Println("[App] Run with config: \n" + utils.ToIndentJson(cfg))
	a.b = bot.GetBot(cfg.BotConfig.BotType)
	err := a.b.Init(cfg)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) Run() {
	go a.b.Run()
}

func (a *App) Block() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit
	log.Println("Shutting down...")
}
