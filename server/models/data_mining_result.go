package models

import (
	"time"
)

type DataMiningResult struct {
	AnalysisID      uint      `gorm:"primaryKey" json:"analysis_id"`
	ImageID         uint      `json:"image_id"`  // FK ke MedicalImage
	PatientID       uint      `json:"patient_id"`
	DoctorID        uint      `json:"doctor_id"`
	DiseaseDetected string    `json:"disease_detected"`  // Penyakit yang terdeteksi
	ConfidenceScore float64   `json:"confidence_score"` // Tingkat kepercayaan prediksi
	AnalysisDate    time.Time `json:"analysis_date"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
