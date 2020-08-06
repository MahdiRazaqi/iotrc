package v1

import (
	"strconv"

	"github.com/MahdiRazaqi/iotrc/sensorlog"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
)

type sensorlogForm struct {
	TempDHT      float64 `json:"temp_dht" form:"temp_dht"`
	HumidityDHT  float64 `json:"humidity_dht" form:"humidity_dht"`
	DustHumidity float64 `json:"dust_humidity" form:"dust_humidity"`
	Light        float64 `json:"light" form:"light"`
	MQ9          float64 `json:"mq9" form:"mq9"`
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
		MQ9:          formData.MQ9,
	}
	if err := l.InsertOne(); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"message": "added log successfully",
	})
}

func listSensorlogs(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	sensorlogs, err := sensorlog.Find(bson.M{}, page, limit)
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"sensorlogs": sensorlogs,
	})
}
