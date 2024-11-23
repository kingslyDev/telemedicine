package models

import (
	"time"
)

type LabResult struct {
    LabResultID uint      `gorm:"primaryKey" json:"lab_result_id"`
    PatientID   uint      `json:"patient_id"`
    TestType    string    `gorm:"not null" json:"test_type"`
    TestResults string    `gorm:"type:text;not null" json:"test_results"` // Bisa berupa JSON atau teks
    TestDate    time.Time `gorm:"type:date" json:"test_date"`
    UploadedBy  uint      `json:"uploaded_by"`
    CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`

    // Relations
    Patient  Patient `gorm:"foreignKey:PatientID" json:"patient"`
    Staff    Staff   `gorm:"foreignKey:UploadedBy" json:"uploaded_by_staff"`
}
