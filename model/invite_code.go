package model

import (
	"crypto/rand"
	"encoding/base64"
	"time"
)

type InviteCode struct {
	Id        idType `gorm:"primaryKey"`
	Code      string
	CreatedBy idType
	ExpiresAt time.Time
}

func CreateInviteCode(createdBy idType) *InviteCode {
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

func VerifyInviteCode(code string) (createdBy idType) {
	deleteExpiredCodes()

	c := new(InviteCode)
	Db.First(c, "Code = ?", code)

	if c.Id == 0 {
		return 0
	}

	createdBy = c.CreatedBy
	Db.Delete(c)
	return createdBy
}

func deleteExpiredCodes() {
	var codes []InviteCode
	var ids []idType
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
