package models

import (
	"time"
)

type MedicalRecord struct {
	MedicalRecordID uint      `gorm:"primaryKey" json:"medical_record_id"`
	PatientID       uint      `json:"patient_id"`
	DoctorID        uint      `json:"doctor_id"`
	AppointmentID   uint      `json:"appointment_id"`
	Diagnosis       string    `json:"diagnosis"`
	TreatmentPlan   string    `json:"treatment_plan"`
	Notes           string    `json:"notes"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
