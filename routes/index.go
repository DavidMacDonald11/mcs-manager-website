package routes

import (
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

func SetupIndexRoutes(app *fiber.App) {
    app.Get("/", getRoot)
    app.Get("*", getOther)
}

func getRoot(c *fiber.Ctx) error {
    return c.Render("index", fiber.Map{})
}

func getOther(c *fiber.Ctx) error {
    return c.Render("notfound", fiber.Map{})
}
