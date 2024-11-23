package models

import (
	"time"
)

type MedicalRecord struct {
    MedicalRecordID uint         `gorm:"primaryKey" json:"medical_record_id"`
    PatientID       uint         `json:"patient_id"`
    DoctorID        uint         `json:"doctor_id"`
    AppointmentID   uint         `json:"appointment_id"`
    Diagnosis       string       `json:"diagnosis"`
    TreatmentPlan   string       `json:"treatment_plan"`
    Notes           string       `json:"notes"`
    CreatedAt       time.Time    `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt       time.Time    `gorm:"autoUpdateTime" json:"updated_at"`

    // Relasi ke Patient, Doctor, dan Appointment
    Patient     *Patient     `gorm:"foreignKey:PatientID;references:PatientID;constraint:OnDelete:CASCADE;" json:"patient"`
    Doctor      *Doctor      `gorm:"foreignKey:DoctorID;references:DoctorID;constraint:OnDelete:CASCADE;" json:"doctor"`
    Appointment *Appointment `gorm:"foreignKey:AppointmentID;references:AppointmentID;constraint:OnDelete:CASCADE;" json:"appointment"`
}

// TableName untuk menghindari pluralisasi yang tidak diinginkan
func (MedicalRecord) TableName() string {
    return "medical_records"
}
