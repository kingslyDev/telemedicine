package models

import (
	"time"
)

type Admin struct {
	AdminID    uint      `gorm:"primaryKey" json:"admin_id"`
	UserID     uint      `gorm:"unique;not null" json:"user_id"`
	FirstName  string    `gorm:"not null" json:"first_name"`
	LastName   string    `gorm:"not null" json:"last_name"`
	Privileges string    `json:"privileges"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
