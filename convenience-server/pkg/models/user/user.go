package user

type UserInfo struct {
	ID             string `json:"id"`
	Email          string `json:"email"`
	Verified_email string `json:"verified_email"`
	Name           string `json:"name"`
	Given_name     string `json:"given_name"`
	Picture        string `json:"picture"`
	Locale         string `json:"locale"`
}
