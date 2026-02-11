package models

type ProposalDTO struct {
	Email   string `gorm:"type:text;not null"`
	Sender  string `gorm:"type:text;not null"`
	Reciver string `gorm:"type:text;not null"`
	Message string `gorm:"type:text;not null"`
}
