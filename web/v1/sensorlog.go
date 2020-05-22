package v1

import (
	"fmt"

	"github.com/MahdiRazaqi/iotrc/bot"
	"github.com/MahdiRazaqi/iotrc/sensorlog"
	"github.com/labstack/echo"
)

type sensorlogForm struct {
	TempDHT      int  `json:"temp_dht" form:"temp_dht"`
	HumidityDHT  int  `json:"humidity_dht" form:"humidity_dht"`
	DustHumidity int  `json:"dust_humidity" form:"dust_humidity"`
	Light        int  `json:"light" form:"light"`
	PompStatus   bool `json:"pomp_status" form:"pomp_status"`
	LampStatus   bool `json:"lamp_status" form:"lamp_status"`
}

func addSensorlog(c echo.Context) error {
	formData := new(sensorlogForm)
	if err := c.Bind(formData); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	l := &sensorlog.Sensorlog{
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

	msg := fmt.Sprintf("<b>Temp: </b>%v%%0A<b>Humidity: </b>%v%%0A<b>DustHumidity: </b>%v%%0A<b>Light: </b>%v%%0A<b>Pomp Status: </b>%v%%0A<b>Lamp Status: </b>%v%%0A%%0A<code>%v</code>", l.TempDHT, l.HumidityDHT, l.DustHumidity, l.Light, l.PompStatus, l.LampStatus, l.Created.Format("02-01-2006 15:04:05"))

	go bot.SendMessage(1147932122, msg)
	go bot.SendMessage(187805876, msg)

	return c.JSON(200, echo.Map{
		"message": "added log successfully",
	})
}
