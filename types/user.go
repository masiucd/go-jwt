package types

// UserResponse result type
type UserResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

// UserRequest for body request
type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
