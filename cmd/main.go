package main

import (
	"log"

	"github.com/davidmacdonald11/mcsm/cmd/env"
	"github.com/davidmacdonald11/mcsm/model"
	"github.com/davidmacdonald11/mcsm/route"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	model.MustConnectToDb()
	model.MustCreateStore()

	app := echo.New()
	app.Static("/public/", "public")

	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     env.AccessControlOrigin(),
		AllowHeaders: []string{
			"Content-Type",
			"Origin",
			"Accept",
			"Authorization",
		},
	}))

	app.Use(session.Middleware(model.Store))
	app.Use(route.IsLoggedIn, route.IsAdmin)

	route.SetupRoutes(app)

	if !env.IsProd() {
		log.Println("WARNING: This is not a production client!")
		log.Println("Security is at risk!")

		app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: `[${time_rfc3339}]  ${status}  ${method} ${host}${path} ${latency_human}` + "\n",
		}))
	}

	log.Fatal(app.Start(env.Port()))
}
