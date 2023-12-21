package route

import (
	"github.com/davidmacdonald11/mcsm/model"
	"github.com/davidmacdonald11/mcsm/view/layout"
	"github.com/labstack/echo/v4"
)

func SetupAuthRoutes(app *echo.Echo) {
	app.GET("/auth", getAuth)
	app.POST("/auth/query", postAuthQuery)
}

func getAuth(c echo.Context) error {
	return render(c, layout.Auth(""))
}

func postAuthQuery(c echo.Context) error {
	username := c.FormValue("username")

	if username == "" {
		return render(c, layout.Auth("Username must be provided"))
	}

	user := model.FindUser(username)

	// TODO display login or signup forms
	if user == nil {
		return c.String(200, "NEW")
	}

	return c.String(200, "OLD")
}
