package models

import (
	"time"
)

type AccessControl struct {
	AccessID    uint      `gorm:"primaryKey" json:"access_id"`
	UserID      uint      `json:"user_id"`     // FK ke User
	Resource    string    `gorm:"not null" json:"resource"` // Resource yang diakses
	Permission  string    `gorm:"type:varchar(20);not null" json:"permission"` // read, write, update, delete
	GrantedBy   uint      `json:"granted_by"`  // FK ke Admin
	GrantedDate time.Time `gorm:"autoCreateTime" json:"granted_date"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
