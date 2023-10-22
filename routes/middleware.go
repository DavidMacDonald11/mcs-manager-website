package routes

import (
	"mcsm/data"

	"github.com/gofiber/fiber/v2"
)

func HTMXRedirect(c *fiber.Ctx, route string) error {
    hxReqHeader := c.GetReqHeaders()["Hx-Request"]

    if len(hxReqHeader) > 0 && hxReqHeader[0] == "true" {
        c.Set("Hx-Redirect", route)
        return nil
    }

    return c.Redirect(route)
}

func LoggedIn(c *fiber.Ctx) error {
    authPaths := []string {
        "/auth",
        "/auth/query",
        "/auth/login",
        "/auth/signup",
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
    return HTMXRedirect(c, "/auth")
}
