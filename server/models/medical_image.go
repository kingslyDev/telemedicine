package models

import (
	"time"
)

type MedicalImage struct {
    ID             uint      `gorm:"primaryKey" json:"id"`
    PatientID      uint      `gorm:"not null" json:"patient_id"` // Foreign key ke Patient
    Patient        Patient   `gorm:"constraint:OnDelete:CASCADE;"` // Relasi ke Patient
    UploadedBy     uint      `gorm:"not null" json:"uploaded_by"`
    ImageType      string    `gorm:"type:text" json:"image_type"`
    BodyPart       string    `json:"body_part"`
    ImagePath      string    `json:"image_path"`
    UploadDate     time.Time `json:"upload_date"`
    AnalysisStatus string    `gorm:"type:text" json:"analysis_status"`
    CreatedAt      time.Time `json:"created_at"`
    UpdatedAt      time.Time `json:"updated_at"`
}
