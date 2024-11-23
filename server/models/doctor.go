package models

import (
	"time"
)

type Doctor struct {
    DoctorID           uint      `gorm:"primaryKey" json:"doctor_id"`
    UserID             uint      `gorm:"unique;not null" json:"user_id"`
    FirstName          string    `gorm:"not null" json:"first_name"`
    LastName           string    `gorm:"not null" json:"last_name"`
    Specialization     string    `gorm:"not null" json:"specialization"`
    LicenseNumber      string    `gorm:"unique;not null" json:"license_number"`
    YearsOfExperience  int       `json:"years_of_experience"`
    Bio                string    `json:"bio"`
    CreatedAt          time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt          time.Time `gorm:"autoUpdateTime" json:"updated_at"`

    // Relasi ke User
    User *User `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE;" json:"user"`

    // Relations
    Appointments      []Appointment      `gorm:"foreignKey:DoctorID;references:DoctorID;constraint:OnDelete:CASCADE;" json:"appointments,omitempty"`
    MedicalRecords    []MedicalRecord    `gorm:"foreignKey:DoctorID;references:DoctorID;constraint:OnDelete:CASCADE;" json:"medical_records,omitempty"`
    DoctorSchedules   []DoctorSchedule   `gorm:"foreignKey:DoctorID;references:DoctorID;constraint:OnDelete:CASCADE;" json:"doctor_schedules,omitempty"`
    DataMiningResults []DataMiningResult `gorm:"foreignKey:DoctorID;references:DoctorID;constraint:OnDelete:CASCADE;" json:"data_mining_results,omitempty"`
}

// TableName untuk menghindari pluralisasi yang tidak diinginkan
func (Doctor) TableName() string {
    return "doctors"
}
