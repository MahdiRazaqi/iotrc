package main

import (
	"github.com/MahdiRazaqi/iotrc/database"
	"github.com/MahdiRazaqi/iotrc/telebot"
	"github.com/MahdiRazaqi/iotrc/web"
)

func main() {
	database.Connect()
	go telebot.Start()
	web.Start()
}
