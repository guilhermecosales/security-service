package model

type UserCredentials struct {
	Email    string
	Password string
}

type AccessTokenData struct {
	AccessToken string `json:"access_token"`
}
