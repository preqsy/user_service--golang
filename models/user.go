package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id          uint8     `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	Email       string    `json:"email" gorm:"unique:not null"`
	Password    string    `json:"password" gorm:"not null"`
	TimeCreated time.Time `json:"timeCreated"`
	TimeUpdated time.Time `json:"timeUpdated"`
}
