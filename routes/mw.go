package routes

import (
	"mcsm/data"

	"github.com/gofiber/fiber/v2"
)

func LoggedIn(c *fiber.Ctx) error {
    authPaths := []string {
        "/user/auth",
        "/user/auth/query",
        "/user/auth/login",
        "/user/auth/signup",
    }

    sess, err := data.Store.Get(c)
    loggedIn := err == nil && sess.Get(data.AUTH_KEY) != nil

    for _, path := range authPaths {
        if c.Path() == path {
            if loggedIn { return HTMXRedirect(c, "/") }
            return c.Next()
        }
    }

    if loggedIn { return c.Next() }
    return HTMXRedirect(c, "/user/auth")
}
