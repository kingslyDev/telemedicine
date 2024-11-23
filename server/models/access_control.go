package models

import (
	"time"
)

type AccessControl struct {
    AccessID    uint      `gorm:"primaryKey" json:"access_id"`
    UserID      uint      `json:"user_id"`
    Resource    string    `gorm:"not null" json:"resource"`
    Permission  string    `gorm:"type:varchar(20);not null" json:"permission"` // 'read', 'write', 'update', 'delete'
    GrantedBy   uint      `json:"granted_by"`
    GrantedDate time.Time `gorm:"autoCreateTime" json:"granted_date"`
    CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`

    // Relations
    User     User  `gorm:"foreignKey:UserID" json:"user"`
    GrantedByUser Admin `gorm:"foreignKey:GrantedBy" json:"granted_by_admin"`
}
