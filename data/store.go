package data

import (
	"mcsm/env"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/sqlite3"
)

var Store = session.New(session.Config {
    CookieHTTPOnly: true,
    CookieSecure: env.IsProd(),
    Expiration: time.Hour * 24,
    Storage: sqlite3.New(sqlite3.Config{Database: env.DbPath()}),
})

var (
    AUTH_KEY = "authenticated"
    USER_ID = "user_id"
)

func GetUserSession(c *fiber.Ctx) (*session.Session, *User) {
    sess, err := Store.Get(c)

    if err != nil {
        return sess, nil
    }

    id := sess.Get(USER_ID).(uint64)
    user := FindUserByID(id)

    return sess, user
}
