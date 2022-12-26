package api

type UserResponse struct {
	UserID    string `json:"user_id"`
	Token     string `json:"token"`
	ExpiresIn int64  `json:"expires_in"`
}
