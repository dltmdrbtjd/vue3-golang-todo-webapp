package user

import "golang.org/x/oauth2"

type GoogleUserInfo struct {
	UserInfo
	ID             string        `json:"id"`
	Email          string        `json:"email"`
	Verified_email bool          `json:"verified_email"`
	Given_name     string        `json:"given_name"`
	Locale         string        `json:"locale"`
	Token          *oauth2.Token `json:"access_token"`
}

type UserInfo struct {
	Name    string `json:"name"`
	Picture string `json:"picture"`
}
