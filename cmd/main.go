package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"html/template"

	"io"
	"net/http"
	"os"
	"studyecho/logger"
)

//type Template struct {
//	templates *template.Template
//}
//
//func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
//	return t.templates.ExecuteTemplate(w, name, data)
//}

type layoutTemplate struct{}

var LayoutTemplate = &layoutTemplate{}

func (l *layoutTemplate) Render(w io.Writer, contentTpl string, data interface{}, ctx echo.Context) error {
	layout := "layout.html"
	tpl, err := template.New(layout).ParseFiles("../template/common/"+layout, "../template/"+contentTpl)
	if err != nil {
		ctx.Logger().Debug("there is err ", err)
		return err
	}
	return tpl.Execute(w, data)
}

func main() {
	e := echo.New()

	//tpl := &Template{
	//	templates: template.Must(template.ParseGlob("template/*.html")),
	//}
	//e.Renderer = tpl

	e.Renderer = LayoutTemplate

	e.Logger = logger.New(os.Stdout)
	e.Logger.SetLevel(log.DEBUG)

	e.Use(middleware.Recover())

	e.GET("/", func(ctx echo.Context) error {
		ctx.Logger().Debug("this is echo logger debug msg.")

		zerolog := ctx.Logger().(*logger.Logger).ZeroLog
		zerolog.Debug().Str("path", ctx.Path()).Msg("this is zerolog debug msg")

		return ctx.HTML(http.StatusOK, "hello world")
	})

	//e.GET("/render", func(ctx echo.Context) error {
	//	return ctx.Render(http.StatusOK, "key", "gebitang")
	//})

	e.GET("/layout", func(ctx echo.Context) error {
		return ctx.Render(http.StatusOK, "myIndex.html", nil)
	})

	e.Logger.Fatal(e.Start(":2222"))
}
