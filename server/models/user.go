package models

import (
	"time"
)

// Definisikan konstanta untuk Role
const (
    RolePatient = "patient"
    RoleDoctor  = "doctor"
    RoleStaff   = "staff"
    RoleAdmin   = "admin"
)

type User struct {
    UserID       uint      `gorm:"primaryKey" json:"user_id"`
    Username     string    `gorm:"unique;not null" json:"username"`
    PasswordHash string    `gorm:"not null" json:"-"`
    Email        string    `gorm:"unique;not null" json:"email"`
    PhoneNumber  string    `json:"phone_number"`
    Role         string    `gorm:"type:varchar(20);not null" json:"role"` // 'patient', 'doctor', 'staff', 'admin'
    CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// TableName untuk menghindari pluralisasi yang tidak diinginkan
func (User) TableName() string {
    return "users"
}
