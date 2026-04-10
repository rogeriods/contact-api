package model

// DTO login
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// DTO contact fields
type Contact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}