package auth

import "time"

type AuthModel struct {
	ID           string    `json:"id"`
	UserID       string    `json:"userId"`
	RefreshToken string    `json:"refreshToken`
	Country      string    `json:"country`
	Browser      string    `json:"browser"`
	ExpireTime   time.Time `json:"expireTime"`
	IsValid		 bool	   `json:"isValid"`
}