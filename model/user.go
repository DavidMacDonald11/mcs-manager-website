package model

import (
	"time"

	"github.com/davidmacdonald11/mcsm/cmd/env"
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

func (u User) IsAdmin() bool {
	return u.Username == env.Admin()
}

func (u User) Delete() {
	Db.Where("created_by = ?", u.Id).Delete(&[]InviteCode{})
	Db.Delete(&u)
}

func FindUser(username string) *User {
	user := new(User)
	Db.First(user, "username = ?", username)

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
