package model

import (
	"crypto/rand"
	"encoding/base64"
	"time"
)

type InviteCode struct {
	ID        uint64 `gorm:"primaryKey"`
	Code      string
	ExpiresAt time.Time
}

func CreateInviteCode() *InviteCode {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)

	if err != nil {
		return nil
	}

	code := base64.URLEncoding.EncodeToString(bytes)
	expiresAt := time.Now().Add(24 * time.Hour)

	c := &InviteCode{Code: code, ExpiresAt: expiresAt}
	Db.Create(c)

	if c.ID == 0 {
		return nil
	}

	return c
}

func VerifyInviteCode(code string) int {
	c := new(InviteCode)
	Db.First(c, "Code = ?", code)

	if c.ID == 0 {
		return -1
	}

	expired := c.ExpiresAt.Before(time.Now())
	Db.Delete(c)

	if expired {
		return 1
	}

	return 0
}
