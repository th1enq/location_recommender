package models

import "time"

type User struct {
	ID       uint      `json:"id"`
	Username string    `json:"user_name"`
	Email    string    `json:"email"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
	DeleteAt time.Time `json:"delete_at"`
}
