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

    // Relasi ke User
    User *User `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE;" json:"user"`

    // Relations
    LabResults []LabResult `gorm:"foreignKey:UploadedBy;references:StaffID;constraint:OnDelete:CASCADE;" json:"lab_results,omitempty"`
}

// TableName untuk menghindari pluralisasi yang tidak diinginkan
func (Staff) TableName() string {
    return "staff"
}
