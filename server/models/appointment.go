package models

import (
	"time"
)

type Appointment struct {
    AppointmentID    uint         `gorm:"primaryKey" json:"appointment_id"`
    PatientID        uint         `json:"patient_id"`
    DoctorID         uint         `json:"doctor_id"`
    AppointmentDate  time.Time    `gorm:"type:date" json:"appointment_date"`
    AppointmentTime  string       `gorm:"type:time" json:"appointment_time"`
    Status           string       `gorm:"type:varchar(20);not null" json:"status"` // 'pending', 'confirmed', 'cancelled', 'completed'
    Reason           string       `json:"reason"`
    CreatedAt        time.Time    `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt        time.Time    `gorm:"autoUpdateTime" json:"updated_at"`

    // Relasi ke Patient, Doctor, dan MedicalRecord
    Patient       *Patient       `gorm:"foreignKey:PatientID;references:PatientID;constraint:OnDelete:CASCADE;" json:"patient"`
    Doctor        *Doctor        `gorm:"foreignKey:DoctorID;references:DoctorID;constraint:OnDelete:CASCADE;" json:"doctor"`
    MedicalRecord *MedicalRecord `gorm:"foreignKey:AppointmentID;references:AppointmentID;constraint:OnDelete:CASCADE;" json:"medical_record,omitempty"`
}

// TableName untuk menghindari pluralisasi yang tidak diinginkan
func (Appointment) TableName() string {
    return "appointments"
}
