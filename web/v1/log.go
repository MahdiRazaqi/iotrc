package v1

import (
	"github.com/MahdiRazaqi/iotrc/log"
	"github.com/labstack/echo"
)

type logForm struct {
	TempDHT      int  `json:"temp_dht" form:"temp_dht"`
	HumidityDHT  int  `json:"humidity_dht" form:"humidity_dht"`
	DustHumidity int  `json:"dust_humidity" form:"dust_humidity"`
	Light        int  `json:"light" form:"light"`
	PompStatus   bool `json:"pomp_status" form:"pomp_status"`
	LampStatus   bool `json:"lamp_status" form:"lamp_status"`
}

func addLog(c echo.Context) error {
	formData := new(logForm)
	if err := c.Bind(formData); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	l := &log.Log{
		TempDHT:      formData.TempDHT,
		HumidityDHT:  formData.HumidityDHT,
		DustHumidity: formData.DustHumidity,
		Light:        formData.Light,
		PompStatus:   formData.PompStatus,
		LampStatus:   formData.LampStatus,
	}
	if err := l.InsertOne(); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"message": "added log successfully",
	})
}
