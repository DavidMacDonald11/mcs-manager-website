package model

import (
	"log"
	"net/http"
	"time"

	"github.com/davidmacdonald11/mcsm/cmd/env"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/michaeljs1990/sqlitestore"
)

var (
	AUTH_KEY = "authenticated"
	USER_ID  = "user_id"
)

var Store *sqlitestore.SqliteStore

func MustCreateStore() {
	var err error

	Store, err = sqlitestore.NewSqliteStore(
		env.DbPath(),
		"sessions",
		"/",
		3600*24,
		[]byte(env.SessionKey()),
	)

	if err != nil {
		log.Fatal("Failed to create store\n", err.Error())
	}

	log.Println("Created store")
}

func CreateSession(c echo.Context, u *User) bool {
	sess, err := GetSession(c)

	if err != nil {
		return false
	}

	sess.Values[AUTH_KEY] = true
	sess.Values[USER_ID] = u.Id
	sess.Options.SameSite = http.SameSite(http.SameSiteStrictMode)
	sess.Options.Secure = env.IsProd()

	err = sess.Save(c.Request(), c.Response())
	Db.Model(u).Update("last_login", time.Now())

	return err == nil
}

func GetSession(c echo.Context) (*sessions.Session, error) {
	return session.Get("session", c)
}

func GetUserSession(c echo.Context) (*sessions.Session, *User) {
	sess, err := GetSession(c)

	if err != nil || sess.Values[USER_ID] == nil {
		return sess, nil
	}

	id := sess.Values[USER_ID].(IdType)
	user := FindUserById(id)

	return sess, user
}

func EndSession(c echo.Context) bool {
	sess, err := GetSession(c)

	if err != nil {
		return true
	}

	sess.Options.SameSite = http.SameSite(http.SameSiteStrictMode)
	sess.Options.Secure = env.IsProd()
	sess.Save(c.Request(), c.Response())

	err = Store.Delete(c.Request(), c.Response(), sess)
	return err == nil
}
