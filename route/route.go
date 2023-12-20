package route

import (
	"github.com/a-h/templ"
	"github.com/davidmacdonald11/mcsm/view/layout"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(app *echo.Echo) {
	SetupUserRoutes(app)
	app.GET("/", getRoot)
}

func getRoot(c echo.Context) error {
	return render(c, layout.Root())
}

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}
