package main

import (
	"log"

	"github.com/davidmacdonald11/mcsm/cmd/env"
	"github.com/davidmacdonald11/mcsm/model"
	"github.com/davidmacdonald11/mcsm/route"
	"github.com/labstack/echo/v4"
)

func main() {
	model.MustConnectToDb()

    app := echo.New()
	app.Static("/", "./public")

	route.SetupRoutes(app)

	if !env.IsProd() {
		log.Println("WARNING: This is not a production client!")
		log.Println("Security is at risk!")
	}

    log.Fatal(app.Start(env.Port()))
}
