package v1

import (
	"github.com/labstack/echo"
)

// Register routes
func Register(e *echo.Echo) {
	v1 := e.Group("/api/v1")

	// authGroup := v1.Group("/auth")
	// authGroup.POST("/register", register)
	// authGroup.POST("/login", login)

	r := v1.Group("/")
	// r.Use(middleware.JWT([]byte(signature)), userRequired)

	airtempGroup := r.Group("airtemp")
	airtempGroup.POST("", addAirtemp)
}
