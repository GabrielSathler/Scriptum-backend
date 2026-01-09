package models

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Email        string    `gorm:"unique;not null" json:"email"`
	Name         string    `gorm:"not null" json:"name"`
	Password string    `gorm:"not null" json:"password_hash"`
	TypeUser     string    `gorm:"not null" json:"type_user"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
}
