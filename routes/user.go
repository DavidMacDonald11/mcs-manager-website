package routes

import (
	"mcsm/db"
	"mcsm/models"

	"github.com/gofiber/fiber/v2"
)

type User struct {
    ID uint `json:"id"`
    UserName string `json:"user_name"`
}

func CreateResponseUser(userModel models.User) User {
    return User{ID: userModel.ID, UserName: userModel.UserName}
}

func CreateUser(c *fiber.Ctx) error {
    var user models.User

    if err := c.BodyParser(&user); err != nil {
        return c.Status(400).JSON(err.Error())
    }

    db.Instance.Create(&user)
    responseUser := CreateResponseUser(user)

    return c.Status(200).JSON(responseUser)
}

func Test(c *fiber.Ctx) error {
    return c.SendString("Successful test!")
}
