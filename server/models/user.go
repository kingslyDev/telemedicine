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
    
    // Relations
    Patient      *Patient     `gorm:"foreignKey:UserID" json:"patient,omitempty"`
    Doctor       *Doctor      `gorm:"foreignKey:UserID" json:"doctor,omitempty"`
    Staff        *Staff       `gorm:"foreignKey:UserID" json:"staff,omitempty"`
    Admin        *Admin       `gorm:"foreignKey:UserID" json:"admin,omitempty"`
    Notifications []Notification `gorm:"foreignKey:UserID" json:"notifications,omitempty"`
}
