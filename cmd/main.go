package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"net/http"
	"os"
	"studyecho/logger"
)

func main() {
	e := echo.New()

	e.Logger = logger.New(os.Stdout)
	e.Logger.SetLevel(log.DEBUG)

	e.Use(middleware.Recover())

	e.GET("/", func(ctx echo.Context) error {
		ctx.Logger().Debug("this is echo logger debug msg.")

		zerolog := ctx.Logger().(*logger.Logger).ZeroLog
		zerolog.Debug().Str("path", ctx.Path()).Msg("this is zerolog debug msg")

		return ctx.HTML(http.StatusOK, "hello world")
	})

	e.Logger.Fatal(e.Start(":2222"))
}
