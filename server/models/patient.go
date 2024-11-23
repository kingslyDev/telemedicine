package models

import (
	"time"
)

type Patient struct {
	PatientID       uint      `gorm:"primaryKey" json:"patient_id"`
	UserID          uint      `gorm:"unique;not null" json:"user_id"`
	FirstName       string    `gorm:"not null" json:"first_name"`
	LastName        string    `gorm:"not null" json:"last_name"`
	DateOfBirth     time.Time `gorm:"type:date" json:"date_of_birth"`
	Gender          string    `json:"gender"`
	Address         string    `json:"address"`
	MedicalHistory  string    `gorm:"type:text" json:"medical_history"`
	EmergencyContact string   `json:"emergency_contact"`
	BloodType       string    `json:"blood_type"`
	Allergies       string    `json:"allergies"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
