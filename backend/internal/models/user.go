package models

import "time"

type User struct {
	ID          uint64    `json:"id"`
	Username    string    `json:"username"`
	PhoneNumber string    `json:"phone_number"`
	Password    string    `json:"password"`
	LastLogin   time.Time `json:"lastLogin"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
