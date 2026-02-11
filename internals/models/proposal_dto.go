package models

type ProposalDTO struct {
	Email   string `json:"Email" binding:"required,email"`
	Sender  string `json:"Sender" binding:"required"`
	Reciver string `json:"Reciver" binding:"required"`
	Message string `json:"Message" binding:"required"`
}
