package main

import "github.com/labstack/echo/v4"

// work as a middleware. auto login for user
func AutoLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		cookie, err := ctx.Cookie("username")
		if err == nil && cookie.Value != "" {
			// search username in db in real env
			user := &User{Username: cookie.Value}

			// set in context
			ctx.Set("user", user)
		}
		return next(ctx)
	}

}
