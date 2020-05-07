package main

import (
	"github.com/MahdiRazaqi/iotrc/telebot"
	"github.com/MahdiRazaqi/iotrc/web"
)

func main() {
	go telebot.Start()
	web.Start()
}
