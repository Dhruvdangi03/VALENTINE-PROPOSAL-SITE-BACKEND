package models

import "time"

type Proposal struct {
	ID          uint      `gorm:"primaryKey"`
	Email       string    `gorm:"type:text;not null"`
	Sender      string    `gorm:"type:text;not null"`
	Reciver     string    `gorm:"type:text;not null"`
	Massege     string    `gorm:"type:text;not null"`
	CreatedTime time.Time `gorm:"autoCreateTime"`
	ExpiryTime  time.Time
	Count       int64 `gorm:"default:0"`
}
