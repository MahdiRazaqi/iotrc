package v1

import (
	"github.com/labstack/echo"
)

// Register routes
func Register(e *echo.Echo) {
	v1 := e.Group("/api/v1")

	sensorlogGroup := v1.Group("/sensorlog")
	sensorlogGroup.GET("", listSensorlogs)
	sensorlogGroup.POST("", addSensorlog)
}
