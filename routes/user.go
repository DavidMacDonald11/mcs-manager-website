package routes

import (
	"mcsm/data"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserData struct {
    Username string `json:"Username"`
    Password string `json:"Password"`
    Code string `json:"Code"`
}

func (u *UserData) CanLogin() bool {
    return u.Username != "" && u.Password != ""
}

func (u *UserData) CanSignup() bool {
    return u.CanLogin() && u.Code != ""
}

func renderAuth(c *fiber.Ctx, tmpl string, u UserData, err string) error {
    return c.Render(tmpl, fiber.Map {
        "Username": u.Username,
        "Error": err,
    })
}

const INTERNAL_ERROR = "Internal server error"

func SetupUserRoutes(app *fiber.App) {
    app.Get("/user/auth", getUserAuth)
    app.Post("/user/auth/query", postUserAuthQuery)
    app.Post("/user/auth/login", postUserAuthLogin)

    app.Post("/user/auth/signup", postUserAuthSignup)
    app.Post("/user/auth/logout", postUserAuthLogout)
}

func getUserAuth(c *fiber.Ctx) error {
    return c.Render("user/auth", fiber.Map{})
}

func postUserAuthQuery(c *fiber.Ctx) error {
    var u UserData
    err := c.BodyParser(&u)

    if err != nil {
        return renderAuth(c, "auth-form", u, INTERNAL_ERROR)
    }

    if u.Username == "" {
        return renderAuth(c, "auth-form", u, "Must provide a username")
    }

    var user data.User
    data.Db.First(&user, "Username = ?", u.Username)

    if user.ID == 0 {
        return renderAuth(c, "signup-form", u, "")
    }

    return renderAuth(c, "login-form", u, "")
}

func postUserAuthLogin(c *fiber.Ctx) error {
    var u UserData
    err := c.BodyParser(&u)

    if err != nil {
        return renderAuth(c, "login-form", u, INTERNAL_ERROR)
    }

    if !u.CanLogin() {
        return renderAuth(c, "login-form", u, "Must provide both fields")
    }

    var user data.User
    data.Db.First(&user, "Username = ?", u.Username)

    if user.ID == 0 {
        return renderAuth(c, "signup-form", u, "")
    }

    hash, password := []byte(user.Password), []byte(u.Password)
    err = bcrypt.CompareHashAndPassword(hash, password)

    if err != nil {
        return renderAuth(c, "login-form", u, "Invalid credentials")
    }

    sess, err := data.Store.Get(c)

    if err != nil {
        return renderAuth(c, "login-form", u, INTERNAL_ERROR)
    }

    sess.Set(data.AUTH_KEY, true)
    sess.Set(data.USER_ID, user.ID)

    err = sess.Save()

    if err != nil {
        return renderAuth(c, "login-form", u, INTERNAL_ERROR)
    }

    return HTMXRedirect(c, "/")
}

func postUserAuthSignup(c *fiber.Ctx) error {
    var u UserData
    err := c.BodyParser(&u)

    if err != nil {
        return renderAuth(c, "signup-form", u, INTERNAL_ERROR)
    }

    if !u.CanSignup() {
        return renderAuth(c, "signup-form", u, "Must provide all fields")
    }

    passBytes := []byte(u.Password)

    if len(passBytes) > 72 {
        return renderAuth(c, "signup-form", u, "Password is too long")
    }

    if len(passBytes) < 16 {
        return renderAuth(c, "signup-form", u, "Password is too short")
    }

    password, err := bcrypt.GenerateFromPassword(passBytes, 14)

    if err != nil {
        return renderAuth(c, "signup-form", u, INTERNAL_ERROR)
    }

    // TODO verify invite code

    var user data.User
    data.Db.First(&user, "Username = ?", u.Username)

    if user.ID != 0 {
        return renderAuth(c, "signup-form", u, "Username already taken")
    }

    user = data.User {
        Username: u.Username,
        Password: string(password),
    }

    data.Db.Create(&user)
    sess, err := data.Store.Get(c)

    if err != nil {
        return renderAuth(c, "login-form", u, INTERNAL_ERROR)
    }

    sess.Set(data.AUTH_KEY, true)
    sess.Set(data.USER_ID, user.ID)

    err = sess.Save()

    if err != nil {
        return renderAuth(c, "login-form", u, INTERNAL_ERROR)
    }

    return HTMXRedirect(c, "/")
}

func postUserAuthLogout(c *fiber.Ctx) error {
    sess, err := data.Store.Get(c)

    if err == nil {
        err = sess.Destroy()

        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
                "message": err.Error(),
            })
        }
    }

    return HTMXRedirect(c, "/user/auth")
}
