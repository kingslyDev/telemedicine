package models

import (
	"time"
)

type MedicalImage struct {
    ImageID         uint      `gorm:"primaryKey" json:"image_id"`
    PatientID       uint      `json:"patient_id"`
    UploadedBy      uint      `json:"uploaded_by"`
    ImageType       string    `gorm:"type:varchar(20);not null" json:"image_type"` // 'X-ray', 'CT scan'
    BodyPart        string    `json:"body_part"`
    ImagePath       string    `gorm:"not null" json:"image_path"`
    UploadDate      time.Time `gorm:"autoCreateTime" json:"upload_date"`
    AnalysisStatus  string    `gorm:"type:varchar(20);not null" json:"analysis_status"` // 'pending', 'completed'
    CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`

    // Relations
    Patient       Patient      `gorm:"foreignKey:PatientID" json:"patient"`
    UploadedByUser User        `gorm:"foreignKey:UploadedBy" json:"uploaded_by_user"`
    DataMiningResults []DataMiningResult `gorm:"foreignKey:ImageID" json:"data_mining_results,omitempty"`
}
