package models

import (
	"time"
)

type DataMiningResult struct {
    AnalysisID      uint      `gorm:"primaryKey" json:"analysis_id"`
    ImageID         uint      `json:"image_id"`
    PatientID       uint      `json:"patient_id"`
    DoctorID        uint      `json:"doctor_id"`
    DiseaseDetected string    `json:"disease_detected"`
    ConfidenceScore float64   `gorm:"type:decimal(5,2)" json:"confidence_score"`
    AnalysisDate    time.Time `gorm:"autoCreateTime" json:"analysis_date"`
    CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`

    // Relations
    Image    MedicalImage `gorm:"foreignKey:ImageID" json:"image"`
    Patient  Patient      `gorm:"foreignKey:PatientID" json:"patient"`
    Doctor   Doctor       `gorm:"foreignKey:DoctorID" json:"doctor"`
}
