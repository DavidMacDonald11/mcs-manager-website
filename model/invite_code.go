package model

import (
	"crypto/rand"
	"encoding/base64"
	"time"
)

type idType = uint64

type InviteCode struct {
	ID        idType `gorm:"primaryKey"`
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
	expiresAt := time.Now().Add(48 * time.Hour)

	c := &InviteCode{Code: code, ExpiresAt: expiresAt}
	Db.Create(c)

	if c.ID == 0 {
		return nil
	}

	return c
}

func VerifyInviteCode(code string) bool {
	deleteExpiredCodes()

	c := new(InviteCode)
	Db.First(c, "Code = ?", code)

	if c.ID == 0 {
		return false
	}

	Db.Delete(c)
	return true
}

func deleteExpiredCodes() {
	var codes []InviteCode
	var ids []idType
	Db.Find(&codes)

	for _, code := range codes {
		if code.ExpiresAt.Before(time.Now()) {
			ids = append(ids, code.ID)
		}
	}

	if len(ids) > 0 {
		Db.Delete(&codes, ids)
	}
}
