package models

import (
	"time"
)

type Staff struct {
    StaffID      uint      `gorm:"primaryKey" json:"staff_id"`
    UserID       uint      `gorm:"unique;not null" json:"user_id"`
    FirstName    string    `gorm:"not null" json:"first_name"`
    LastName     string    `gorm:"not null" json:"last_name"`
    Department   string    `json:"department"`
    Position     string    `json:"position"`
    CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`

    // Relations
    LabResults   []LabResult `gorm:"foreignKey:UploadedBy" json:"lab_results,omitempty"`
}
