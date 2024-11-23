package models

import (
	"time"
)

type AccessControl struct {
    AccessID    uint   `gorm:"primaryKey" json:"access_id"`
    UserID      uint   `json:"user_id"`
    Resource    string `gorm:"not null" json:"resource"`
    Permission  string `gorm:"type:varchar(20);not null" json:"permission"` // 'read', 'write', 'update', 'delete'
    GrantedBy   uint   `json:"granted_by"`
    GrantedDate time.Time `gorm:"autoCreateTime" json:"granted_date"`
    CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`

    // Relasi ke User dan Admin
    User          *User  `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE;" json:"user"`
    GrantedByUser *Admin `gorm:"foreignKey:GrantedBy;references:AdminID;constraint:OnDelete:CASCADE;" json:"granted_by_admin"`
}

// TableName untuk menghindari pluralisasi yang tidak diinginkan
func (AccessControl) TableName() string {
    return "access_controls"
}
