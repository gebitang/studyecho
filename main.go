package main

/**
follow: https://mp.weixin.qq.com/s/DuEGITdOYHYOXk-v-u2V5A
 */
import (
	echo "github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	// get an instance
	e := echo.New()

	// register router
	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Hello, World")
	})

	// start server
	e.Logger.Fatal(e.Start(":2020"))


}


