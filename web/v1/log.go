package v1

import (
	"github.com/MahdiRazaqi/iotrc/log"
	"github.com/labstack/echo"
)

type logForm struct {
	Light float64 `json:"light" form:"light"`
}

func addLog(c echo.Context) error {
	formData := new(logForm)
	if err := c.Bind(formData); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	l := &log.Log{
		Light: formData.Light,
	}
	if err := l.InsertOne(); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"message": "added log successfully",
	})
}
