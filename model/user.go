package model

import (
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        idType `gorm:"primaryKey"`
	Username  string `gorm:"unique"`
	Hash      string
	InvitedBy idType
	CreatedAt time.Time
	LastLogin time.Time
}

func (u User) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Hash), []byte(password))
	return err == nil
}

func (u User) CreateSession(c echo.Context) bool {
	sess, err := GetSession(c)

	if err != nil {
		return false
	}

	sess.Values[AUTH_KEY] = true
	sess.Values[USER_ID] = u.Id
	err = sess.Save(c.Request(), c.Response())

	return err == nil
}

func EndSession(c echo.Context) bool {
	sess, err := GetSession(c)

	if err != nil {
		return true
	}

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

func FindUserById(id idType) *User {
	user := new(User)
	Db.First(user, id)

	if user.Id == 0 {
		return nil
	}

	return user
}

func CreateUser(username, hash string, invitedBy idType) *User {
	user := &User{Username: username, Hash: hash, InvitedBy: invitedBy}
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
