package auth

import "fmt"

type MemoryRepository struct {
	data map[string]*AuthModel
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		data: map[string]*AuthModel{},
	}
}

func (r *MemoryRepository) AddNewLogin(auth *AuthModel) error {
	r.data[auth.ID] = auth
	return nil
}

func (r *MemoryRepository) RevokeToken(token string) error {
	for _, auth := range r.data {
		if auth.RefreshToken == token {
			auth.IsValid = false
			return nil
		}
	}
	return fmt.Errorf("Token not found")
}

func (r *MemoryRepository) RevokeAllTokens(userId string) error {
	for _, auth := range r.data {
		if auth.UserID == userId {
			auth.IsValid = false
		}
	}
	return nil
}
