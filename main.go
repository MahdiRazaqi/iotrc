package main

import (
	"github.com/MahdiRazaqi/iotrc/bot"
	"github.com/MahdiRazaqi/iotrc/database"
	"github.com/MahdiRazaqi/iotrc/web"
)

func main() {
	database.Connect()
	bot.Start()
	web.Start()
}
