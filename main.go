package main

import (
	"gopkg.in/labstack/echo.v3"
)

func main() {
	e := echo.New()
	e.GET("/api/v1/ping", ping)
	e.Logger.Fatal(e.Start(":1323"))
}
