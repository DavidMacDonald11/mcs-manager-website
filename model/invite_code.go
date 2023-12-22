package model

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/davidmacdonald11/mcsm/cmd/env"
)

type InviteCode struct {
	Id        IdType `gorm:"primaryKey"`
	Code      string
	CreatedBy IdType
	ExpiresAt time.Time
}

func CreateInviteCode(createdBy IdType) *InviteCode {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)

	if err != nil {
		return nil
	}

	code := base64.URLEncoding.EncodeToString(bytes)
	expiresAt := time.Now().Add(48 * time.Hour)

	c := &InviteCode{Code: code, CreatedBy: createdBy, ExpiresAt: expiresAt}
	Db.Create(c)

	if c.Id == 0 {
		return nil
	}

	return c
}

func VerifyInviteCode(code string) (createdBy IdType) {
	deleteExpiredCodes()

	c := new(InviteCode)
	Db.First(c, "Code = ?", code)

	if c.Id == 0 {
		c.Code = env.BootstrapInviteCode()

		if len(FindAllUsers()) == 0 && c.Code != "" && code == c.Code {
			return 1
		}

		return 0
	}

	createdBy = c.CreatedBy
	Db.Delete(c)
	return createdBy
}

func deleteExpiredCodes() {
	var codes []InviteCode
	var ids []IdType
	Db.Find(&codes)

	for _, code := range codes {
		if code.ExpiresAt.Before(time.Now()) {
			ids = append(ids, code.Id)
		}
	}

	if len(ids) > 0 {
		Db.Delete(&codes, ids)
	}
}
