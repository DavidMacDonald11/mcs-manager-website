package main

import (
	"log"
	"mcsm/data"
	"mcsm/env"
	"mcsm/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
)

func main() {
    data.MustConnectToDb()

    app := fiber.New(fiber.Config{Views: html.New("./views", ".html") })
    app.Static("/", "./public")

    app.Use(routes.LoggedIn, cors.New(cors.Config {
        AllowCredentials: true,
        AllowOriginsFunc: func(origin string) bool {
            return !env.IsProd()
        },
        AllowHeaders: "Content-Type, Origin, Accept",
    }))

    routes.Setup(app)

    if !env.IsProd() {
        log.Println("WARNING: This is not a production client!")
        log.Println("Security is at risk!")
    }

    log.Fatal(app.Listen(env.Port()))
}
