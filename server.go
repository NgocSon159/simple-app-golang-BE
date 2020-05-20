package main

import (
	"./route"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e:= echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	_, _ = route.NewHotelAppRoute(e)
	e.Logger.Fatal(e.Start(":1323"))

}
