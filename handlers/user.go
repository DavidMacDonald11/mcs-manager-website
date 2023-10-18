package handlers

import (
	"mcsm/data"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserData struct {
    Username string `json:"Username"`
    Password string `json:"Password"`
}

func CreateUser(c *fiber.Ctx, userData *UserData) error {
    passBytes := []byte(userData.Password)

    if len(passBytes) > 72 {
        return c.Render("tmpl/attachedError", fiber.Map {
            "Message": "Password is too long",
        })
    }

    if len(passBytes) < 16 {
        return c.Render("tmpl/attachedError", fiber.Map {
            "Message": "Password is too short",
        })
    }

    password, err := bcrypt.GenerateFromPassword(passBytes, 14)

    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
            "message": err.Error(),
        })
    }

    user := &data.User {
        Username: userData.Username,
        Password: string(password),
    }

    data.Db.Create(user)
    return LoginUser(c, userData, user)
}

func LoginUser(c *fiber.Ctx, userData *UserData, user *data.User) error {
    hash, password := []byte(user.Password), []byte(userData.Password)
    err := bcrypt.CompareHashAndPassword(hash, password)

    if err != nil {
        return c.Render("tmpl/attachedError", fiber.Map {
            "Message": "Invalid password",
        })
    }

    sess, err := data.Store.Get(c)

    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
            "message": err.Error(),
        })
    }

    sess.Set(data.AUTH_KEY, true)
    sess.Set(data.USER_ID, user.ID)

    err = sess.Save()

    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
            "message": err.Error(),
        })
    }

    c.Set("hx-redirect", "/")
    return nil
}
