package main

import (
	"net/http"

	"gopkg.in/labstack/echo.v3"
)

func ping(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
