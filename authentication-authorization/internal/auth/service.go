package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthService struct {
	repo Repository
}

type JwtClaims struct {
	UserId string `json:"userId"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

var JWTsecret string = "laksdjf8947234jlksjdf089342d"

func NewAuthService(repo Repository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (a *AuthService) AddNewLogin(country, browser, userId string) error {
	var auth AuthModel = AuthModel{Country: country, Browser: browser, UserID: userId}

	auth.ID = uuid.NewString()

	expireDate := jwt.NewNumericDate(time.Now().Add(24 * time.Hour))

	claims := JwtClaims{
		UserId: userId,
		Role:   "user",

		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userId,
			ExpiresAt: expireDate,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "iabdulghaffar-backend",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	refreshToken, err := token.SignedString(JWTsecret)

	if err != nil {
		return fmt.Errorf("Couldn't create a session")
	}

	auth.RefreshToken = refreshToken
	auth.ExpireTime = time.Now().Add(24 * time.Hour)
	auth.IsValid = true

	// saving the login
	err = a.repo.AddNewLogin(&auth)
	if err != nil {
		return fmt.Errorf("Service->AddNewLogin: %s", err.Error())
	}

	return nil
}

func (a *AuthService) RevokeToken(token string) error {
	if token == "" {
		return fmt.Errorf("Token not found")
	}
	return a.repo.RevokeAllTokens(token)
}

func (a *AuthService) RevokeAllTokens(userId string) error {
	if userId == "" {
		return fmt.Errorf("Userid cannot be empty")
	}
	return a.repo.RevokeAllTokens(userId)
}
