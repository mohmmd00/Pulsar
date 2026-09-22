package handlers

type RegisterRequest struct { // using gin binding spells
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}
type CreateNotificationRequest struct {
	Channel   string `json:"channel" binding:"required"`
	Recipient string `json:"recipient" binding:"required"`
	Message   string `json:"message" binding:"required"`
}
