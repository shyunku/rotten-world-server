package models

type User struct {
	ID                int64  `json:"id"`
	AuthID            string `json:"auth_id"`
	EncryptedPassword string `json:"encrypted_password"`
}
