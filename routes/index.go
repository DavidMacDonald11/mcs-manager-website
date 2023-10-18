package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupIndexRoutes(app *fiber.App) {
    app.Get("/", getRoot)
}

func getRoot(c *fiber.Ctx) error {
    return c.Render("index", fiber.Map{})
}
