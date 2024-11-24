package models

import "time"

type DoctorSchedule struct {
    ID              uint      `gorm:"primaryKey" json:"id"`
	DoctorID          uint      `gorm:"not null" json:"doctor_id"` // Foreign key ke tabel Doctor
	Doctor            Doctor    `gorm:"constraint:OnDelete:CASCADE;" json:"doctor"` // Relasi ke Doctor
    AvailableDate   time.Time `json:"available_date"`
    AvailableTimeStart time.Time `json:"available_time_start"`
    AvailableTimeEnd   time.Time `json:"available_time_end"`
    IsAvailable     bool      `gorm:"default:true" json:"is_available"`
    CreatedAt       time.Time `json:"created_at"`
    UpdatedAt       time.Time `json:"updated_at"`
}
