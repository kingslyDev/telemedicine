package models

import "time"

type AccessControl struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `gorm:"not null" json:"user_id"`
    User        User      `gorm:"constraint:OnDelete:CASCADE;" json:"user"`
    Resource    string    `json:"resource"`
    Permission  string    `gorm:"type:text;not null" json:"permission"` // Ganti enum dengan TEXT
    GrantedBy   uint      `gorm:"not null" json:"granted_by"`
    GrantedDate time.Time `json:"granted_date"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
