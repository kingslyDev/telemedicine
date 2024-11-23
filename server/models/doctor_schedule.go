package models

import (
	"time"
)

type DoctorSchedule struct {
	ScheduleID        uint      `gorm:"primaryKey" json:"schedule_id"`
	DoctorID          uint      `json:"doctor_id"`
	AvailableDate     time.Time `gorm:"type:date" json:"available_date"`
	AvailableTimeStart string   `gorm:"type:time" json:"available_time_start"`
	AvailableTimeEnd   string   `gorm:"type:time" json:"available_time_end"`
	IsAvailable        bool     `gorm:"default:true" json:"is_available"`
	CreatedAt          time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
