package routes

import (
	"mcsm/data"
	"mcsm/handlers"

	"github.com/gofiber/fiber/v2"
)

func LoggedIn(c *fiber.Ctx) error {
    if c.Path() == "/user/login" {
        return c.Next()
    }

    sess, err := data.Store.Get(c)

    if err != nil || sess.Get(data.AUTH_KEY) == nil {
        return c.Redirect("/user/login")
    }

    return c.Next()
}

func SetupUserRoutes(app *fiber.App) {
    app.Get("/user/login", getUserLogin)
    app.Post("/user/login", postUserLogin)

    app.Post("/user/logout", postUserLogout)
}

func getUserLogin(c *fiber.Ctx) error {
    return c.Render("user/login", fiber.Map{})
}

func postUserLogin(c *fiber.Ctx) error {
    userData := new(handlers.UserData)

    if err := c.BodyParser(userData); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
            "message": err.Error(),
        })
    }

    if userData.Username == "" || userData.Password == "" {
        return c.Render("tmpl/attachedError", fiber.Map {
            "Message": "Both fields are required",
        })
    }

    user := new(data.User)
    data.Db.First(user, "Username = ?", userData.Username)

    if user.ID == 0 {
        return handlers.CreateUser(c, userData)
    }

    return handlers.LoginUser(c, userData, user)
}

func postUserLogout(c *fiber.Ctx) error {
    sess, err := data.Store.Get(c)

    if err == nil {
        err = sess.Destroy()

        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
                "message": err.Error(),
            })
        }
    }

    c.Set("hx-redirect", "/user/login")
    return nil
}
