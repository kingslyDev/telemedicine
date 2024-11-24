package models

import "time"

type User struct {
    ID           uint      `gorm:"primaryKey" json:"id"`
    Username     string    `gorm:"unique;not null" json:"username"`
    PasswordHash string    `gorm:"not null" json:"-"`
    Email        string    `gorm:"unique;not null" json:"email"`
    PhoneNumber  string    `json:"phone_number"`
    Role         string    `gorm:"type:user_role;not null" json:"role"` // Menggunakan enum user_role
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
}
