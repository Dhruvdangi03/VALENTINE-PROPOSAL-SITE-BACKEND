package models

import "time"

type Proposal struct {
	ID          uint      `gorm:"primaryKey"`
	Email       string    `json:"Email" binding:"required,email"`
	Sender      string    `json:"Sender" binding:"required"`
	Reciver     string    `json:"Reciver" binding:"required"`
	Message     string    `json:"Message" binding:"required"`
	CreatedTime time.Time `gorm:"autoCreateTime"`
	ExpiryTime  time.Time
	Count       int64 `gorm:"default:0"`
}
