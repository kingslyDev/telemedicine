package models

import "time"

type Patient struct {
    ID              uint      `gorm:"primaryKey" json:"id"`
    UserID          uint      `gorm:"not null;unique" json:"user_id"` // Foreign key
    User            User      `gorm:"constraint:OnDelete:CASCADE;"`   // Relasi ke User
    FirstName       string    `gorm:"not null" json:"first_name"`
    LastName        string    `gorm:"not null" json:"last_name"`
    DateOfBirth     time.Time `json:"date_of_birth"`
    Gender          string    `gorm:"type:text" json:"gender"`
    Address         string    `json:"address"`
    MedicalHistory  string    `json:"medical_history"`
    EmergencyContact string   `json:"emergency_contact"`
    BloodType       string    `json:"blood_type"`
    Allergies       string    `json:"allergies"`
    CreatedAt       time.Time `json:"created_at"`
    UpdatedAt       time.Time `json:"updated_at"`
}
