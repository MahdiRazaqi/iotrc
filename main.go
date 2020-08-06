package main

import (
	"github.com/MahdiRazaqi/iotrc/config"
	"github.com/MahdiRazaqi/iotrc/database"
	"github.com/MahdiRazaqi/iotrc/web"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := config.Load(); err != nil {
		logrus.Error(err)
	}

	if err := database.Connect(); err != nil {
		logrus.Error(err)
	}

	web.Start()
}
