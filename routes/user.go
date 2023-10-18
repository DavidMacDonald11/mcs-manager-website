package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
    app.Get("/user/login", getUserLogin)
}

func getUserLogin(c *fiber.Ctx) error {
    return c.Render("user/login", fiber.Map{})
}
