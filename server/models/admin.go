package models

import "time"

type Admin struct {
    ID         uint      `gorm:"primaryKey" json:"id"`
    UserID    uint      `gorm:"not null;unique" json:"user_id"` // Foreign key ke tabel User
	User      User      `gorm:"constraint:OnDelete:CASCADE;" json:"user"` // Relasi ke User
	FirstName string    `json:"first_name"`
    LastName   string    `gorm:"not null" json:"last_name"`
    Privileges string    `json:"privileges"` // Can store JSON
    CreatedAt  time.Time `json:"created_at"`
    UpdatedAt  time.Time `json:"updated_at"`
}
