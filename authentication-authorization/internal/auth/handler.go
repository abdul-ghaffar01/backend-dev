package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/mssola/useragent"
)

type AuthHandler struct {
	service *AuthService
}

func NewAuthHandler(s *AuthService) *AuthHandler {
	return &AuthHandler{
		service: s,
	}
}

func (h *AuthHandler) AddNewLogin(w http.ResponseWriter, r *http.Request) {
	userId := "12321" // later will take it from db using email
	userAgent := r.UserAgent()
	ua := useragent.New(userAgent)
	browser, _ := ua.Browser()
	country := "Pakistan" // Later will try to extract from IP
	h.service.AddNewLogin(country, browser, userId)
}

func (h *AuthHandler) RevokeToken(w http.ResponseWriter, r *http.Request) {
	token := strings.Split(r.Header.Get("Authorization"), " ")[1]
	h.service.RevokeToken(token)
}

func (h *AuthHandler) RevokeAllTokens(w http.ResponseWriter, r *http.Request) {
	token := strings.Split(r.Header.Get("Authorization"), " ")[1]
	// extract userId from token pass it to the service
	fmt.Println(token)
}
