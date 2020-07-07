package main

import (
	"bytes"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"html/template"
	"net/http"
	"time"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/", func(ctx echo.Context) error {
		tpl, err := template.ParseFiles("template/login.html")
		if err != nil {
			ctx.Logger().Error("parse file error: ", err)
			return err
		}

		ctx.Logger().Info("this is login page.")

		data := map[string]interface{}{
			"msg": ctx.QueryParam("msg"),
		}

		if user, ok := ctx.Get("user").(*User); ok {
			data["username"] = user.Username
			data["had_login"] = true
		} else {
			sess := getCookieSession(ctx)
			if flashes := sess.Flashes("username"); len(flashes) > 0 {
				data["username"] = flashes[0]
			}
			// after read Flashes, save session again
			err := sess.Save(ctx.Request(), ctx.Response())
			if err != nil {
				return err
			}
		}
		var buf bytes.Buffer
		err = tpl.Execute(&buf, data)
		if err != nil {
			return err
		}

		return ctx.HTML(http.StatusOK, buf.String())
	})

	// login
	e.POST("/login", func(ctx echo.Context) error {
		username := ctx.FormValue("username")
		passwd := ctx.FormValue("passwd")
		rememberMe := ctx.FormValue("remember_me")

		if username == "gebitang" && passwd == "studyecho" {
			// user standard lib for cookie
			cookie := &http.Cookie{
				Name:     "username",
				Value:    username,
				HttpOnly: true,
			}

			if rememberMe == "1" {
				cookie.MaxAge = 7 * 24 * 3600 // 7 days
			}
			ctx.SetCookie(cookie)

			return ctx.Redirect(http.StatusSeeOther, "/")
		}

		// wrong user or passwd
		sess := getCookieSession(ctx)
		sess.AddFlash(username, "username")
		err := sess.Save(ctx.Request(), ctx.Response())
		if err != nil {
			return ctx.Redirect(http.StatusSeeOther, "/?msg="+err.Error())
		}

		return ctx.Redirect(http.StatusSeeOther, "/?msg=wrong username or password")
	})

	// logout
	e.GET("/logout", func(ctx echo.Context) error {
		cookie := &http.Cookie{
			Name:    "username",
			Value:   "",
			Expires: time.Now().Add(-1e9),
			MaxAge:  -1,
		}
		ctx.SetCookie(cookie)
		return ctx.Redirect(http.StatusSeeOther, "/")
	})

}

func getCookieSession(ctx echo.Context) *sessions.Session {
	sess, _ := cookieStore.Get(ctx.Request(), "request-scope")
	return sess
}
