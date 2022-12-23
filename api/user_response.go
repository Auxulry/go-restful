package api

type UserResponse struct {
	UUID  string
	Name  string
	Email string
}

type UserAuthenticatedData struct {
	UUID  string
	Token string
}
