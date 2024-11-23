package models

import (
	"time"
)

type LabResult struct {
	LabResultID uint      `gorm:"primaryKey" json:"lab_result_id"`
	PatientID   uint      `json:"patient_id"`  // FK ke Patient
	TestType    string    `gorm:"not null" json:"test_type"` // Jenis tes, misalnya tes darah
	TestResults string    `gorm:"type:text;not null" json:"test_results"` // Hasil tes dalam format JSON atau teks
	TestDate    time.Time `gorm:"type:date" json:"test_date"`
	UploadedBy  uint      `json:"uploaded_by"` // FK ke Staff
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
