package routes

import (
	"mcsm/data"
	"mcsm/data/form"

	"github.com/gofiber/fiber/v2"
)

var INTERNAL_ERROR = form.INTERNAL_ERROR

func setupAuth(app *fiber.App) {
    app.Get("/auth", getAuth)
    app.Post("/auth/query", postAuthQuery)
    app.Post("/auth/login", postAuthLogin)
    app.Post("/auth/signup", postAuthSignup)
    app.Post("/auth/logout", postAuthLogout)
}

func getAuth(c *fiber.Ctx) error {
    return c.Render("auth", fiber.Map{})
}

func postAuthQuery(c *fiber.Ctx) error {
    u := form.ParseAuth(c, "auth-form", 1)
    if u == nil { return nil }

    user := data.FindUser(u.Username)

    if user == nil {
        return u.RenderAuth(c, "signup-form", "")
    }

    return u.RenderAuth(c, "login-form", "")
}

func postAuthLogin(c *fiber.Ctx) error {
    u := form.ParseAuth(c, "login-form", 2)
    if u == nil { return nil }

    user := data.FindUser(u.Username)

    if user == nil {
        return u.RenderAuth(c, "signup-form", "")
    }

    if !user.VerifyPassword(u.Password) {
        return u.RenderAuth(c, "login-form", "Invalid credentials")
    }

    if !user.CreateSession(c) {
        return u.RenderAuth(c, "login-form", INTERNAL_ERROR)
    }

    return HTMXRedirect(c, "/")
}

func postAuthSignup(c *fiber.Ctx) error {
    u := form.ParseAuth(c, "signup-form", 3)
    if u == nil { return nil }

    if ok := u.VerifyUsername(); ok == nil {
        return u.RenderAuth(c, "signup-form", INTERNAL_ERROR)
    } else if !*ok {
        return u.RenderAuth(c, "signup-form", "Minecraft username is not valid")
    }

    if user := data.FindUser(u.Username); user != nil {
        return u.RenderAuth(c, "signup-form", "Username is already taken")
    }

    if !u.VerifyPassword() {
        return u.RenderAuth(c, "signup-form", "Password is not sufficient")
    }

    hash := u.HashPassword()

    if hash == nil {
        return u.RenderAuth(c, "signup-form", INTERNAL_ERROR)
    }

    // TODO verify invite code

    user := data.CreateUser(u.Username, *hash)

    if user == nil {
        return u.RenderAuth(c, "signup-form", INTERNAL_ERROR)
    }

    if !user.CreateSession(c) {
        return u.RenderAuth(c, "login-form", INTERNAL_ERROR)
    }

    return HTMXRedirect(c, "/")
}

func postAuthLogout(c *fiber.Ctx) error {
    data.EndSession(c)
    return HTMXRedirect(c, "/auth")
}
