package models

import "github.com/jinzhu/gorm"

type Auth struct {
	gorm.Model
	AuthUUID string `gorm:"size:255;not null;" json:"auth_uuid,omitempty"`
	Username string `gorm:"size:15;not null;" json:"username"`
	Role     string `json:"role"`
}
