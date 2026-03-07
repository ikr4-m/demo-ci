package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

const versionNumber = "v0"

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(
			http.StatusOK,
			fmt.Sprintf("Hello world from %s revision!", versionNumber),
		)
	})

	if err := e.Start(":8080"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
