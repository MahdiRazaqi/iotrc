package v1

import "github.com/labstack/echo"

func addAirtemp(c echo.Context) error {
	return c.JSON(200,echo.Map{
		"test":""
	})
}