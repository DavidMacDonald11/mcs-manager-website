package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
    setupAuth(app)
    setupAdmin(app)
    app.Get("/", getRoot)
    app.Get("*", getOther)
}

func getRoot(c *fiber.Ctx) error {
    return c.Render("index", fiber.Map{})
}

func getOther(c *fiber.Ctx) error {
    return c.Render("notfound", fiber.Map{})
}
