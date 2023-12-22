package model

import (
	"net/http"
	"time"

	"github.com/davidmacdonald11/mcsm/cmd/env"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        IdType `gorm:"primaryKey"`
	Username  string `gorm:"unique"`
	Hash      string
	InvitedBy IdType
	CreatedAt time.Time
	LastLogin time.Time
}

func (u User) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Hash), []byte(password))
	return err == nil
}

func (u *User) CreateSession(c echo.Context) bool {
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

func FindUser(username string) *User {
	user := new(User)
	Db.First(user, "Username = ?", username)

	if user.Id == 0 {
		return nil
	}

	return user
}

func FindUserById(id IdType) *User {
	user := new(User)
	Db.First(user, id)

	if user.Id == 0 {
		return nil
	}

	return user
}

func CreateUser(username, password string, invitedBy IdType) *User {
	h, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	hash := string(h)

	if err != nil {
		return nil
	}

	user := &User{
		Username:  username,
		Hash:      hash,
		InvitedBy: invitedBy,
		CreatedAt: time.Now(),
	}

	Db.Create(user)

	if user.Id == 0 {
		return nil
	}

	return user
}

func FindAllUsers() []User {
	users := new([]User)
	Db.Find(users)
	return *users
}
