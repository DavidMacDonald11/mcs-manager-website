package routes

import (
	"mcsm/data"

	"github.com/gofiber/fiber/v2"
)

func setupAdmin(app *fiber.App)  {
    app.Get("/admin", getAdmin)
}

func getAdmin(c *fiber.Ctx) error {
    users := data.FindAllUsers()
    return c.Render("admin", fiber.Map{"Users": users})
}
