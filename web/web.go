package web

import (
	"net/http"

	v1 "github.com/MahdiRazaqi/iotrc/web/v1"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Start server
func Start() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	v1.Register(e)
	e.Start(":8080")
}
