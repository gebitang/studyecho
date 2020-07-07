package main

import (
	"github.com/labstack/echo/v4"
	"github.com/vmihailenco/msgpack"
	"net/http"
	"strings"
)

type User struct {
	Name string `query:"name" form:"name" json:"name" xml:"name"`
	Sex  string `query:"sex" form:"sex" json:"sex" xml:"sex"`
}

func main() {
	e := echo.New()

	//register binder
	e.Binder = new(MsgpackBinder)

	e.Any("/", func(ctx echo.Context) error {
		user := new(User)

		if err := ctx.Bind(user); err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, user)
	})

	e.Logger.Fatal(e.Start(":2020"))
}

type MsgpackBinder struct{}

func (b *MsgpackBinder) Bind(i interface{}, ctx echo.Context) (err error) {
	// support default binder
	db := new(echo.DefaultBinder)
	if err = db.Bind(i, ctx); err != echo.ErrUnsupportedMediaType {
		return
	}
	req := ctx.Request()
	ctype := req.Header.Get(echo.HeaderContentType)
	if strings.HasPrefix(ctype, echo.MIMEApplicationMsgpack) {
		if err = msgpack.NewDecoder(req.Body).Decode(i); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error()).SetInternal(err)
		}
		return
	}

	return echo.ErrUnsupportedMediaType
}
