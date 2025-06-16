package models

type Login struct {
	HashedPassword string
	SessionToken   string
	CSRFToken      string
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
