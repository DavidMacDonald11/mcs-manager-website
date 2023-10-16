package main

import (
	"log"
	"mcsm/db"
	"mcsm/env"
	"mcsm/routes"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/gofiber/template/html/v2"
)

func welcome(c *fiber.Ctx) error {
    return c.Render("index", fiber.Map{})
}

func setupRoutes(app *fiber.App) {
    jwt := jwtware.New(jwtware.Config{SigningKey: env.JwtSecret()})

    app.Get("/", welcome)
    app.Get("/test", jwt, routes.Test)

    app.Post("/user/create", routes.CreateUser)
}

func main() {
    db.MustConnectToDb()

    app := fiber.New(fiber.Config{Views: html.New("./views", ".html") })
    app.Static("/", "./public")
    setupRoutes(app)

    log.Fatal(app.Listen(env.Port()))
}
