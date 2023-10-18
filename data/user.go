package data

import "time"

type User struct {
    ID uint64 `gorm:"primaryKey"`
    Username string `gorm:"unique"`
    Password string
    CreatedAt time.Time
}
