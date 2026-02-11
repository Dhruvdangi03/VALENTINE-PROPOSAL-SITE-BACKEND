package models

type ProposalDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Sender   string `json:"sender" binding:"required"`
	Receiver string `json:"receiver" binding:"required"`
	Message  string `json:"message" binding:"required"`
}
