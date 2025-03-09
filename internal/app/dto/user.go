package dto

import "time"

type User struct {
	UserID            int32     `json:"userID"`
	Email             string    `json:"email"`
	EncryptedPassword string    `json:"encryptedPassword"`
	Token             string    `json:"token"`
	Verified          bool      `json:"verified"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}
