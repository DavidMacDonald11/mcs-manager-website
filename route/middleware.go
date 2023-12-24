package route

import (
	"strings"

	"github.com/davidmacdonald11/mcsm/cmd/env"
	"github.com/davidmacdonald11/mcsm/model"
	"github.com/labstack/echo/v4"
)

var AUTH_PATHS = []string {
	"/auth",
	"/auth/query",
	"/auth/login",
	"/auth/signup",
}

func IsLoggedIn(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if strings.HasPrefix(c.Path(), "/public") {
			return next(c)
		}

		_, user := model.GetUserSession(c)
		loggedIn := user != nil

		for _, path := range AUTH_PATHS {
			if c.Path() == path {
				if loggedIn {
					return redirect(c, "/")
				}

				return next(c)
			}
		}

		if loggedIn {
			return next(c)
		}

		return redirect(c, "/auth")
	}
}

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Path() != "/admin" {
			return next(c)
		}

		_, user := model.GetUserSession(c)

		if user == nil {
			return redirect(c, "/auth")
		}

		if user.Username != env.Admin() {
			return redirect(c, "/")
		}

		return next(c)
	}
}
