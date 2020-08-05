package web

import (
	v1 "github.com/MahdiRazaqi/iotrc/web/v1"
	"github.com/labstack/echo"
)

// Start server
func Start() {
	e := echo.New()
	v1.Register(e)
	e.Start(":80")
}
