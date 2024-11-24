package models

import "time"

type MedicalRecord struct {
    ID            uint      `gorm:"primaryKey" json:"id"`
	PatientID      uint        `gorm:"not null" json:"patient_id"`      // Foreign key ke tabel Patients
	Patient        Patient     `gorm:"constraint:OnDelete:CASCADE;" json:"patient"` // Relasi ke Patients
	DoctorID       uint        `gorm:"not null" json:"doctor_id"`       // Foreign key ke tabel Doctors
	Doctor         Doctor      `gorm:"constraint:OnDelete:CASCADE;" json:"doctor"`  // Relasi ke Doctors
	AppointmentID  *uint       `json:"appointment_id"`                 // Foreign key ke tabel Appointments (opsional)
	Appointment    *Appointment `gorm:"constraint:OnDelete:SET NULL;" json:"appointment"` // Relasi ke Appointments
    Diagnosis     string    `json:"diagnosis"`
    TreatmentPlan string    `json:"treatment_plan"`
    Notes         string    `json:"notes"`
    CreatedAt     time.Time `json:"created_at"`
    UpdatedAt     time.Time `json:"updated_at"`
}
