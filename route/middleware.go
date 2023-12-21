package route

import (
	"strings"

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

		sess, err := model.GetSession(c)
		loggedIn := err == nil && sess.Values[model.AUTH_KEY] != nil

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
