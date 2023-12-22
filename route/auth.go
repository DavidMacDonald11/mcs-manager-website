package route

import (
	"github.com/davidmacdonald11/mcsm/cmd/form"
	"github.com/davidmacdonald11/mcsm/model"
	"github.com/davidmacdonald11/mcsm/view/layout"
	"github.com/labstack/echo/v4"
)

func SetupAuthRoutes(app *echo.Echo) {
	app.GET("/auth", getAuth)
	app.POST("/auth/query", postAuthQuery)
	app.POST("/auth/login", postAuthLogin)
	app.POST("/auth/signup", postAuthSignup)
	app.POST("/auth/logout", postAuthLogout)
}

func getAuth(c echo.Context) error {
	return render(c, layout.Auth("", ""))
}

func postAuthQuery(c echo.Context) error {
	username := c.FormValue("username")
	err := form.VerifyUsername(username)

	if err != "" {
		return render(c, layout.Auth(username, err))
	}

	if model.FindUser(username) == nil {
		return render(c, layout.Signup(username, ""))
	}

	return render(c, layout.Login(username, ""))
}

func postAuthLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	err := form.VerifyLogin(username, password)

	if err != "" {
		return render(c, layout.Login(username, err))
	}

	user := model.FindUser(username)

	if user == nil {
		return render(c, layout.Signup(username, "User not found"))
	}

	if !user.VerifyPassword(password) {
		return render(c, layout.Login(username, "Invalid credentials"))
	}

	if !user.CreateSession(c) {
		return render(c, layout.Login(username, "Internal server error"))
	}

	return redirect(c, "/")
}

func postAuthSignup(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	code := c.FormValue("invite-code")
	err := form.VerifySignup(username, password, code)

	if err != "" {
		return render(c, layout.Signup(username, err))
	}

	invitedBy := model.VerifyInviteCode(code)

	if invitedBy == 0 {
		return render(c, layout.Signup(username, "Invite Code is invalid"))
	}

	user := model.CreateUser(username, password, invitedBy)

	if user == nil {
		return render(c, layout.Signup(username, "Internal server error"))
	}

	if !user.CreateSession(c) {
		return render(c, layout.Login(username, "Internal server error"))
	}

	return redirect(c, "/")
}

func postAuthLogout(c echo.Context) error {
	model.EndSession(c)
	return redirect(c, "/")
}
