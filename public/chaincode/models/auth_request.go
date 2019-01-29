package models

type LoginRequest struct {
	Email    string     `json:"email"`
	Password string     `json:"password"`
	Role     MemberRole `json:"role"`
}

type LoginResponse struct {
	Message string `json:"message"`
	UserID  string `json:"user_id"`
	Email   string `json:"email"`
}
