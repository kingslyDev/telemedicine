package models

import "time"

type Notification struct {
    ID             uint      `gorm:"primaryKey" json:"id"`
    UserID          uint      `gorm:"not null" json:"user_id"`
    User            User      `gorm:"constraint:OnDelete:CASCADE;" json:"user"`
    Message        string    `json:"message"`
    NotificationType string  `json:"notification_type"`
    IsRead         bool      `gorm:"default:false" json:"is_read"`
    SentDate       time.Time `json:"sent_date"`
    CreatedAt      time.Time `json:"created_at"`
    UpdatedAt      time.Time `json:"updated_at"`
}
