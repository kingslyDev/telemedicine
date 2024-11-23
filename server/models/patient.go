package models

import (
	"time"
)

type Patient struct {
    PatientID        uint      `gorm:"primaryKey" json:"patient_id"`
    UserID           uint      `gorm:"unique;not null" json:"user_id"`
    FirstName        string    `gorm:"not null" json:"first_name"`
    LastName         string    `gorm:"not null" json:"last_name"`
    DateOfBirth      time.Time `gorm:"type:date" json:"date_of_birth"`
    Gender           string    `json:"gender"`
    Address          string    `json:"address"`
    MedicalHistory   string    `json:"medical_history"`
    EmergencyContact string    `json:"emergency_contact"`
    BloodType        string    `json:"blood_type"`
    Allergies        string    `json:"allergies"`
    CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt        time.Time `gorm:"autoUpdateTime" json:"updated_at"`

    // Relasi ke User
    User *User `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE;" json:"user"`

    // Relations
    Appointments   []Appointment   `gorm:"foreignKey:PatientID;references:PatientID;constraint:OnDelete:CASCADE;" json:"appointments,omitempty"`
    MedicalRecords []MedicalRecord `gorm:"foreignKey:PatientID;references:PatientID;constraint:OnDelete:CASCADE;" json:"medical_records,omitempty"`
    MedicalImages  []MedicalImage  `gorm:"foreignKey:PatientID;references:PatientID;constraint:OnDelete:CASCADE;" json:"medical_images,omitempty"`
    LabResults     []LabResult     `gorm:"foreignKey:PatientID;references:PatientID;constraint:OnDelete:CASCADE;" json:"lab_results,omitempty"`
}

// TableName untuk menghindari pluralisasi yang tidak diinginkan
func (Patient) TableName() string {
    return "patients"
}
