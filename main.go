package main

import (
	"log"
	"mcsm/data"
	"mcsm/env"
	"mcsm/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func setupRoutes(app *fiber.App) {
    routes.SetupUserRoutes(app)
    routes.SetupIndexRoutes(app)
}

func main() {
    data.MustConnectToDb()

    app := fiber.New(fiber.Config{Views: html.New("./views", ".html") })
    app.Static("/", "./public")
    setupRoutes(app)

    if !env.IsProd() {
        log.Println("WARNING: This is not a production client!")
        log.Println("Security is at risk!")
    }

    log.Fatal(app.Listen(env.Port()))
}
