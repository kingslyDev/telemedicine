package models

import "time"

type Staff struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null;unique" json:"user_id"` // Foreign key ke tabel Users
	User      User      `gorm:"constraint:OnDelete:CASCADE;" json:"user"` 
    FirstName string    `gorm:"not null" json:"first_name"`
    LastName  string    `gorm:"not null" json:"last_name"`
    Department string   `json:"department"`
    Position   string   `json:"position"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
