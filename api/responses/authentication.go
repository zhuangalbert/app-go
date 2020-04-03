package responses

import "time"

type Authentication struct {
	AccessToken string    `json:"access_token"`
	ExpiredIn   int       `json:"expired_in"`
	Username    string    `json:"username"`
	Name        string    `json:"name"`
	Role        string    `json:"role"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
