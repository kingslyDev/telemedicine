package models

import (
	"time"
)

type Notification struct {
    NotificationID   uint      `gorm:"primaryKey" json:"notification_id"`
    UserID           uint      `json:"user_id"`
    Message          string    `gorm:"not null" json:"message"`
    NotificationType string    `gorm:"type:varchar(50)" json:"notification_type"`
    IsRead           bool      `gorm:"default:false" json:"is_read"`
    SentDate         time.Time `gorm:"default:current_timestamp" json:"sent_date"`
    CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt        time.Time `gorm:"autoUpdateTime" json:"updated_at"`

    // Relasi ke User
    User *User `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE;" json:"user"`
}

// TableName untuk menghindari pluralisasi yang tidak diinginkan
func (Notification) TableName() string {
    return "notifications"
}
