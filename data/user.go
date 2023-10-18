package data

import "time"

type User struct {
    ID uint `json:"id" gorm:"primaryKey"`
    CreatedAt time.Time
    Username string `json:"username"`
    Password string `json:"password"`
}
