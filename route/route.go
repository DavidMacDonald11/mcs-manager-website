package route

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/davidmacdonald11/mcsm/model"
	"github.com/davidmacdonald11/mcsm/view/layout"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(app *echo.Echo) {
	app.GET("/", getRoot)
	app.GET("/status", getStatus)
	app.GET("/info", getInfo)
	app.GET("/admin", getAdmin)
	app.POST("/create-invite-code", postCreateInviteCode)

	SetupAuthRoutes(app)
}

func getRoot(c echo.Context) error {
	return redirect(c, "/status")
}

func getStatus(c echo.Context) error {
	_, user := model.GetUserSession(c)

	if user == nil {
		return redirect(c, "/auth")
	}

	return render(c, layout.Status(user.IsAdmin()))
}

func getInfo(c echo.Context) error {
	_, user := model.GetUserSession(c)

	if user == nil {
		return redirect(c, "/auth")
	}

	return render(c, layout.Info(user.IsAdmin()))
}

func getAdmin(c echo.Context) error {
	_, user := model.GetUserSession(c)

	if user == nil {
		return redirect(c, "/auth")
	}

	users := model.FindAllUsers()
	return render(c, layout.Admin(user.IsAdmin(), users))
}

func postCreateInviteCode(c echo.Context) error {
	_, user := model.GetUserSession(c)

	if user == nil {
		return c.String(echo.ErrInternalServerError.Code, "Session Error")
	}

	code := model.CreateInviteCode(user.Id)

	if code == nil {
		return c.String(echo.ErrInternalServerError.Code, "Error")
	}

	return c.String(200, code.Code)
}

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

func redirect(c echo.Context, route string) error {
	hxReqHeader, ok := c.Request().Header["Hx-Request"]

	if ok && len(hxReqHeader) > 0 && hxReqHeader[0] == "true" {
		c.Response().Header().Set("Hx-Redirect", route)
		return nil
	}

	return c.Redirect(http.StatusSeeOther, route)
}
