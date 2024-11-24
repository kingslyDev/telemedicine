package models

import "time"

type Doctor struct {
    ID               uint      `gorm:"primaryKey" json:"id"`
	UserID         uint      `gorm:"not null" json:"user_id"`
    User           User      `gorm:"constraint:OnDelete:CASCADE;" json:"user"`
    FirstName        string    `gorm:"not null" json:"first_name"`
    LastName         string    `gorm:"not null" json:"last_name"`
    Specialization   string    `json:"specialization"`
    LicenseNumber    string    `gorm:"unique" json:"license_number"`
    YearsOfExperience int      `json:"years_of_experience"`
    Bio              string    `json:"bio"`
    AvailableHours   string    `json:"available_hours"` // Can store JSON
    CreatedAt        time.Time `json:"created_at"`
    UpdatedAt        time.Time `json:"updated_at"`
}
