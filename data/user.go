package data

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
    ID uint64 `gorm:"primaryKey"`
    Username string `gorm:"unique"`
    Hash string
    CreatedAt time.Time
}

func (u User) VerifyPassword(password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(u.Hash), []byte(password))
    return err == nil
}

func (u User) CreateSession(c *fiber.Ctx) bool {
    sess, err := Store.Get(c)
    if err != nil { return false }

    sess.Set(AUTH_KEY, true)
    sess.Set(USER_ID, u.ID)
    err = sess.Save()

    return err == nil
}

func EndSession(c *fiber.Ctx) bool {
    sess, err := Store.Get(c)
    if err != nil { return true }

    if err = sess.Destroy(); err != nil {
        return false
    }

    return true
}

func FindUser(username string) *User {
    user := new(User)
    Db.First(user, "Username = ?", username)

    if user.ID == 0 { return nil }
    return user
}

func CreateUser(username string, hash string) *User {
    user := &User{Username: username, Hash: hash}
    Db.Create(user)

    if user.ID == 0 { return nil }
    return user
}
