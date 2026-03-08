package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

	e.GET("/crash", func(c echo.Context) error {
		log.Println("Calling crash url")
		os.Exit(1)
		return c.String(http.StatusOK, "Lmaooo")
	})

	if err := e.Start(":8080"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
