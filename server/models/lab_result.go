package models

import "time"

type LabResult struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
	PatientID  uint      `gorm:"not null" json:"patient_id"`
    Patient    Patient   `gorm:"constraint:OnDelete:CASCADE;" json:"patient"`
    TestType    string    `json:"test_type"`
    TestResults string    `json:"test_results"`
    TestDate    time.Time `json:"test_date"`
    UploadedBy  uint      `gorm:"not null" json:"uploaded_by"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
