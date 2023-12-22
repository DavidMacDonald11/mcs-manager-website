package model

import (
	"log"

	"github.com/davidmacdonald11/mcsm/cmd/env"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/michaeljs1990/sqlitestore"
)

var (
	AUTH_KEY = "authenticated"
	USER_ID = "user_id"
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

func GetSession(c echo.Context) (*sessions.Session, error) {
	return session.Get("session", c)
}

func GetUserSession(c echo.Context) (*sessions.Session, *User) {
	sess, err := GetSession(c)

	if err != nil {
		return sess, nil
	}

	id := sess.Values[USER_ID].(IdType)
	user := FindUserById(id)

	return sess, user
}
