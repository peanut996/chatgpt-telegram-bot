package main

import (
	"chatgpt-telegram-bot/app"
	"chatgpt-telegram-bot/cfg"
)

func main() {
	c, err := cfg.InitConfig()
	if err != nil {
		panic(err)
	}
	a := app.GetApp()

	err = a.Init(c)
	if err != nil {
		panic(err)
	}
	a.Run()
	a.Block()
}
