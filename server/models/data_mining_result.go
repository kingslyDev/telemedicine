package models

import "time"

type DataMiningResult struct {
    ID             uint      `gorm:"primaryKey" json:"id"`
    ImageID        uint      `gorm:"not null" json:"image_id"`
    PatientID      uint      `gorm:"not null" json:"patient_id"`
    DoctorID       uint      `gorm:"not null" json:"doctor_id"`
    DiseaseDetected string   `json:"disease_detected"`
    ConfidenceScore float64  `json:"confidence_score"`
    AnalysisDate   time.Time `json:"analysis_date"`
    CreatedAt      time.Time `json:"created_at"`
    UpdatedAt      time.Time `json:"updated_at"`
}
