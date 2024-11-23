package models

import (
	"time"
)

type Appointment struct {
	AppointmentID   uint      `gorm:"primaryKey" json:"appointment_id"`
	PatientID       uint      `json:"patient_id"`
	DoctorID        uint      `json:"doctor_id"`
	AppointmentDate time.Time `gorm:"type:date" json:"appointment_date"`
	AppointmentTime string    `gorm:"type:time" json:"appointment_time"`
	Status          string    `gorm:"type:varchar(20);not null" json:"status"` // pending, confirmed, cancelled, completed
	Reason          string    `json:"reason"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
