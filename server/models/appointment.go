package models

import (
	"time"
)

type Appointment struct {
    ID              uint      `gorm:"primaryKey" json:"id"`
    PatientID       uint      `gorm:"not null" json:"patient_id"` // Foreign key ke Patient
    Patient         Patient   `gorm:"constraint:OnDelete:CASCADE;"` // Relasi ke Patient
    DoctorID        uint      `gorm:"not null" json:"doctor_id"`  // Foreign key ke Doctor
    Doctor          Doctor    `gorm:"constraint:OnDelete:CASCADE;"` // Relasi ke Doctor
    AppointmentDate time.Time `json:"appointment_date"`
    AppointmentTime time.Time `json:"appointment_time"`
    Status          string    `gorm:"type:text" json:"status"`
    Reason          string    `json:"reason"`
    CreatedAt       time.Time `json:"created_at"`
    UpdatedAt       time.Time `json:"updated_at"`
}

