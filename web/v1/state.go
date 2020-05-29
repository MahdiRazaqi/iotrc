package v1

import (
	"github.com/MahdiRazaqi/iotrc/bot"
	"github.com/labstack/echo"
)

func currentState(c echo.Context) error {
	return c.JSON(200, bot.CurrentState)
}
