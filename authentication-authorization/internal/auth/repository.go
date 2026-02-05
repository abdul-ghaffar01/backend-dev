package auth


type Repository interface{
	AddNewLogin(auth *AuthModel) error
	RevokeToken(token string) error
	RevokeAllTokens(userId string) error
}

